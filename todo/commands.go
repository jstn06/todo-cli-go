package todo

import (
	"fmt"
	"os"
)

func (tl *TaskList) HandleCommand() {
	if len(os.Args) < 2 {
		PrintUsage()
		return
	}

	command := os.Args[1]
	if !tl.executeCommnd(command) {
		fmt.Printf("Command '%s' is invalid.\n", command)
		PrintUsage()
		return
	}
}

func (tl *TaskList) executeCommnd(command string) bool {
	commandFound := true
	switch command {
	case "add":
		tl.addCommand()
	case "a":
		tl.addCommand()
	case "delete":
		tl.deleteCommand()
	case "d":
		tl.deleteCommand()
	case "toggle":
		tl.toggleCommand()
	case "t":
		tl.toggleCommand()
	case "clear":
		tl.clearCommand()
	case "c":
		tl.clearCommand()
	case "list":
		tl.listTasks()
	case "l":
		tl.listTasks()
	default:
		commandFound = false
	}

	return commandFound
}

func (tl *TaskList) addCommand() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: todo add \"Task Name\"")
	}

	taskName := argsToTaskName()
	tl.addTask(taskName)
}

func (tl *TaskList) deleteCommand() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: todo delete \"Task Name\"")
	}

	taskNameOrIndex := os.Args[2]
	tl.deleteTask(taskNameOrIndex)
}

func (tl *TaskList) toggleCommand() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: todo toggle \"Task Name\"")
	}

	taskName := argsToTaskName()
	tl.toggleTask(taskName)
}

func (tl *TaskList) clearCommand() {
	*tl = TaskList{}
	fmt.Println("Task list cleared.")
}
