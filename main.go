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
	if err := loadTasks(taskFileName); err != nil {
		fmt.Println("Error while loading tasks:", err)
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		taskName := os.Args[2]
		addTask(taskName)
		if err := saveTasks(taskFileName); err != nil {
			fmt.Println("Error while saving tasks:", err)
		}
	case "delete":
		taskName := os.Args[2]
		deleteTask(taskName)
		if err := saveTasks(taskFileName); err != nil {
			fmt.Println("Error while saving tasks:", err)
		}
	case "toggle":
		taskName := os.Args[2]
		toggleDone(taskName)
		if err := saveTasks(taskFileName); err != nil {
			fmt.Println("Error while saving tasks:", err)
		}
	case "list":
		listTasks()
	}
}

func addTask(name string) {
	tasks = append(tasks, Task{Name: name, Done: false})
	fmt.Printf("Task '%s' added.\n", name)
}

func deleteTask(name string) {
	found := false
	for i, t := range tasks {
		if name == t.Name {
			found = true
			fmt.Println("Deleted Task:", tasks[i].Name)
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	if !found {
		fmt.Println("Task not found:", name)
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
