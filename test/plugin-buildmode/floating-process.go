package main

import "fmt"

type Anigo struct {
	Properties map[string]interface{}
}

type Status struct {
	Status string
	Name   string
}

type ProcessPlugin[_ float32 | int32] struct {
	Handler func(*Anigo)
	Test    func() []Status

	Name string
}

var Plugin = ProcessPlugin[float32]{
	Handler: func(a *Anigo) {
		fmt.Println("a")
	},
}
