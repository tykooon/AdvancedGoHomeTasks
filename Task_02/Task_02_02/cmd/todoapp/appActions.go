package main

import (
	"fmt"
	"os"

	"example.com/tykoon/pkg/dataaccess"
	"example.com/tykoon/pkg/todoplanner"
)

func (app *appModel) ShowTaskList() {
	if !app.CheckFile() {
		fmt.Println("No tasks found")
		return
	}
	app.MustReadData()
	for _, task := range app.todolist {
		fmt.Println(task.FmtString())
	}
}

func (app *appModel) AddNewTask() {
	if !app.CheckFile() {
		fmt.Println("New Task File will be created")
	} else {
		app.MustReadData()
	}
	newTask := todoplanner.NewTodoTask(len(app.todolist)+1, app.taskToCreate)
	app.todolist = append(app.todolist, *newTask)
	app.MustWriteData()
}

func (app *appModel) CompleteTask() {
	if !app.CheckFile() {
		fmt.Println("No tasks found")
		return
	}
	app.MustReadData()

	if app.taskToComplete > len(app.todolist) || app.taskToComplete < 1 {
		fmt.Printf("No Task with id %d \n", app.taskToComplete)
		return
	}

	if app.todolist[app.taskToComplete-1].IsCompleted {
		fmt.Printf("Task with id %d is already completed\n", app.taskToComplete)
		return
	}
	app.todolist[app.taskToComplete-1].Finish()
	app.MustWriteData()
}

func (app *appModel) MustReadData() {
	list, err := dataaccess.ReadTasksFromJson(app.fileName)
	BreakIfError(err, "Failed to read Tasks File.")
	app.todolist = list
}

func (app *appModel) MustWriteData() {
	err := dataaccess.WriteTasksToJson(app.todolist, app.fileName)
	BreakIfError(err, "Failed to create new Task")
}

func BreakIfError(err error, msg string) {
	if err == nil {
		return
	}
	FinishWithError(msg)
}

func FinishWithError(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
