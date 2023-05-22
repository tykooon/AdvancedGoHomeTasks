package main

import (
	"log"
	"os"

	"database/sql"
	"net/http"

	_ "modernc.org/sqlite"

	"example.com/tykoon/pkg/usersystem"
)

const DbFileName string = "sql/users.db"

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := OpenDB(DbFileName)
	if err != nil {
		logger.Fatal("Sorry. Failed to open Database... ", err.Error())
		return
	}
	defer db.Close()

	app := app{
		users:     &usersystem.UserSysModel{DB: db},
		log:       logger,
		templates: initTemplates(),
	}

	err = http.ListenAndServe(":5000", app.routes())
	app.log.Fatal(err)
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
