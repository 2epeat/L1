package main

import (
	"context"
	"fmt"
	"math/rand"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

/*Реализовать все возможные способы остановки выполнения горутины.
Классические подходы: выход по условию, через канал уведомления,
через контекст, прекращение работы runtime.Goexit() и др.
Продемонстрируйте каждый способ в отдельном фрагменте кода.*/

var wg sync.WaitGroup

func main() {
	// ---
	go func() {
		fmt.Println("native stopping")
		return
	}()

	// ---
	go func() {
		fmt.Println("stopped by Goexit")
		runtime.Goexit()
	}()

	// ---
	go func() {
		for {
			if rand.Intn(100) > 50 {
			} else {
				fmt.Println("stopped by condition")
				return
			}
		}
	}()

	// ---
	timeout := time.After(10 * time.Millisecond)
	go func() {
		for {
			select {
			case <-timeout:
				fmt.Println("stopped by After timeout")
				return
			}
		}
	}()

	// ---
	go func() {
		shutdown, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		for {
			select {
			case <-shutdown.Done():
				fmt.Println("stopped by context with timeout")
				return
			default:
			}
		}
	}()

	// ---
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("stopped by channel")
				return
			default:
			}
		}
	}()
	time.Sleep(1 * time.Second)
	done <- true

	// ---
	go func() {
		for {
			select {
			case v, ok := <-done:
				if !ok {
					fmt.Println("stopped by closing channel")
					return
				}
				done <- v
			}
		}
	}()
	time.Sleep(1 * time.Second)
	close(done)

	// ---
	wg.Add(1)
	go func() {
		defer wg.Done()

		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("stopped by context with signal")
				return
			default:
			}
		}
	}()

	wg.Wait()
}
