package main

import (
	"gopkg.in/gomail.v2"
)

func SendEmail() error {
	message := gomail.NewMessage()
	message.SetHeader("From", "rizkyian78")
	message.SetHeader("To", "rizkyian78")
	message.SetHeader("Subject", "rizkyian78")
	message.SetBody("text/html", "rizkyian78")

	return nil
}
