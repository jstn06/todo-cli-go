package todo

import (
	"strconv"
)

func (tl TaskList) findTaskByHumanIndex(index string) (int, bool) {
	taskIndex, err := strconv.Atoi(index)
	taskIndex--
	if err == nil && taskIndex >= 0 && taskIndex < len(tl) {
		return taskIndex, true
	}

	return -1, false
}

func (tl TaskList) findTaskByName(name string) (int, bool) {
	for i, t := range tl {
		if name == t.Name {
			return i, true
		}
	}

	return -1, false
}
