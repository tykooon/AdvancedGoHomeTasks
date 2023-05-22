package main

import (
	"log"
	"net/http"

	"example.com/tykoon/pkg/usersystem"
)

type app struct {
	users     *usersystem.UserSysModel
	log       *log.Logger
	templates *templates
}

func (app *app) notFound(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func (app *app) httpVerbError(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}
