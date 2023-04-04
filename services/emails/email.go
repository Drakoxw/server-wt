package services

import (
	"fmt"
	"os"

	mail "github.com/xhit/go-simple-mail/v2"
)

func createServer() (*mail.SMTPClient, error) {
	server := mail.NewSMTPClient()
	server.Host = "smtp.gmail.com"
	server.Port = 587
	server.Username = os.Getenv("FROM_EMAIL")
	server.Password = os.Getenv("PASS_EMAIL")
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		return smtpClient, err
	}
	return smtpClient, nil
}

func SendEmailPlantilla(subject string, emailCli string, template string) bool {
	smtpClient, err := createServer()
	if err != nil {
		fmt.Println("Error createServer : " + err.Error())
		return false
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom("we Tourism <dragon12xw@gmail.com>")
	email.AddTo(emailCli)
	// email.AddCc("another_you@example.com")
	email.SetSubject(subject)

	email.SetBody(mail.TextHTML, template)
	// add adjuntar archivo
	// email.AddAttachment("static/img/DrakoImg.png")

	err = email.Send(smtpClient)
	if err != nil {
		fmt.Println("Error emailSend : " + err.Error())
		return false
	} else {
		return true
	}
}
