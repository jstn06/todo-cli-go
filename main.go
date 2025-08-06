package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jstn06/todo-cli-go/todo"
)

func main() {
	var taskList todo.TaskList

	fmt.Println("--------- T O D O ---------")

	if err := taskList.LoadTasks(); err != nil {
		log.Fatal("Error while loading tasks:", err)
	}

	taskList.HandleCommand(os.Args)

	taskList.SaveTasks()

	todo.PrintSeparator()
}
