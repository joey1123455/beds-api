package notifier

// import (
// 	"bytes"
// 	"errors"
// 	"fmt"
// 	"html/template"
// 	"path/filepath"
// 	"strconv"

// 	"github.com/joey1123455/beds-api/internal/common"
// 	"github.com/joey1123455/beds-api/internal/config"
// 	"gopkg.in/gomail.v2"
// )

// var (
// 	ErrNoEmailData            = errors.New("no email data provided in notification struct, email data should be keyed with 'email'")
// 	ErrIncorrectEmailDataType = errors.New("incorrect email data type, notification data keyed with 'email' should be of type 'EmailData'")
// )

// // EmailData struct holds the data for sending an email.
// type EmailData struct {
// 	To          string
// 	Subject     string
// 	Body        string
// 	Data        map[string]interface{}
// 	Attachments []string
// }

// // mailerImpl struct implements the Mailer interface.
// type mailerImpl struct {
// 	config    config.Config
// 	templates map[string]*template.Template
// }

// // NewMailer initializes the mailer with the given configuration.
// func NewMailer(config config.Config) (Notifier, error) {
// 	m := &mailerImpl{
// 		config:    config,
// 		templates: make(map[string]*template.Template),
// 	}
// 	fmt.Println(m)
// 	err := m.loadTemplates()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return m, nil
// }

// // func (m *mailerImpl) convertNotification(templateName string, emailData EmailData) error {

// // loadTemplates loads HTML templates from the specified directory.
// func (m *mailerImpl) loadTemplates() error {
// 	files, err := filepath.Glob(m.config.TemplateDirectory + "/*.html")
// 	if err != nil {
// 		return fmt.Errorf("error loading templates: %v", err)
// 	}
// 	fmt.Println(files)

// 	for _, file := range files {
// 		name := filepath.Base(file)
// 		tmpl, err := template.ParseFiles(file)
// 		if err != nil {
// 			return fmt.Errorf("error parsing template %s: %v", name, err)
// 		}
// 		m.templates[name] = tmpl
// 	}

// 	return nil
// }

// func (m *mailerImpl) SendNotification(templateName string, data *Notification) error {

// 	sentData, ok := data.Data["email"]
// 	if !ok {
// 		return ErrNoEmailData
// 	}

// 	emailData, ok := sentData.(EmailData)
// 	if !ok {
// 		return ErrIncorrectEmailDataType
// 	}

// 	m.SendEmail(templateName, emailData)
// 	return nil
// }

// // SendEmail parses the selected template and sends the email.
// func (m *mailerImpl) SendEmail(templateName string, emailData EmailData) error {

// 	tmpl, ok := m.templates[templateName]
// 	if !ok {
// 		return fmt.Errorf("template %s not found", templateName)
// 	}
// 	emailData.Data["Company"] = template.HTML(common.CompanyName)

// 	var body bytes.Buffer
// 	err := tmpl.Execute(&body, emailData.Data)
// 	if err != nil {
// 		return fmt.Errorf("error executing template: %v", err)
// 	}

// 	emailData.Body = body.String()

// 	return m.sendEmail(emailData)
// }

// // send sends an email with the provided data using gomail.
// func (m *mailerImpl) sendEmail(emailData EmailData) error {
// 	msg := gomail.NewMessage()
// 	msg.SetHeader("From", m.config.SmtpEmail)
// 	msg.SetHeader("To", emailData.To)
// 	msg.SetHeader("Subject", emailData.Subject)
// 	msg.SetBody("text/html", emailData.Body)

// 	for _, attachment := range emailData.Attachments {
// 		msg.Attach(attachment)
// 	}

// 	intSMTPPort, err := strconv.Atoi(m.config.SmtpPort)
// 	if err != nil {
// 		return fmt.Errorf("error: %v", err)
// 	}

// 	dialer := gomail.NewDialer(m.config.SmtpHost, intSMTPPort, m.config.SmtpUsername, m.config.SmtpPassword)
// 	fmt.Println(dialer)
// 	if err := dialer.DialAndSend(msg); err != nil {
// 		return fmt.Errorf("error sending email: %v", err)
// 	}

// 	return nil
// }
