package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const taskDirName string = ".todo"
const taskFileName string = "tasks.json"

func getTaskFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not find home directory: %w", err)
	}

	tasksDir := filepath.Join(homeDir, taskDirName)

	err = os.MkdirAll(tasksDir, 0755)
	if err != nil {
		return "", fmt.Errorf("could not creat .todo directory: %w", err)
	}

	return filepath.Join(tasksDir, taskFileName), nil
}

func (tasks TaskList) SaveTasks() error {
	filePath, err := getTaskFilePath()
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("error while marshalling tasks: %w", err)
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error while writing tasks into file: %w", err)
	}

	return nil
}

func (tasks *TaskList) LoadTasks() error {
	filePath, err := getTaskFilePath()
	if err != nil {
		return err
	}

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			*tasks = TaskList{}
			return nil
		}
		return fmt.Errorf("error while reading tasks file: %w", err)
	}

	err = json.Unmarshal(jsonData, tasks)
	if err != nil {
		return fmt.Errorf("error while unmarshalling tasks: %w", err)
	}

	return nil
}
