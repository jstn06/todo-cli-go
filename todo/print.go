package todo

import "fmt"

func PrintUsage() {
	fmt.Println("Usage: todo <command> [arguments]")
	fmt.Println("Commands:")
	fmt.Println("  add|a    <\"Task Name\">           - Adds a new task")
	fmt.Println("  delete|d <index|\"Task Name\">     - Deletes a task by its index")
	fmt.Println("  toggle|t <index|\"Task Name\">     - Toggles the done status of a task by its index")
	fmt.Println("  list|l                           - Lists all tasks")
	fmt.Println("  clear|c                          - Clears all tasks")
}

func PrintSeparator() {
	fmt.Println("---------------------------")
}
