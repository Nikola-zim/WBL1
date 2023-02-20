package main

import "fmt"

// Функция для выполнения бинарного поиска
func binarySearch(arr []int, x int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] < x {
			low = mid + 1
		} else if arr[mid] > x {
			high = mid - 1
		} else {
			return mid
		}
	}

	// Если элемент не найден, возвращаем -1
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9}
	x := 7

	result := binarySearch(arr, x)

	if result == -1 {
		fmt.Printf("Элемент %d не найден\n", x)
	} else {
		fmt.Printf("Элемент %d найден в позиции %d\n", x, result)
	}
}
