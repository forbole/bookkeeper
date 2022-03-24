package email

import (
	"net/mail"
	"net/smtp"
	"time"

	"github.com/forbole/bookkeeper/types"
	"github.com/scorredoira/email"
)

// SendEmail send an an email by Gmail. It attach an zip which compress the provided file
func SendEmail(emailDetails types.EmailDetails,path []string)error {
	// compose the message
	m := email.NewMessage(emailDetails.Subject, emailDetails.Details)
	m.From = mail.Address{Name: emailDetails.From.Name, Address: emailDetails.From.Address}
	m.To = emailDetails.To

	//zip the files
	zipName:="bookkeeper_"+time.Now().String()+".zip"
	if err := ZipFiles(zipName, path); err != nil {
        return err
    }

	

	// add attachments
	if err := m.Attach(zipName); err != nil {
		return nil
	} 

	// send it
	auth := smtp.PlainAuth("", emailDetails.From.Name, emailDetails.From.Password, "smtp.gmail.com")
	if err := email.Send("smtp.gmail.com:587", auth, m); err != nil {
		return err
	}
	return nil
}