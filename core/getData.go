package core

import (
	"fmt"
	"strconv"
)

// GetData get the information for the tasks
func (o OptionsImpl) GetData(transaction string) Task {
	var id, name, desc, status, dateInit, dateEnd, priority string

	//if it's a create transaction
	if transaction == "C" {
		fmt.Println("Enter the ID: ")
		_, _ = fmt.Scan(&id)
	}

	fmt.Println("Enter the Name: ")
	_, _ = fmt.Scan(&name)

	fmt.Println("Enter the Description: ")
	_, _ = fmt.Scan(&desc)

	fmt.Println("Enter the Status: ")
	_, _ = fmt.Scan(&status)

	fmt.Println("Enter the Init Date: ")
	_, _ = fmt.Scan(&dateInit)

	fmt.Println("Enter the End Date: ")
	_, _ = fmt.Scan(&dateEnd)

	fmt.Println("Enter the Priority: ")
	_, _ = fmt.Scan(&priority)

	statusInt, _ := strconv.Atoi(status)
	priorityInt, _ := strconv.Atoi(priority)
	task := NewTask(id, name, desc, dateInit, dateEnd, statusInt, priorityInt)

	return task
}
