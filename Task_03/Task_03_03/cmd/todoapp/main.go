package main

import (
	"flag"
)

const FileName string = "todolist.json"

func main() {

	var app appModel = *NewAppModel(FileName)

	flag.BoolVar(&app.listFlag, "list", false, "show list of all tasks (optional)")
	flag.IntVar(&app.taskToComplete, "complete", 0, "id-number of task to finish now")
	flag.StringVar(&app.taskToCreate, "task", "", "add name of new task to start now")

	flag.Parse()
	app.CheckArgs()

	if len(app.actionList) == 0 {
		flag.Usage()
		return
	}

	for _, v := range app.actionList {
		switch v {
		case "-list":
			app.ShowTaskList()
		case "-task":
			app.AddNewTask()
		case "-complete":
			app.CompleteTask()
		default:
			continue
		}
	}
}
