package main

import "fmt"

// Функция quickSort выбирает элемент опорный элемент из массива arr и использует функцию partition, чтобы разделить массив на две части:
// одну с элементами меньше опорного элемента и другую с элементами больше опорного элемента.
// Затем рекурсивно вызываем quickSort для каждой из этих двух частей массива.
func quickSort(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}

// Выбираем опорный элемент и перемещаем все элементы меньше него влево, а все элементы больше вправо.
// В конце функции partition возвращается индекс pivot, чтобы разделить массив на две части.
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	arr := []int{3, 5, 1, 4, 2, 0, 0, -1}
	fmt.Println("Before sorting: ", arr)

	quickSort(arr, 0, len(arr)-1)

	fmt.Println("After sorting: ", arr)
}
