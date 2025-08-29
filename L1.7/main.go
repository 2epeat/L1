package main

import (
	"fmt"
	"sync"
)

/*Реализовать безопасную для конкуренции запись данных в структуру map.
Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).
Проверьте работу кода на гонки (util go run -race).*/

var myMap = make(map[int]int)
var wg sync.WaitGroup
var mutex = &sync.Mutex{}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			myMap[i] = i
			mutex.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(myMap)
}
