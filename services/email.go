package services

import (
	// "log"
	"net/http"
	"server/utils"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

const (
	user = "dragon12xw@gmail.com"
	pass = "wolfW12300"
)

var htmlBody = `
<html>
<head>
   <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
   <title>Hello, World</title>
</head>
<body>
   <p>This is an email using Go</p>
</body>
`

func SendEmail(w http.ResponseWriter) {
	server := mail.NewSMTPClient()
	server.Host = "smtp.gmail.com"
	server.Port = 465
	server.Username = user
	server.Password = pass
	server.Encryption = mail.EncryptionSTARTTLS

	server.KeepAlive = false

	// Timeout for connect to SMTP Server
	server.ConnectTimeout = 25 * time.Second
	server.SendTimeout = 25 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		// log.Fatal(err)
		utils.BadResponse(w, utils.RespBad{
			Message:    "Error Connect:" + err.Error(),
			StatusCode: http.StatusForbidden,
		})
		return
	}

	email := mail.NewMSG()
	email.SetFrom("From Me <" + user + ">")
	email.AddTo("desarrollo3@aveonline.co")
	// email.AddCc("another_you@example.com")
	email.SetSubject("New Go Email")

	email.SetBody(mail.TextHTML, htmlBody)
	// email.AddAttachment("super_cool_file.png")

	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		utils.SendResponse(w, utils.RespOk{
			Message: "Correo enviado",
		})
		return
	}
}
