package core

import "fmt"

// Get Retrieve one task by ID
func (o OptionsImpl) Get(ID string, tasks map[string]Task) Task {
	//Retrieve the task
	task := tasks[ID]

	fmt.Println("ID | Nombre | Descripcion | Estado | Fecha Inicio | Fecha Fin | Prioridad")
	fmt.Printf("%v  %v  %v  %v  %v  %v  %v ", task.ID, task.Nombre, task.Descripcion, task.Estado, task.FechaIni, task.FechaFin, task.Prioridad)

	return task
}
