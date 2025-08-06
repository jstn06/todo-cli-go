package todo

import (
	"fmt"
)

func (tl *TaskList) addTask(name string) {
	*tl = append(*tl, Task{Name: name, Done: false})
	fmt.Printf("Task '%s' added.\n", name)
}

func (tl *TaskList) deleteTask(taskNameOrIndex string) {
	deleted, name := tl.deleteTaskByIndex(taskNameOrIndex)
	if !deleted {
		deleted, name = tl.deleteTaskByName(taskNameOrIndex)
	}

	if deleted {
		fmt.Printf("Task '%s' deleted.\n", name)
	} else {
		fmt.Printf("Task '%s' not found.\n", taskNameOrIndex)
	}
}

func (tl *TaskList) deleteTaskByIndex(taskNameOrIndex string) (bool, string) {
	deleted := false
	name := ""

	tl.findTaskByIndex(taskNameOrIndex, func(taskIndex int) {
		name = (*tl)[taskIndex].Name
		*tl = append((*tl)[:taskIndex], (*tl)[taskIndex+1:]...)
		deleted = true
	})

	return deleted, name
}

func (tl *TaskList) deleteTaskByName(taskName string) (bool, string) {
	deleted := false
	name := ""

	tl.findTaskByName(taskName, func(taskIndex int) {
		name = (*tl)[taskIndex].Name
		*tl = append((*tl)[:taskIndex], (*tl)[taskIndex+1:]...)
		deleted = true
	})

	return deleted, name
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
