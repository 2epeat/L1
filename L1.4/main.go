package main

import (
	"context"
	"io"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

/*Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).
Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала прерывания.
Подсказка: можно использовать контекст (context.Context) или канал для оповещения о завершении.*/

var wg sync.WaitGroup

func main() {
	//fmt.Println("Start")
	//defer fmt.Println("End")

	shutdown, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	ch := make(chan string)

	workers := "1"
	if os.Args[1] != "" {
		workers = os.Args[1]
	}
	j, _ := strconv.Atoi(workers)
	for i := 1; i <= j; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work(shutdown, ch, i)
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		send(shutdown, ch)
	}()
	<-shutdown.Done()
	wg.Wait()
}

func send(ctx context.Context, ch chan<- string) {
	//fmt.Println("Start sending")
	//defer fmt.Println("End sending")
	for {
		select {
		case <-ctx.Done():
			//fmt.Println("Close channel")
			close(ch)
			return
		default:
			ch <- "data"
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func work(ctx context.Context, ch <-chan string, n int) {
	//fmt.Println("Start worker " + strconv.Itoa(n))
	//defer fmt.Println("End worker " + strconv.Itoa(n))
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-ch:
			if !ok {
				return
			}
			io.WriteString(os.Stdout, v+"\n")
			//io.WriteString(os.Stdout, v+" | "+strconv.Itoa(n)+"\n")
		}
	}
}
