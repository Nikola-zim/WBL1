package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var num1, num2 int64 = -5, 0

	//1. Арифметический
	num1 += num2
	num2 = num1 - num2
	num1 = num1 - num2
	fmt.Printf("num1 = %d; num2 = %d\n", num1, num2)

	//2. Присвоение
	num1, num2 = num2, num1
	fmt.Printf("num1 = %d; num2 = %d\n", num1, num2)

	//3. Битовые операции
	num1 ^= num2 // (x ^ y, y)
	num2 ^= num1 // (x ^ y, y ^ x ^ y) = (x ^ y, x)
	num1 ^= num2 // (x ^ y ^ x, x)     = (y, x)
	fmt.Printf("num1 = %d; num2 = %d\n", num1, num2)

	//4.Atomic
	num2 = atomic.SwapInt64(&num1, num2)
	fmt.Printf("num1 = %d; num2 = %d\n", num1, num2)
}
