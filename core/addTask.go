package core

type OptionsImpl struct {
}

func NewOptionsImpl() OptionsImpl {
	return OptionsImpl{}
}

// Add add a task
func (o OptionsImpl) Add(task Task, tasks map[string]Task) map[string]Task {
	//Adding the task
	tasks[task.ID] = task

	return tasks
}
