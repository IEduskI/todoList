package core

import "fmt"

// Update updates a task
func (o OptionsImpl) Update(tasks map[string]Task) map[string]Task {

	var ID string
	fmt.Println("enter the id of the task: ")
	_, _ = fmt.Scan(&ID)

	//Retrieve the task
	task := o.Get(ID, tasks)

	fmt.Println("Insert the new information")
	fmt.Println("===========================")

	updatedTask := o.GetData("U")

	updatedTask.ID = task.ID

	newTasks := o.Delete(ID, tasks)
	updatedTasks := o.Add(updatedTask, newTasks)

	return updatedTasks
}
