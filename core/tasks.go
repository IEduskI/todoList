package core

type Task struct {
	ID          string
	Nombre      string
	Descripcion string
	Estado      int
	FechaIni    string
	FechaFin    string
	Prioridad   int
}

func NewTask(id, nombre, desc, fIni, fFin string, estado, prioridad int) Task {
	return Task{
		ID:          id,
		Nombre:      nombre,
		Descripcion: desc,
		Estado:      estado,
		FechaIni:    fIni,
		FechaFin:    fFin,
		Prioridad:   prioridad,
	}
}

type AddTask interface {
	Add(task Task, tasks map[string]Task) map[string]Task
}

type GetTask interface {
	Get(ID string, tasks map[string]Task) Task
}

type ListTasks interface {
	List(tasks map[string]Task)
}

type UpdateTasks interface {
	Update(tasks map[string]Task) map[string]Task
}

type DeleteTask interface {
	Delete(ID string, tasks map[string]Task) map[string]Task
}

type CaptureData interface {
	GetData(transaction string) Task
}

type Options interface {
	AddTask
	ListTasks
	UpdateTasks
	DeleteTask
	CaptureData
}
