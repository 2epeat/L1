package main

import (
	"fmt"
	"math"
	"time"
)

/*
Написать программу, которая конкурентно рассчитает значения квадратов чисел,
взятых из массива [2,4,6,8,10], и выведет результаты в stdout.
Подсказка: запусти несколько горутин, каждая из которых возводит число в квадрат.
*/

var nums = [...]int{2, 4, 6, 8, 10}

func main() {
	for _, num := range nums {
		go fmt.Printf("%.0f ", func(num int) float64 {
			return math.Pow(float64(num), 2)
		}(num))
	}
	time.Sleep(1 * time.Millisecond)
}
