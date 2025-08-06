package main

import (
	"fmt"
	"os"

	"github.com/jstn06/cli-todo/todo"
)

const taskFileName string = "tasks.json"

var taskList todo.TaskList

func main() {
	fmt.Println("--------- T O D O ---------")

	if err := taskList.LoadTasks(taskFileName); err != nil {
		fmt.Println("Error while loading tasks:", err)
		os.Exit(1)
	}

	taskList.HandleCommand()

	taskList.SaveTasks(taskFileName)

	fmt.Println("---------------------------")
}
