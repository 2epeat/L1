package main

import (
	"fmt"
	"math"
)

/*Дана переменная типа int64. Разработать программу,
которая устанавливает i-й бит этого числа в 1 или 0.
Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).
Подсказка: используйте битовые операции (|, &^).*/

var num int64 = 12

func main() {
	res := swapBit(3)
	fmt.Printf("orig: %d | %b\n", num, num)
	fmt.Printf("resu: %d | %b\n", res, res)
}

func swapBit(bit int64) int64 {
	var bit10 int64 = 1
	if bit > 1 {
		bit10 = int64(math.Pow(float64(2), float64(bit-1)))
	}
	return num ^ bit10
}
