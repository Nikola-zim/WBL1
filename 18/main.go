package main

import (
	"fmt"
	"sync"
)

type concurrentCounter struct {
	wg    sync.WaitGroup
	mut   sync.Mutex
	count int
}

func (c *concurrentCounter) IncrementSync() {
	// Обеспечиваем доступ только одной горутине
	c.mut.Lock()
	c.count++
	c.mut.Unlock()
	c.wg.Done()
}

func main() {
	counter := &concurrentCounter{sync.WaitGroup{}, sync.Mutex{}, 0}
	const numbOfIterations = 1000 //Задаём желаемое
	for i := 0; i < numbOfIterations; i++ {
		counter.wg.Add(1)
		go counter.IncrementSync()
	}
	counter.wg.Wait()
	fmt.Println(counter.count)
}
