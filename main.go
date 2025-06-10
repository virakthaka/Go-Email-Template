package main

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
	"path/filepath"
)

type EmailData struct {
	Username string
	Body     string
	Link     string
}

func send(to, subject string, data EmailData) error {
	from := "your_email@gmail.com"
	password := "xxxxx"

	tmplPath := filepath.Join("templates", "change-pass.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return err
	}

	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n"
	msg += "Subject: " + subject + "\n"
	msg += "From: " + from + "\n"
	msg += "To: " + to + "\n\n"
	msg += body.String()

	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")
	return smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(msg))
}

func main() {
	if err := send(
		"recipient@example.com",
		"Change Account Password",
		EmailData{
			Username: "JohnDoe",
			Body:     "Your password has been changed successfully.",
			Link:     "https://domain.com/ref=xxxxx",
		},
	); err != nil {
		log.Fatal("Email send failed:", err)
	}
	log.Println("Email sent successfully.")
}
