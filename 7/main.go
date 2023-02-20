package main

import (
	"fmt"
	"sync"
)

type myMap struct {
	sync.Mutex
	data map[int]int
}

func main() {
	const n = 5           //число записей
	const writersNum = 12 //число записывающих горутин
	myMap := &myMap{sync.Mutex{}, make(map[int]int)}
	wg := sync.WaitGroup{}
	wg.Add(writersNum * n)

	//Конкурентная запись
	for w := 0; w < writersNum; w++ {
		go func(w int) {
			for i := w * n; i < w*n+n; i++ {
				myMap.Lock()
				myMap.data[i] = i
				myMap.Unlock()
				wg.Done()
			}
		}(w)
	}
	wg.Wait()

	//Результат
	for k, v := range myMap.data {
		fmt.Printf("key %d\tval %d\n", k, v)
	}
	fmt.Println("\nsize of map", len(myMap.data))
}
