package main

import (
	"log"
	"net/http"
	"os"
)

const FileName string = "todolist.json"
const cert string = "cert/localhost.cer"
const pkey string = "cert/localhost.pkey"

func main() {

	model := NewWebModel(FileName)

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	mainHandler := &MainHandler{
		model: model,
		log:   logger,
	}

	//go StartServer(mainHandler)

	StartServerTLS(mainHandler)
}

func StartServerTLS(m *MainHandler) {
	err := http.ListenAndServeTLS(":5500", cert, pkey, m)
	m.log.Fatal(err)
}

func StartServer(m *MainHandler) {
	err := http.ListenAndServe(":5000", m)
	m.log.Fatal(err)
}
