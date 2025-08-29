package main

import (
	"fmt"
	"sync"
	"time"
)

/*Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала – читать эти значения. По истечении N секунд программа должна завершаться.
Подсказка: используйте time.After или таймер для ограничения времени работы.*/

var wg sync.WaitGroup

func main() {
	timeout := time.After(10 * time.Millisecond)
	ch := make(chan int)

	go func() {
		i := 0
		for {
			select {
			case <-timeout:
				close(ch)
				return
			default:
				ch <- i
				i++
			}
		}

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case v, ok := <-ch:
				if !ok {
					return
				}
				fmt.Println(v)
			}
		}
	}()

	wg.Wait()
}
