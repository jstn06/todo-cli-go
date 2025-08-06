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
	err := tl.dispatchCommand(commandArgs)

	if err != nil {
		fmt.Printf("Error: %v\n\n", err)
		PrintUsage()
	}
}

func (tl *TaskList) dispatchCommand(args []string) error {
	command := args[0]
	commandArgs := args[1:]

	switch args[0] {
	case "add", "a":
		return tl.addCommand(commandArgs)
	case "delete", "d":
		return tl.deleteCommand(commandArgs)
	case "toggle", "t":
		return tl.toggleCommand(commandArgs)
	case "clear", "c":
		return tl.clearCommand()
	case "list", "l":
		return tl.listTasks()
	default:
		return fmt.Errorf("invalid command '%s'", command)
	}
}

func (tl *TaskList) addCommand(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("missing task name for 'add'")
	}

	tl.addTask(args)
	return nil
}

func (tl *TaskList) deleteCommand(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("missing task name for 'delete'")
	}

	tl.deleteTask(args)
	return nil
}

func (tl *TaskList) toggleCommand(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("missing task name for 'toggle'")
	}

	tl.toggleTask(args)
	return nil
}

func (tl *TaskList) clearCommand() error {
	*tl = TaskList{}
	fmt.Println("Task list cleared.")
	return nil
}
