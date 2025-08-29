package main

import "fmt"

/*Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree").
Создать для неё собственное множество.
Ожидается: получить набор уникальных слов.
Для примера, множество = {"cat", "dog", "tree"}.*/

var arr = []string{"cat", "cat", "dog", "cat", "tree"}

var m = make(map[string]bool)
var s = []string{}

func main() {
	for _, v := range arr {
		if _, ok := m[v]; !ok {
			m[v] = true
			s = append(s, v)
		}
	}
	fmt.Println(s)
}
