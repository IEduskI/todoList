package core

import "testing"

func TestOptionsImpl_List(t *testing.T) {
	type args struct {
		tasks map[string]Task
	}

	tasks := make(map[string]Task)
	tasks["1"] = Task{
		ID:          "1",
		Nombre:      "Prueba",
		Descripcion: "Prueba",
		Estado:      1,
		FechaIni:    "01/03/2023",
		FechaFin:    "02/02/2023",
		Prioridad:   3,
	}

	tasks["2"] = Task{
		ID:          "2",
		Nombre:      "Prueba2",
		Descripcion: "Prueba2",
		Estado:      2,
		FechaIni:    "02/02/2023",
		FechaFin:    "02/05/2023",
		Prioridad:   3,
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test List",
			args: args{tasks: tasks},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := OptionsImpl{}
			o.List(tt.args.tasks)
		})
	}
}
