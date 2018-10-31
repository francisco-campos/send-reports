package email

import (
	"log"
	"net/smtp"
)

const addrMailSender = "smtp.gmail.com:587"
const hostMailSender = "smtp.gmail.com"

//SendEmail func send mail
func SendEmail(from string, pass string, to []string, body string, file string) {
	err := smtp.SendMail(addrMailSender,
		smtp.PlainAuth("", from, pass, hostMailSender),
		from, to, []byte(body))

	log.Printf("correo enviado: %s", from)
	if err != nil {
		log.Printf("error SendEmail: %s", err)
	}
}
