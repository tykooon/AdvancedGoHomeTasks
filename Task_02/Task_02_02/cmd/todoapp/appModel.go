package main

import (
	"errors"
	"os"
	"regexp"

	"example.com/tykoon/pkg/todoplanner"
)

type appModel struct {
	fileName       string                   // name of json-file
	listFlag       bool                     // -list value
	taskToComplete int                      // -complete value
	taskToCreate   string                   // -task value
	todolist       todoplanner.TodoTaskList // list of all tasks
	actionList     []string                 // actions in order from argument string
}

func NewAppModel(fname string) *appModel {
	return &appModel{
		fileName:   fname,
		todolist:   make(todoplanner.TodoTaskList, 0),
		actionList: make([]string, 0),
	}
}

func (app *appModel) CheckArgs() {
	var cmdRe regexp.Regexp = *regexp.MustCompile(`^-task*|^-list*|^-complete*`)
	for _, v := range os.Args[1:] {
		switch cmdRe.FindString(v) {
		case "-list":
			app.actionList = append(app.actionList, "-list")
		case "-task":
			{
				if contains(app.actionList, "-task") {
					FinishWithError("Invalid arguments. Parameter -task duplicate detected.")
				}
				app.actionList = append(app.actionList, "-task")
			}
		case "-complete":
			{
				if contains(app.actionList, "-complete") {
					FinishWithError("Invalid arguments. Parameter -complete duplicate detected")
				}
				app.actionList = append(app.actionList, "-complete")
			}
		default:
			continue
		}
	}
}

func (app *appModel) CheckFile() bool {
	_, err := os.Stat(app.fileName)
	if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		BreakIfError(err, "Sorry. Unexpected I/O error")
	}
	return true
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if str == v {
			return true
		}
	}
	return false
}
