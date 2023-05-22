package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"example.com/tykoon/pkg/dataaccess"
)

//const templBase string = "template/base.html"
//const templList string = "template/list.html"

type MainHandler struct {
	model *webModel
	log   *log.Logger
}

func (m *MainHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/list":
		m.ShowTodoList(writer)
	case "/":
		http.Redirect(writer, request, "/list", http.StatusTemporaryRedirect)
	default:
		http.NotFound(writer, request)
	}
}

func (m *MainHandler) ShowTodoList(writer http.ResponseWriter) {
	if m.model.CheckFile() {
		m.model.MustReadData()
	}

	listView := template.New("listView")
	fmap := template.FuncMap(map[string]any{
		"timeformat": timeformat,
		"statusconv": statusconvert,
	})

	listView = listView.Funcs(fmap)

	listView = template.Must(listView.ParseFiles("template/base.html", "template/list.html"))
	listView.ExecuteTemplate(writer, "base.html", m.model.todoList)
}

func (web *webModel) CheckFile() bool {
	_, err := os.Stat(web.fileName)
	if errors.Is(err, os.ErrNotExist) {
		return false
	} else if err != nil {
		log.Fatal("Sorry. Unexpected I/O error", err.Error())
		os.Exit(1)
	}
	return true
}

func timeformat(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("02 Jan 2006 15:04")
}

func statusconvert(s bool) template.HTML {
	if s {
		return "<span class='text-success fw-semibold'>Done</span>"
	}
	return "<span class='text-danger fw-semibold'>Undone</span>"
}

func (web *webModel) MustReadData() {
	list, err := dataaccess.ReadTasksFromJson(FileName)
	if err != nil {
		log.Fatal("Failed to read Tasks File.")
		os.Exit(1)
	}
	web.todoList = list
}
