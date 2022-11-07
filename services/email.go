package services

import (
	// "log"
	"net/http"
	"server/utils"

	mail "github.com/xhit/go-simple-mail/v2"
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
	server.Port = 587
	server.Username = "dragon12xw@gmail.com"
	server.Password = "wolfW12300"
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		// log.Fatal(err)
		utils.BadResponse(w, utils.RespBad{
			Message:    err.Error(),
			StatusCode: http.StatusForbidden,
		})
		return
	}

	email := mail.NewMSG()
	email.SetFrom("From Me <dragon12xw@gmail.com>")
	email.AddTo("desarrollo3@aveonline.co")
	// email.AddCc("another_you@example.com")
	email.SetSubject("New Go Email")

	email.SetBody(mail.TextHTML, htmlBody)
	// email.AddAttachment("super_cool_file.png")

	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		utils.BadResponse(w, utils.RespBad{
			Message:    err.Error(),
			StatusCode: http.StatusForbidden,
		})
		return
	}
}
