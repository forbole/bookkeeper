package email

import (
	"net/mail"
	"net/smtp"

	"github.com/forbole/bookkeeper/types"
	"github.com/scorredoira/email"
)

func SendEmail(emailDetails types.EmailDetails)error {
	// compose the message
	m := email.NewMessage(emailDetails.Subject, emailDetails.Details)
	m.From = mail.Address{Name: emailDetails.From.Name, Address: emailDetails.From.Address}
	m.To = emailDetails.To

	// add attachments
	if err := m.Attach("bitcoin.csv"); err != nil {
		return nil
	} 

	// add headers
	m.AddHeader("X-CUSTOMER-id", "xxxxx")

	// send it
	auth := smtp.PlainAuth("", emailDetails.From.Name, emailDetails.From.Password, "smtp.gmail.com")
	if err := email.Send("smtp.gmail.com:587", auth, m); err != nil {
		return err
	}
	return nil
}