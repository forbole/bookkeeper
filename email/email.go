package email

import (
	"log"
	"net/mail"
	"net/smtp"

	"github.com/scorredoira/email"
)

func SendEmail(subject string, body string)error {
	// compose the message
	m := email.NewMessage("Bookkeeper Monthy Report", "Enjoy")
	m.From = mail.Address{Name: "From", Address: "[email protected]"}
	m.To = []string{"[email protected]"}

	// add attachments
	if err := m.Attach("email.go"); err != nil {
		log.Fatal(err)
	}

	// add headers
	m.AddHeader("X-CUSTOMER-id", "xxxxx")

	// send it
	auth := smtp.PlainAuth("", "[email protected]", "pwd", "smtp.zoho.com")
	if err := email.Send("smtp.zoho.com:587", auth, m); err != nil {
		return err
	}
}