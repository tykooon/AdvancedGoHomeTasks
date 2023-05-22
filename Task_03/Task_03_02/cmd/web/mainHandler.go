package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"example.com/tykoon/pkg/usersystem"
)

type MainHandler struct {
	users *usersystem.UserSysModel
	log   *log.Logger
}

func (m *MainHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	urlRe := regexp.MustCompile(`^\/users$|^\/users\/[0-9]+$`)

	if !urlRe.MatchString(request.URL.Path) {
		http.NotFound(writer, request)
		return
	}
	switch request.URL.Path {
	case "/users":
		if request.Method != http.MethodGet {
			http.NotFound(writer, request)
			return
		}
		users, err := m.users.GetAll()
		CheckAndSendJson(writer, users, err)
	default:
		numRe := regexp.MustCompile("[0-9]+")
		id, _ := strconv.ParseInt(numRe.FindString(request.URL.Path), 10, 64)
		switch request.Method {
		case http.MethodGet:
			user, err := m.users.GetById(id)
			CheckAndSendJson(writer, user, err)
		case http.MethodDelete:
			err := m.users.Delete(id)
			var answer = map[string]string{"status": "Deleted"}
			CheckAndSendJson(writer, answer, err)
		default:
			http.NotFound(writer, request)
		}
	}
}

func CheckAndSendJson(writer http.ResponseWriter, obj any, err error) {
	if err != nil {
		sendServerError(writer, err)
	} else {
		err = json.NewEncoder(writer).Encode(obj)
		if err != nil {
			sendServerError(writer, err)
		}
	}
}

func sendServerError(writer http.ResponseWriter, err error) {
	http.Error(writer, err.Error(), http.StatusInternalServerError)
}
