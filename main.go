package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jstn06/cli-todo/todo"
)

func main() {
	var taskList todo.TaskList

	fmt.Println("--------- T O D O ---------")

	if err := taskList.LoadTasks(); err != nil {
		log.Fatal("Error while loading tasks:", err)
	}

	taskList.HandleCommand(os.Args)

	taskList.SaveTasks()

	fmt.Println("---------------------------")
}
