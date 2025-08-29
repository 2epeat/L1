package main

import (
	"fmt"
	"math/rand"
	"slices"
)

/*Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов) —
т.е. вывести элементы, присутствующие и в первом, и во втором.
Пример:
A = {1,2,3}
B = {2,3,4}
Пересечение = {2,3}*/

var a = []int{}
var b = []int{}

func main() {
	var j = 3 + rand.Intn(5)
	a = makeSlice(a, j)
	b = makeSlice(b, j)

	fmt.Println(a)
	fmt.Println(b)

	var c = map[int]int{}
	for _, v := range b {
		if slices.Contains(a, v) {
			c[v] = v
		}
	}
	fmt.Println(c)
}

func makeSlice(slice []int, length int) []int {
	length = length - 1
	for {
		if len(slice) == length {
			break
		}
		r := rand.Intn(10)
		if !slices.Contains(slice, r) {
			slice = append(slice, r)
		}
	}

	return slice
}
