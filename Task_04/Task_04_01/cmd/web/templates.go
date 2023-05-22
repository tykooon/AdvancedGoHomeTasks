package main

import (
	"html/template"
)

type templates struct {
	home  *template.Template
	info  *template.Template
	login *template.Template
}

func initTemplates() *templates {
	res := templates{}
	res.home = template.Must(template.New("home").ParseFiles("template/base.html"))
	res.info = template.Must(template.New("info").ParseFiles("template/base.html", "template/info.html"))
	res.login = template.Must(template.New("login").ParseFiles("template/base.html", "template/login.html"))
	return &res
}

func anonymusTemplate() *template.Template {
	return template.Must(template.New("welcome").ParseFiles("template/base.html", "template/welcome.html"))
}
