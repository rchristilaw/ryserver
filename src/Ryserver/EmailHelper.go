// EmailHelper
package main

import (
	"bytes"
	"fmt"
	"html"
	"log"
	"net/http"
	"net/smtp"
)

type EmailUser struct {
	Username    string
	Password    string
	EmailServer string
	Port        int
}

func SendEmail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sending Test Email, %q", html.EscapeString(r.URL.Path))
	CreateEmail("username", "password") // these must be actual credentials to send emails
}

func CreateEmail(username string, password string) {
	emailUser := &EmailUser{username, password, "smtp.gmail.com", 587}
	var doc bytes.Buffer

	doc.WriteString("Test Email")
	auth := smtp.PlainAuth("", emailUser.Username, emailUser.Password, emailUser.EmailServer)

	err := smtp.SendMail(emailUser.EmailServer+":587", // in our case, "smtp.google.com:587"
		auth,
		emailUser.Username,
		[]string{"ryan.christilaw@gmail.com"},
		doc.Bytes())
	if err != nil {
		log.Print("ERROR: attempting to send a mail ", err)
	}
}
