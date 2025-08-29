package main

import (
	"io"
	"os"
	"strconv"
	"sync"
	"time"
)

/*Разработать конвейер чисел. Даны два канала: в первый пишутся числа x из массива,
во второй – результат операции x*2. После этого данные из второго канала должны выводиться в stdout.
То есть, организуйте конвейер из двух этапов с горутинами: генерация чисел и их обработка.
Убедитесь, что чтение из второго канала корректно завершается.*/

// send
var s = make(chan int)

// result
var r = make(chan int)

var arr = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			v, ok := <-s
			if !ok {
				close(r)
				return
			}
			r <- v * 2
		}
	}()

	wg.Add(1)
	go func() {
		wg.Done()
		for i := 0; i < 10; i++ {
			s <- arr[i]
		}
		close(s)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			time.Sleep(1 * time.Second)
			v, ok := <-r
			if !ok {
				return
			}
			io.WriteString(os.Stdout, strconv.Itoa(v)+"\n")
		}
	}()
	wg.Wait()
}
