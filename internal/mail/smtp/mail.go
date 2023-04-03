package smtp

import (
	"bytes"
	"github.com/go-gomail/gomail"
	"html/template"
	"log"
)

// Mail ...
type Mail struct {
	From              string
	Password          string
	Host              string
	Port              int
	verifyTmpl        *template.Template
	likeNotifyTmpl    *template.Template
	commentNotifyTmpl *template.Template
}

// New ...
func New(from, password, host string, port int) *Mail {
	verifyTmpl, err := template.ParseFiles("web/templates/verify.html")
	if err != nil {
		log.Println(err)
	}

	likeTmpl, err := template.ParseFiles("web/templates/like_notify.html")
	if err != nil {
		log.Println(err)
	}

	commentTmpl, err := template.ParseFiles("web/templates/comment_notify.html")
	if err != nil {
		log.Println(err)
	}

	return &Mail{
		From:              from,
		Password:          password,
		Host:              host,
		Port:              port,
		verifyTmpl:        verifyTmpl,
		likeNotifyTmpl:    likeTmpl,
		commentNotifyTmpl: commentTmpl,
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

	mail := gomail.NewMessage()
	mail.SetHeader("From", m.From)
	mail.SetHeader("To", email)
	mail.SetHeader("Subject", "Camagru email verification")
	mail.SetBody("text/html", body.String())

	d := gomail.NewDialer(m.Host, m.Port, m.From, m.Password)

	if err := d.DialAndSend(mail); err != nil {
		return err
	}

	return nil
}

// CommentNotify ...
func (m *Mail) CommentNotify(email string) error {
	var body bytes.Buffer

	data := struct {
		User string
	}{
		User: "user",
	}

	if err := m.commentNotifyTmpl.Execute(&body, data); err != nil {
		return err
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", m.From)
	mail.SetHeader("To", email)
	mail.SetHeader("Subject", "New activity")
	mail.SetBody("text/html", body.String())

	d := gomail.NewDialer(m.Host, m.Port, m.From, m.Password)

	if err := d.DialAndSend(mail); err != nil {
		return err
	}

	return nil
}

// LikeNotify ...
func (m *Mail) LikeNotify(email string) error {
	var body bytes.Buffer

	data := struct {
		User string
	}{
		User: "user",
	}

	if err := m.likeNotifyTmpl.Execute(&body, data); err != nil {
		return err
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", m.From)
	mail.SetHeader("To", email)
	mail.SetHeader("Subject", "New activity")
	mail.SetBody("text/html", body.String())

	d := gomail.NewDialer(m.Host, m.Port, m.From, m.Password)

	if err := d.DialAndSend(mail); err != nil {
		return err
	}

	return nil
}
