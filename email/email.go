package email

import (
	"net/mail"
	"net/smtp"

	"github.com/forbole/bookkeeper/types"
	"github.com/scorredoira/email"
)

// SendEmail send an an email by Gmail. It attach an attachment which locate on the input relative path
func SendEmail(emailDetails types.EmailDetails,path string)error {
	// compose the message
	m := email.NewMessage(emailDetails.Subject, emailDetails.Details)
	m.From = mail.Address{Name: emailDetails.From.Name, Address: emailDetails.From.Address}
	m.To = emailDetails.To

	// add attachments
	if err := m.Attach(path); err != nil {
		return nil
	} 

	// send it
	auth := smtp.PlainAuth("", emailDetails.From.Name, emailDetails.From.Password, "smtp.gmail.com")
	if err := email.Send("smtp.gmail.com:587", auth, m); err != nil {
		return err
	}
	return nil
}