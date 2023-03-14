package core

import (
	"reflect"
	"testing"
)

func TestOptionsImpl_Delete(t *testing.T) {
	type args struct {
		ID    string
		tasks map[string]Task
	}

	id := "1"
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

	want := make(map[string]Task)

	want["2"] = Task{
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
		want map[string]Task
	}{
		{
			name: "Test Delete",
			args: args{
				ID:    id,
				tasks: tasks,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := OptionsImpl{}
			if got := o.Delete(tt.args.ID, tt.args.tasks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
