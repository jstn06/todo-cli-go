package tools

import (
	"fmt"
)

func PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("-> todo add 'name of your task'")
	fmt.Println("-> todo delete 'name of your task'")
	fmt.Println("-> todo toggle 'name of your task'")
	fmt.Println("-> todo list")
}
