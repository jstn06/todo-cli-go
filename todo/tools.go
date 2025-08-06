package todo

import (
	"os"
	"strconv"
)

func (tl TaskList) findTaskByIndex(index string, op func(taskIndex int)) {
	if len(os.Args) == 3 {
		taskIndex, err := strconv.Atoi(index)
		taskIndex--
		if err == nil && taskIndex >= 0 && taskIndex < len(tl) {
			op(taskIndex)
		}
	}
}

func (tl TaskList) findTaskByName(name string, op func(taskIndex int)) {
	for i, t := range tl {
		if name == t.Name {
			op(i)
			break
		}
	}
}
