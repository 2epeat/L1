package main

import (
	"fmt"
	"math/rand"
)

/*Поменять местами два числа без использования временной переменной.
Подсказка: примените сложение/вычитание или XOR-обмен.*/

var a int = rand.Intn(100)
var b int = rand.Intn(100)

func main() {
	fmt.Println("a:", a, "b:", b)

	a = a ^ b
	b = b ^ a
	a = a ^ b
	fmt.Println("a:", a, "b:", b)

	a = a + b
	b = a - b
	a = a - b
	fmt.Println("a:", a, "b:", b)

	// bonus
	a, b = b, a
	fmt.Println("a:", a, "b:", b)
}
