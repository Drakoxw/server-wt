package services

import (
	"html/template"
	"strings"
)

func SendWelcomeProvider(Provider string, email string, id string, enterprise string) {
	link := "http://drako-domain/providers/validateEmail?id=" + id
	templateFile, _ := template.ParseFiles("static/template/welcomeProvider.html")
	templateString := strings.Replace(templateFile.Tree.Root.String(), "{{.Provider}}", Provider, 1)
	templateString = strings.Replace(templateString, "{{.Link}}", link, 1)
	templateString = strings.Replace(templateString, "{{.Enterprise}}", enterprise, 1)
	SendEmailPlantilla("Correo Bienvenida -noreplay", email, templateString)
}

func SendWelcomeUser(Client string, email string, id string) {
	link := "http://drako-domain/users/validateEmail?id=" + id
	templateFile, _ := template.ParseFiles("static/template/welcomeUser.html")
	templateString := strings.Replace(templateFile.Tree.Root.String(), "{{.cliente}}", Client, 1)
	templateString = strings.Replace(templateString, "{{.Link}}", link, 1)
	SendEmailPlantilla("Correo Bienvenida -noreplay", email, templateString)
}
