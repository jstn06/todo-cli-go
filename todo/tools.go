package todo

import (
	"fmt"
	"strconv"
)

func (tl TaskList) findTask(taskNameOrIndex string) (int, error) {
	index, err := tl.findTaskByHumanIndex(taskNameOrIndex)

	if err == nil {
		return index, nil
	}

	return tl.findTaskByName(taskNameOrIndex)
}

func (tl TaskList) findTaskByHumanIndex(taskIndex string) (int, error) {
	index, err := strconv.Atoi(taskIndex)
	if err != nil {
		return -1, fmt.Errorf("'%s' is not a valid number", taskIndex)
	}

	index--
	if index < 0 || index >= len(tl) {
		return -1, fmt.Errorf("index out of range")
	}

	return index, nil
}

func (tl TaskList) findTaskByName(name string) (int, error) {
	for index, task := range tl {
		if name == task.Name {
			return index, nil
		}
	}

	return -1, fmt.Errorf("task called '%s' not found", name)
}
