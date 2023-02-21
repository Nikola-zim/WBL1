package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	array := []int{2, 4, 6, 8, 10}
	sumSquares(array) // решение с Mutex и WaitGroup
	SolutionAtomic(array)
}

// Способ с использование WaitGroup
func sumSquares(nums []int) {
	var wg sync.WaitGroup
	wg.Add(len(nums))
	sum := 0
	fmt.Println("SolutionWaitGroup")
	for _, val := range nums {
		go func(val int) {
			val = val * val
			sum = sum + val
			wg.Done()
		}(val)
	}
	wg.Wait()
	fmt.Println("array :=", nums)
	fmt.Println("sum :=", sum)
}

// SolutionAtomic с помощью пакета Atomic
func SolutionAtomic(nums []int) {
	var sum int64
	wg := sync.WaitGroup{}
	wg.Add(len(nums))
	fmt.Println("SolutionAtomic")
	for _, v := range nums {
		go func(vi int) {
			atomic.AddInt64(&sum, int64(vi*vi))
			wg.Done()
		}(v)
	}
	wg.Wait()
	fmt.Println("array :=", nums)
	fmt.Println("sum :=", sum)
}
