package core

import "fmt"

// Delete delete one task by ID
func (o OptionsImpl) Delete(ID string, tasks map[string]Task) map[string]Task {

	if ID == "" {
		fmt.Println("enter the id of the task: ")
		_, _ = fmt.Scan(&ID)
	}
	//Delete the task
	delete(tasks, ID)

	return tasks
}
