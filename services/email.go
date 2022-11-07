package services

import (
	"net/http"
	"net/smtp"
	"os"
	"server/utils"
)

const (
	user = "dragon12xw@gmail.com"
	pass = "wolfW12300"
	to   = "desarrollo3@aveonline.co"
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
	// from := "dragon12xw@gmail.com"
	from := os.Getenv("FROM_EMAIL")
	password := os.Getenv("PASS_EMAIL")
	toMail := "desarrollo3@aveonline.co"
	to := []string{toMail}

	host := "smtp.gmail.com"
	port := "587"
	addres := host + ":" + port
	subject := "Asunto de prueba -noreplay"
	body := htmlBody
	message := []byte(subject + body)
	auth := smtp.PlainAuth("", user, password, host)
	err := smtp.SendMail(addres, auth, from, to, message)
	if err != nil {
		utils.BadResponse(w, utils.RespBad{
			Message:    "Error Connect: " + err.Error(),
			StatusCode: http.StatusForbidden,
		})
		return
	}
	utils.SendResponse(w, utils.RespOk{
		Message: "Correo enviado",
	})
}
