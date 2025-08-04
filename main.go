package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

const taskFileName string = "tasks.json"

var tasks []Task = []Task{}

func main() {
	fmt.Println("--------- T O D O ---------")

	if err := loadTasks(taskFileName); err != nil {
		fmt.Println("Error while loading tasks:", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		printUsage()
		fmt.Println("---------------------------")
		return
	}

	command := os.Args[1]
	if !executeCommnd(command) {
		fmt.Printf("Command '%s' is invalid.\n", command)
		printUsage()
	}

	if command != "list" && command != "l" {
		if err := saveTasks(taskFileName); err != nil {
			fmt.Println("Error while saving tasks:", err)
		}
	}

	fmt.Println("---------------------------")
}

func executeCommnd(command string) bool {
	commandFound := true
	switch command {
	case "add":
		addCommand()
	case "a":
		addCommand()
	case "delete":
		deleteCommand()
	case "d":
		deleteCommand()
	case "toggle":
		toggleCommand()
	case "t":
		toggleCommand()
	case "list":
		listTasks()
	case "clear":
		clearCommand()
	case "c":
		clearCommand()
	case "l":
		listTasks()
	default:
		commandFound = false
	}

	return commandFound
}

func addCommand() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: todo add \"Task Name\"")
	}
	taskName := os.Args[2]
	tasks = append(tasks, Task{Name: taskName, Done: false})
	fmt.Printf("Task '%s' added.\n", taskName)
}

func deleteCommand() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: todo delete \"Task Name\"")
	}
	taskName := os.Args[2]
	deleteTask(taskName)
}

func toggleCommand() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: todo toggle \"Task Name\"")
	}
	taskName := os.Args[2]
	toggleDone(taskName)
}

func clearCommand() {
	tasks = []Task{}
	fmt.Println("Task list cleared.")
}

func deleteTask(name string) {
	found := false
	for i, t := range tasks {
		if name == t.Name {
			found = true
			fmt.Printf("Task '%s' deleted.\n", tasks[i].Name)
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	if !found {
		fmt.Printf("Task '%s' not found.\n", name)
	}
}

func toggleDone(name string) {
	found := false
	for i, t := range tasks {
		if name == t.Name {
			found = true
			tasks[i].Done = !tasks[i].Done
			fmt.Println("Updated Task Done State:")
			printFormattedTask(tasks[i])
			break
		}
	}

	if !found {
		fmt.Println("Task not found:", name)
	}
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("Task list is empty.")
	}

	for _, task := range tasks {
		printFormattedTask(task)
	}
}

func printFormattedTask(task Task) {
	doneMarker := " "
	if task.Done {
		doneMarker = "X"
	}

	fmt.Printf("[%s]: %s\n", doneMarker, task.Name)
}

func saveTasks(filename string) error {
	jsonData, err := json.MarshalIndent(tasks, "", "  ") // MarshalIndent für schön formatierte JSON
	if err != nil {
		return fmt.Errorf("Fehler beim Marshalling der Aufgaben: %w", err)
	}

	err = os.WriteFile(filename, jsonData, 0644) // 0644 sind Standard-Dateirechte (lesbar für alle, schreibbar für Besitzer)
	if err != nil {
		return fmt.Errorf("Fehler beim Schreiben der Aufgaben in die Datei: %w", err)
	}

	return nil
}

func loadTasks(filename string) error {
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			tasks = []Task{}
			return nil
		}
		return fmt.Errorf("Fehler beim Lesen der Aufgaben-Datei: %w", err)
	}

	err = json.Unmarshal(jsonData, &tasks)
	if err != nil {
		return fmt.Errorf("Fehler beim Unmarshalling der Aufgaben: %w", err)
	}

	return nil
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("-> todo add 'name of your task'")
	fmt.Println("-> todo delete 'name of your task'")
	fmt.Println("-> todo toggle 'name of your task'")
	fmt.Println("-> todo list")
}
