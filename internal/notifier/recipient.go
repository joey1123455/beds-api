package notifier

import "fmt"

var (
	ErrRecipientCannotBeNil = fmt.Errorf("recipient cannot be nil")
)

type Recipient struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func NewRecipient(email, phone string) Recipient {
	return Recipient{
		Email: email,
		Phone: phone,
	}
}

func (e Recipient) GetEmail() string {
	return e.Email
}

func (e Recipient) GetPhone() string {
	return e.Phone
}
