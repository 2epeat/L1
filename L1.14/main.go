package main

import (
	"fmt"
	"os"
)

/*Разработать программу, которая в runtime способна определить тип переменной,
переданной в неё (на вход подаётся interface{}). Типы,
которые нужно распознавать: int, string, bool, chan (канал).
Подсказка: оператор типа switch v.(type) поможет в решении.*/

func main() {
	var t string
	if os.Args[1] != "" {
		t = typeName(os.Args[1])
	}
	fmt.Printf("%T | %v\n", t, t)
}

func typeName(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan interface{}:
		return "chan"
	default:
		return "unknown"
	}
}
