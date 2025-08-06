package todo

import (
	"fmt"
	"strings"
)

func (tl *TaskList) addTask(name string) {
	*tl = append(*tl, Task{Name: name, Done: false})
	fmt.Printf("Task '%s' added.\n", name)
}

func (tl *TaskList) deleteTask(args []string) {
	taskNameOrIndex := strings.Join(args, " ")

	index, found := tl.findTaskByHumanIndex(taskNameOrIndex)
	if !found {
		index, found = tl.findTaskByName(taskNameOrIndex)
	}

	if found {
		name := (*tl)[index].Name
		*tl = append((*tl)[:index], (*tl)[index+1:]...)
		fmt.Printf("Task '%s' deleted.\n", name)
	} else {
		fmt.Printf("Task '%s' not found.\n", taskNameOrIndex)
	}
}

func (tl *TaskList) toggleTask(name string) {
	found := false
	for i, t := range *tl {
		if name != t.Name {
			continue
		}

		found = true
		if (*tl)[i].Done {
			fmt.Printf("Task '%s' unchecked.\n", name)
		} else {
			fmt.Printf("Task '%s' checked.\n", name)
		}
		(*tl)[i].Done = !(*tl)[i].Done
		break
	}

	if !found {
		fmt.Println("Task not found:", name)
	}
}
