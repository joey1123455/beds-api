package mailer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"path/filepath"
	"strconv"

	"github.com/joey1123455/beds-api/internal/common"
	"github.com/joey1123455/beds-api/internal/config"
	"github.com/joey1123455/beds-api/internal/notifier"
	"gopkg.in/gomail.v2"
)

var (
	ErrNoEmailData            = errors.New("no email data provided in notification struct, email data should be keyed with 'email'")
	ErrIncorrectEmailDataType = errors.New("incorrect email data type, notification data keyed with 'email' should be of type 'EmailData'")
)

// Mailer interface defines the methods for sending emails.

// EmailData struct holds the data for sending an email.
type EmailData struct {
	To          string
	Subject     string
	Body        string
	Data        map[string]interface{}
	Attachments []string
}

// mailerImpl struct implements the Mailer interface.
type mailerImpl struct {
	key       string
	config    config.Config
	templates map[string]*template.Template
}

// NewMailer initializes the mailer with the given configuration.
func NewMailer(config config.Config) (notifier.Notifier, error) {
	m := &mailerImpl{
		config:    config,
		templates: make(map[string]*template.Template),
		key:       "email",
	}
	err := m.loadTemplates()
	if err != nil {
		return nil, err
	}

	for k := range m.templates {
		log.Println("loaded template: ", k)
	}
	return m, nil
}

// loadTemplates loads HTML templates from the specified directory.
func (m *mailerImpl) loadTemplates() error {
	files, err := filepath.Glob(m.config.TemplateDirectory + "/*.html")
	if err != nil {
		return fmt.Errorf("error loading templates: %v", err)
	}
	for _, file := range files {

		name := filepath.Base(file)
		log.Println(name)
		tmpl, err := template.ParseFiles(file)
		if err != nil {
			return fmt.Errorf("error parsing template %s: %v", name, err)
		}
		m.templates[name] = tmpl
	}

	return nil
}

func (m *mailerImpl) SendNotification(templateName string, data *notifier.Notification) error {
	sentData, ok := data.Data["email"]
	if !ok {
		return ErrNoEmailData
	}

	jsonData, err := json.Marshal(sentData)
	if err != nil {
		return err
	}

	var emailData EmailData
	err = json.Unmarshal(jsonData, &emailData)
	if err != nil {
		return err
	}

	err = m.Send(templateName, emailData)
	if err != nil {
		return err
	}

	return nil
}

func (m *mailerImpl) Key() string {
	return m.key
}

// SendEmail parses the selected template and sends the email.
func (m *mailerImpl) Send(templateName string, emailData EmailData) error {
	tmpl, ok := m.templates[templateName]
	if !ok {
		return fmt.Errorf("template %s not found", templateName)
	}
	emailData.Data["Company"] = template.HTML(common.CompanyName)

	var body bytes.Buffer
	err := tmpl.Execute(&body, emailData.Data)
	if err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	emailData.Body = body.String()

	return m.send(emailData)
}

// send sends an email with the provided data using gomail.
func (m *mailerImpl) send(emailData EmailData) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", m.config.SmtpEmail)
	msg.SetHeader("To", emailData.To)
	msg.SetHeader("Subject", emailData.Subject)
	msg.SetBody("text/html", emailData.Body)

	for _, attachment := range emailData.Attachments {
		msg.Attach(attachment)
	}

	intSMTPPort, err := strconv.Atoi(m.config.SmtpPort)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	dialer := gomail.NewDialer(m.config.SmtpHost, intSMTPPort, m.config.SmtpUsername, m.config.SmtpPassword)
	if err := dialer.DialAndSend(msg); err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	return nil
}
