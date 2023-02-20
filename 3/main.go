package main

import (
	"fmt"
	"sync"
)

func main() {
	array := []int{2, 4, 6, 8, 10}
	sumSquares(array) // решение с Mutex и WaitGroup
}

// Способ с использование WaitGroup
func sumSquares(nums []int) {
	var wg sync.WaitGroup
	wg.Add(len(nums))
	sum := 0
	for _, val := range nums {
		go func(val int) {
			val = val * val
			sum = sum + val
			wg.Done()
		}(val)
	}
	wg.Wait()
	fmt.Println(nums)
	fmt.Println(sum)
}
