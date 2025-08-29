package main

import (
	"io"
	"os"
	"strconv"
	"time"
)

/*Реализовать постоянную запись данных в канал (в главной горутине).
Реализовать набор из N воркеров, которые читают данные из этого канала
и выводят их в stdout.
Программа должна принимать параметром количество воркеров
и при старте создавать указанное число горутин-воркеров.*/

func main() {
	ch := make(chan string)

	workers := "1"
	if os.Args[1] != "" {
		workers = os.Args[1]
	}
	j, _ := strconv.Atoi(workers)
	for i := 1; i <= j; i++ {
		go Worker(ch, i)
	}

	for {
		ch <- "data"
		time.Sleep(100 * time.Millisecond)
	}
}

func Worker(ch chan string, n int) {
	for {
		io.WriteString(os.Stdout, <-ch+"\n")
	}
}
