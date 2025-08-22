package main

import "fmt"

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
