package main

import "fmt"

func removeElem(sl []int, index int) []int {
	// Проверка на корректность индекса
	if index < 0 || index >= len(sl) {
		return sl
	}

	// Объединение двух части слайса до и после элемента который хотим удалить
	if index != len(sl)-1 {
		return append(sl[:index], sl[index+1:]...)
	} else {
		return sl[:index]
	}
}

func main() {
	arr := make([]int, 10, 30)
	for i := range arr {
		arr[i] = i
	}
	fmt.Printf("Исходный массив: %v\n", arr)

	//Индекс элемента который нужно удалить
	const idx = 5

	arr = removeElem(arr, idx)

	// Массим после удаления
	fmt.Printf("After remove: %v\n", arr)
}
