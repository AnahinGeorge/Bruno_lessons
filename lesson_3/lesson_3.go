package lesson_3

import (
	"fmt"
	"time"
)

// Задача уровня Normal

type Task struct {
	id          uint
	Title       string
	Description string
	Status      bool
	Deadline    time.Time
}

type TaskList struct {
	Tasks map[uint]Task
}

func (t *Task) ChangeStatus() {

	t.Status = !t.Status

}

func (l *TaskList) ChangeStatus(id uint) bool {

	if task, ok := l.Tasks[id]; ok {
		task.ChangeStatus()
		l.Tasks[id] = task
		return ok
	}
	return false
}

var counter uint

func NewTask(title string, days int) *Task {

	counter++
	return &Task{
		id:          counter,
		Title:       title,
		Description: "",
		Status:      false,
		Deadline:    time.Now().Add(time.Duration(days) * 24 * time.Hour),
	}
}

func (l *TaskList) Add(task Task) {

	l.Tasks[task.id] = task

}

func (l *TaskList) PrintList() {

	for _, val := range l.Tasks {

		fmt.Println(val.id, val.Title, val.Description, val.Status, val.Deadline.Format(time.DateTime))

	}

}
