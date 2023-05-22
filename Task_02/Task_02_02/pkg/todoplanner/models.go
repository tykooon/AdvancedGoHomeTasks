package todoplanner

import (
	"fmt"
	"time"
)

type TodoTask struct {
	Id          int
	Name        string
	IsCompleted bool
	Created     time.Time
	Finished    time.Time
}

type TodoTaskList []TodoTask

func NewTodoTask(no int, name string) *TodoTask {
	res := &TodoTask{
		Id:          no,
		Name:        name,
		IsCompleted: false,
		Created:     time.Now(),
	}
	return res
}

func (t *TodoTask) Finish() {
	t.IsCompleted = true
	t.Finished = time.Now()
}

func (task *TodoTask) FmtString() (result string) {
	var status = statusName(task.IsCompleted)
	var startTime = task.Created.Format(time.RFC822Z)
	result = fmt.Sprintf("Task %3d | %s\n------- Status: %6s | Started: %s ",
		task.Id, task.Name, status, startTime)
	if !task.Finished.IsZero() {
		result += fmt.Sprint("| Finished: ", task.Finished.Format(time.RFC822Z))
	}
	return result
}  
    
func statusName(s bool) string {
	if s {
		return "Done"
	}
	return "Undone"
}
