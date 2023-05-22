package main

import (
	"database/sql"
	"fmt"
	"io"
	"os"

	"example.com/tykoon/pkg/dataio"
	"example.com/tykoon/pkg/usersystem"
	_ "modernc.org/sqlite"
)

const DbFileName string = "users.db"
const ImportFileName string = "persons.txt"
const ErrorsFileName string = "errors.txt"

func main() {

	db, err := CreateAndOpenDB(DbFileName)
	if err != nil {
		fmt.Println("Sorry. Failed to create Database... ", err.Error())
		return
	}
	var model = usersystem.UserSysModel{DB: db}
	defer model.DB.Close()

	var errorWriter io.StringWriter
	errorWriter, err = dataio.NewErrorFWriter(ErrorsFileName)
	if err != nil {
		errorWriter = os.Stdout
	}

	userSource := dataio.NewUsersFReader(ImportFileName, errorWriter)
	userList, _ := userSource.ImportAll()

	n := model.AddRange(userList)
	fmt.Printf("%d users from %d\n successfully added to database", n, len(userList))

	userList, err = model.GetAll()
	if err != nil {
		fmt.Println("Failed to get data from DataBase. ", err.Error())
	} else {
		for _, u := range userList {
			fmt.Println(u.ToString())
		}
	}
	// u2, _ := model.GetById(2)
	// u3, _ := model.GetById(3)
	// fmt.Println(u2.ToString())
	// fmt.Println(u3.ToString())

}

func CreateAndOpenDB(fname string) (db *sql.DB, err error) {
	query := `
	DROP TABLE IF EXISTS Users;

	CREATE TABLE IF NOT EXISTS Users(
		Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		FullName TEXT NOT NULL,
		Email TEXT UNIQUE NOT NULL,
		Hash TEXT NOT NULL,
		IsActive INTEGER NOT NULL)`

	file, err := os.Create(fname)
	if err == nil {
		file.Close()
		db, err = sql.Open("sqlite", fname)
		if err == nil {
			_, err = db.Exec(query)
		}
	}
	return db, err
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
