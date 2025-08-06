package todo

import (
	"fmt"
	"strconv"
)

func (tl TaskList) findTaskByHumanIndex(index string) (int, error) {
	taskIndex, err := strconv.Atoi(index)
	if err != nil {
		return -1, fmt.Errorf("'%s' is not a valid number", index)
	}

	taskIndex--
	if taskIndex < 0 || taskIndex >= len(tl) {
		return -1, fmt.Errorf("index out of range")
	}

	return taskIndex, nil
}

func (tl TaskList) findTaskByName(name string) (int, error) {
	for i, t := range tl {
		if name == t.Name {
			return i, nil
		}
	}

	return -1, fmt.Errorf("task called '%s' not found", name)
}
