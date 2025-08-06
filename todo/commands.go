package todo

import (
	"fmt"
)

func (tl *TaskList) HandleCommand(args []string) {
	if len(args) < 2 {
		PrintUsage()
		return
	}

	commandArgs := args[1:]
	if !tl.executeCommand(commandArgs) {
		fmt.Printf("Command '%s' is invalid.\n", commandArgs[0])
		PrintUsage()
		return
	}
}

func (tl *TaskList) executeCommand(args []string) bool {
	commandFound := true
	switch args[0] {
	case "add":
		tl.addCommand(args)
	case "a":
		tl.addCommand(args)
	case "delete":
		tl.deleteCommand(args)
	case "d":
		tl.deleteCommand(args)
	case "toggle":
		tl.toggleCommand(args)
	case "t":
		tl.toggleCommand(args)
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

func (tl *TaskList) addCommand(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: todo add \"Task Name\"")
	}

	taskName := argsToTaskName(args)
	tl.addTask(taskName)
}

func (tl *TaskList) deleteCommand(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: todo delete \"Task Name\"")
	}

	taskNameOrIndex := args[1]
	tl.deleteTask(taskNameOrIndex)
}

func (tl *TaskList) toggleCommand(args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: todo toggle \"Task Name\"")
	}

	taskName := argsToTaskName(args)
	tl.toggleTask(taskName)
}

func (tl *TaskList) clearCommand() {
	*tl = TaskList{}
	fmt.Println("Task list cleared.")
}
