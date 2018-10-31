package email

import (
	"log"

	"gopkg.in/gomail.v2"
)

const hostMailSender = "smtp.gmail.com"

//SendEmail func send mail
func SendEmail(from string, pass string, to []string, subject, body string, file string) bool {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to[0])
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	if file != "" {
		m.Attach(file)
	}

	d := gomail.NewPlainDialer(hostMailSender, 587, from, pass)
	err := d.DialAndSend(m)
	if err != nil {
		log.Printf("error SendEmail: %s", err)
		return false
	}

	return true
}
