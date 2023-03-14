package main

import (
	"fmt"
	"todoList/core"
)

func main() {
	menu := true
	var option string
	tasks := make(map[string]core.Task)
	coreImpl := core.NewOptionsImpl()
	for menu {
		fmt.Println("======= Options Menu ===========")
		fmt.Println("1. Agregar tarea")
		fmt.Println("2. Listar Tareas")
		fmt.Println("3. Modificar tarea")
		fmt.Println("4. Eliminar tarea")
		fmt.Println("5. Salir")
		fmt.Println("=================================")
		fmt.Println("")

		fmt.Println("Choose an option: ")
		_, _ = fmt.Scan(&option)

		switch option {
		case "1":
			task := coreImpl.GetData("C")
			coreImpl.Add(task, tasks)
		case "2":
			coreImpl.List(tasks)
		case "3":
			coreImpl.Update(tasks)
		case "4":
			coreImpl.Delete("", tasks)
		case "5":
			menu = false
		default:
			fmt.Println("Opcion no disponible")
		}
	}

}
