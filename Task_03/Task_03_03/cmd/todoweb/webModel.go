package main

import (
	"example.com/tykoon/pkg/todoplanner"
)

type webModel struct {
	fileName string
	todoList todoplanner.TodoTaskList
}

func NewWebModel(fname string) *webModel {
	return &webModel{
		fileName: fname,
	}
}
