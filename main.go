package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	todoTools "github.com/jstn06/cli-todo/tools"
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
		todoTools.PrintUsage()
		fmt.Println("---------------------------")
		return
	}

	command := os.Args[1]
	if !executeCommnd(command) {
		fmt.Printf("Command '%s' is invalid.\n", command)
		todoTools.PrintUsage()
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
	addTask()
}

func deleteCommand() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: todo delete \"Task Name\"")
	}
	deleteTask()
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

func addTask() {
	argList := []string{}
	for i := 2; i < len(os.Args); i++ {
		argList = append(argList, os.Args[i])
	}
	taskName := strings.Join(argList, " ")

	tasks = append(tasks, Task{Name: taskName, Done: false})
	fmt.Printf("Task '%s' added.\n", taskName)
}

func deleteTask() {
	taskNameOrIndex := os.Args[2]

	deleted, name := deleteTaskByIndex(taskNameOrIndex)
	if !deleted {
		deleted, name = deleteTaskByName(taskNameOrIndex)
	}

	if deleted {
		fmt.Printf("Task '%s' deleted.\n", name)
	} else {
		fmt.Printf("Task '%s' not found.\n", taskNameOrIndex)
	}
}

func deleteTaskByIndex(taskNameOrIndex string) (bool, string) {
	deleted := false
	name := ""

	findTaskByIndex(taskNameOrIndex, func(taskIndex int) {
		name = tasks[taskIndex].Name
		tasks = append(tasks[:taskIndex], tasks[taskIndex+1:]...)
		deleted = true
	})

	return deleted, name
}

func findTaskByIndex(index string, op func(taskIndex int)) {
	if len(os.Args) == 3 {
		taskIndex, err := strconv.Atoi(index)
		taskIndex--
		if err == nil && taskIndex >= 0 && taskIndex < len(tasks) {
			op(taskIndex)
		}
	}
}

func deleteTaskByName(taskName string) (bool, string) {
	deleted := false
	name := ""

	findTaskByName(taskName, func(taskIndex int) {
		name = tasks[taskIndex].Name
		tasks = append(tasks[:taskIndex], tasks[taskIndex+1:]...)
		deleted = true
	})

	return deleted, name
}

func findTaskByName(name string, op func(taskIndex int)) {
	for i, t := range tasks {
		if name == t.Name {
			op(i)
			break
		}
	}
}

func toggleDone(name string) {
	found := false
	for i, t := range tasks {
		if name == t.Name {
			found = true
			if tasks[i].Done {
				fmt.Printf("Task '%s' unchecked.\n", name)
			} else {
				fmt.Printf("Task '%s' checked.\n", name)
			}
			tasks[i].Done = !tasks[i].Done
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

	for i, task := range tasks {
		printFormattedTask(task, i+1)
	}
}

func printFormattedTask(task Task, index int) {
	doneMarker := " "
	if task.Done {
		doneMarker = "X"
	}

	fmt.Printf("[%s]: %d. %s\n", doneMarker, index, task.Name)
}

func saveTasks(filename string) error {
	jsonData, err := json.MarshalIndent(tasks, "", "  ") // MarshalIndent für schön formatierte JSON
	if err != nil {
		return fmt.Errorf("fehler beim Marshalling der Aufgaben: %w", err)
	}

	err = os.WriteFile(filename, jsonData, 0644) // 0644 sind Standard-Dateirechte (lesbar für alle, schreibbar für Besitzer)
	if err != nil {
		return fmt.Errorf("fehler beim schreiben der aufgaben in die datei: %w", err)
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
		return fmt.Errorf("fehler beim lesen der aufgaben-Datei: %w", err)
	}

	err = json.Unmarshal(jsonData, &tasks)
	if err != nil {
		return fmt.Errorf("fehler beim unmarshalling der aufgaben: %w", err)
	}

	return nil
}
