package todo

import "fmt"

type Task struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

type TaskList []Task

func (tasks TaskList) listTasks() error {
	if len(tasks) == 0 {
		fmt.Println("Task list is empty.")
	}

	for i, task := range tasks {
		printFormattedTask(task, i+1)
	}

	return nil
}

func printFormattedTask(task Task, index int) {
	doneMarker := " "
	if task.Done {
		doneMarker = "X"
	}

	fmt.Printf("[%s]: %d. %s\n", doneMarker, index, task.Name)
}
