package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"
	"net/http"

	_ "modernc.org/sqlite"

	"example.com/tykoon/pkg/usersystem"
)

const DbFileName string = "users.db"

func main() {

	db, err := OpenDB(DbFileName)
	if err != nil {
		fmt.Println("Sorry. Failed to open Database... ", err.Error())
		return
	}
	model := usersystem.UserSysModel{DB: db}
	defer model.DB.Close()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	mainHandler := &MainHandler{
		users: &model,
		log:   logger,
	}

	StartServer(mainHandler)
}

func StartServer(m *MainHandler) {
	err := http.ListenAndServe(":5000", m)
	m.log.Fatal(err)
}

func OpenDB(fname string) (db *sql.DB, err error) {
	db, err = sql.Open("sqlite", fname)
	if err == nil {
		err = db.Ping()
		if err == nil {
			return db, err
		}
	}
	return nil, err
}
