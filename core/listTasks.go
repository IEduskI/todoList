package core

import (
	"fmt"
	"strconv"
	"todoList/kit"
)

// List list all tasks
func (o OptionsImpl) List(tasks map[string]Task) {
	//List all tasks
	var inProgress int

	status, priority := kit.RetrieveOptions()

	//Para hacer el cambio de "numero" a descripcion en los campos de estado y prioridad
	//pensaba incluir las opciones en un YAML y hacer la conversion en base a la key
	//que vendria siendo el id de cada opcion

	fmt.Println("ID | Nombre | Descripcion | Estado | Fecha Inicio | Fecha Fin | Prioridad")
	for _, task := range tasks {
		fmt.Println("")
		fmt.Printf("%v  %v  %v  %v  %v  %v  %v ", task.ID, task.Nombre, task.Descripcion, status[task.Estado], task.FechaIni, task.FechaFin, priority[task.Prioridad])
		fmt.Println("")
		if task.Estado == 1 {
			inProgress++
		}
	}

	fmt.Println("")
	fmt.Println("Tiene " + strconv.Itoa(inProgress) + " actividades pendientes de realizar")
}
