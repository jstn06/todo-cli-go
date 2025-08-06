package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

func (tasks TaskList) SaveTasks(filename string) error {
	jsonData, err := json.MarshalIndent(tasks, "", "  ") // MarshalIndent für schön formatierte JSON
	if err != nil {
		return fmt.Errorf("fehler beim marshalling der aufgaben: %w", err)
	}

	err = os.WriteFile(filename, jsonData, 0644) // 0644 sind Standard-Dateirechte (lesbar für alle, schreibbar für Besitzer)
	if err != nil {
		return fmt.Errorf("fehler beim schreiben der aufgaben in die datei: %w", err)
	}

	return nil
}

func (tasks *TaskList) LoadTasks(filename string) error {
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			*tasks = TaskList{}
			return nil
		}
		return fmt.Errorf("fehler beim lesen der aufgaben-Datei: %w", err)
	}

	err = json.Unmarshal(jsonData, tasks)
	if err != nil {
		return fmt.Errorf("fehler beim unmarshalling der aufgaben: %w", err)
	}

	return nil
}
