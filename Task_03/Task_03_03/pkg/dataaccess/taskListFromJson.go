package dataaccess

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	todo "example.com/tykoon/pkg/todoplanner"
)

func ReadTasksFromJson(fileName string) (r todo.TodoTaskList, err error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &r)
	return
}

func WriteTasksToJson(tasks todo.TodoTaskList, fileName string) (err error) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("failed to open")
		return
	}
	defer file.Close()

	data, err := json.Marshal(tasks)
	if err != nil {
		return
	}

	_, err = file.Write(data)
	return
}
