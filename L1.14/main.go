package main

import (
	"fmt"
	"reflect"
)

/*Разработать программу, которая в runtime способна определить тип переменной,
переданной в неё (на вход подаётся interface{}). Типы,
которые нужно распознавать: int, string, bool, chan (канал).
Подсказка: оператор типа switch v.(type) поможет в решении.*/

func main() {
	var t string
	var r string
	var a = []interface{}{"a", 1, make(chan interface{}), true, map[string]int{"a": 1}}
	for _, v := range a {
		t = typeName(v)
		fmt.Printf("%T | %v\n", v, t)
		r = reflectType(v)
		fmt.Printf("%T | %v\n", v, r)
	}
	fmt.Println("---")
	for _, v := range a {
		r = reflectType(v)
		fmt.Printf("%T | %v\n", v, r)
	}
}

func typeName(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan any, chan string, chan int, chan bool:
		return "chan"
	default:
		return "unknown"
	}
}

func reflectType(v interface{}) string {
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return "chan"
	default:
		return "unknown"
	}
}
