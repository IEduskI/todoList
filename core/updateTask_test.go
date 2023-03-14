package core

import (
	"reflect"
	"testing"
)

func TestOptionsImpl_Update(t *testing.T) {
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
		want map[string]Task
	}{
		{
			name: "Test Update",
			args: args{tasks: tasks},
			want: tasks,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := OptionsImpl{}
			if got := o.Update(tt.args.tasks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
