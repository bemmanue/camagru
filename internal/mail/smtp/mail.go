package smtp

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
)

// Mail ...
type Mail struct {
	Address    string
	From       string
	Password   string
	Host       string
	verifyTmpl *template.Template
}

// New ...
func New(address, from, password, host string) *Mail {
	t, err := template.ParseFiles("web/templates/verify.html")
	if err != nil {
		log.Println(err)
	}

	return &Mail{
		Address:    address,
		From:       from,
		Password:   password,
		Host:       host,
		verifyTmpl: t,
	}
}

// Verify ...
func (m *Mail) Verify(email, code string) error {
	var body bytes.Buffer

	data := struct {
		Email string
		Code  string
	}{
		Email: email,
		Code:  code,
	}

	if err := m.verifyTmpl.Execute(&body, data); err != nil {
		return err
	}

	msg := "From: " + m.From + "\n" +
		"To: " + email + "\n" +
		"Subject: Email verification\n" +
		"MIME-version: 1.0;\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\n\n" +
		body.String()

	err := smtp.SendMail(
		m.Address,
		smtp.PlainAuth("", m.From, m.Password, "smtp.mail.ru"),
		m.From,
		[]string{email},
		[]byte(msg),
	)

	return err
}
