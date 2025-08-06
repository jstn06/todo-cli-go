package todo

import (
	"fmt"
	"strings"
)

func (tl *TaskList) addTask(args []string) {
	taskName := strings.Join(args, " ")
	*tl = append(*tl, Task{Name: taskName, Done: false})
	fmt.Printf("Task '%s' added.\n", taskName)
}

func (tl *TaskList) deleteTask(args []string) {
	taskNameOrIndex := strings.Join(args, " ")

	index, err := tl.findTaskByHumanIndex(taskNameOrIndex)
	if err != nil {
		index, err = tl.findTaskByName(taskNameOrIndex)
	}

	if err != nil {
		fmt.Printf("Task '%s' not found.\n", taskNameOrIndex)
	} else {
		name := (*tl)[index].Name
		*tl = append((*tl)[:index], (*tl)[index+1:]...)
		fmt.Printf("Task '%s' deleted.\n", name)
	}
}

func (tl *TaskList) toggleTask(args []string) {
	taskNameOrIndex := strings.Join(args, " ")

	index, err := tl.findTaskByHumanIndex(taskNameOrIndex)
	if err != nil {
		index, err = tl.findTaskByName(taskNameOrIndex)
	}

	if err != nil {
		fmt.Printf("Task '%s' not found.\n", taskNameOrIndex)
	} else {
		if (*tl)[index].Done {
			fmt.Printf("Task '%s' unchecked.\n", taskNameOrIndex)
		} else {
			fmt.Printf("Task '%s' checked.\n", taskNameOrIndex)
		}
		(*tl)[index].Done = !(*tl)[index].Done
	}
}
