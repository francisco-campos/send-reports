package email

import "testing"

func TestSendEmail(t *testing.T) {
	from := "franciscojavier.c.c@gmail.com"
	pass := "fr1nc3sc4151984"
	to := []string{"francisco.campos@fusiona.cl", "franciscojavier.c.c@gmail.com"}
	subject := "testing"
	body := "hola!"
	file := ""
	SendEmail(from, pass, to, subject, body, file)
}
