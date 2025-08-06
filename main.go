package main

import (
	"fmt"
	"os"

	"github.com/jstn06/cli-todo/todo"
)

func main() {
	var taskList todo.TaskList

	fmt.Println("--------- T O D O ---------")

	if err := taskList.LoadTasks(); err != nil {
		fmt.Println("Error while loading tasks:", err)
		os.Exit(1)
	}

	taskList.HandleCommand()

	taskList.SaveTasks()

	fmt.Println("---------------------------")
}
