package core

import (
	"reflect"
	"testing"
)

func TestNewOptionsImpl(t *testing.T) {
	tests := []struct {
		name string
		want OptionsImpl
	}{
		{
			name: "Test Constructor",
			want: OptionsImpl{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOptionsImpl(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOptionsImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptionsImpl_Add(t *testing.T) {
	type args struct {
		task  Task
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

	task := Task{
		ID:          "2",
		Nombre:      "Prueba2",
		Descripcion: "Prueba2",
		Estado:      2,
		FechaIni:    "02/02/2023",
		FechaFin:    "02/05/2023",
		Prioridad:   3,
	}

	want := make(map[string]Task)
	want["1"] = Task{
		ID:          "1",
		Nombre:      "Prueba",
		Descripcion: "Prueba",
		Estado:      1,
		FechaIni:    "01/03/2023",
		FechaFin:    "02/02/2023",
		Prioridad:   3,
	}

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
			name: "Test add",
			args: args{
				task:  task,
				tasks: tasks,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := OptionsImpl{}
			if got := o.Add(tt.args.task, tt.args.tasks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
