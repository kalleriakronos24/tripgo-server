package mail

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
	"google.golang.org/appengine"
	"google.golang.org/appengine/mail"
	"gopkg.in/gomail.v2"
)

type TCarbonCopy struct {
	Address string `json:"address,omitempty"`
	Name    string `json:"name,omitempty"`
}

type TSendMail struct {
	From        string         `json:"from,omitempty"`
	MailTo      string         `json:"mailTo,omitempty"`
	Cc          []*TCarbonCopy `json:"cc,omitempty"`
	Subject     string         `json:"subject,omitempty"`
	Body        string         `json:"body,omitempty"`
	Template    string         `json:"template,omitempty"`
	Attachments string         `json:"attachment,omitempty"`
}

func SendEmail(t *TSendMail) *gomail.Message {

	mail := gomail.NewMessage()
	mail.SetHeader("From", t.From)
	mail.SetHeader("To", t.MailTo)

	if len(t.Cc) > 0 {
		for _, v := range t.Cc {
			mail.SetAddressHeader("Cc", v.Address, v.Name)
		}
	}
	mail.SetHeader("Subject", t.Subject)
	mail.SetBody(t.Body, t.Template)
	mail.Attach(t.Attachments)

	d := gomail.NewDialer(viper.GetString("MAIL_HOST"), viper.GetInt("MAIL_PORT"), viper.GetString("MAIL_USERNAME"), viper.GetString("MAIL_PASSWORD"))

	if err := d.DialAndSend(mail); err != nil {
		panic(err)
	}

	return mail
}

func SendMailV2(w http.ResponseWriter, r *http.Request) {
	// Get the App Engine context from the request
	ctx := appengine.NewContext(r)

	// Create a new email message
	msg := &mail.Message{
		Sender:   "example@example.com",
		To:       []string{"recipient@example.com"},
		Subject:  "Hello from App Engine!",
		Body:     "This is a plain text email.",
		HTMLBody: "<html><body><p>This is an <b>HTML</b> email.</p></body></html>",
	}

	// Send the email
	if err := mail.Send(ctx, msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Email sent successfully!"))
}

func SendMailV3(p *TSendMail) (err error) {
	message := gomail.NewMessage()
	message.SetHeader("From", p.From)
	message.SetHeader("To", p.MailTo)
	message.SetHeader("Subject", p.Subject)
	message.AddAlternative("text/html", p.Body)

	dialer := gomail.NewDialer("live.smtp.mailtrap.io", 587, "api", "eeb698e014595388ad1d80ed7e9ffb54")
	if err := dialer.DialAndSend(message); err != nil {
		fmt.Println("Error:", err)
		return err
	} else {
		fmt.Println("HTML Email sent successfully")
	}
	return err
}
