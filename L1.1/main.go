package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
Подсказка: используйте композицию (embedded struct), чтобы Action имел все методы Human.
*/

type Human struct {
	Name string
}

func (h Human) Greet() {
	fmt.Println("Hello, my name is", h.Name)
}

type Action struct {
	Human
	Running bool
}

func (a Action) Run() {
	fmt.Printf("%s running\n", a.Name)
	a.Running = true
}

func main() {
	a := Action{
		Human: Human{
			Name: "Viktor",
		},
		Running: false,
	}

	a.Greet()
	a.Run()
}
