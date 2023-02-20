package main

import "fmt"

/*Для реализации пересечения двух неупорядоченных множеств в Go можно использовать карты (maps).
Можно использовать карту для представления каждого множества,
где ключи карты являются элементами множества, а значения карты - булевые значения
*/

func intersection(a, b []int) []int {
	// Создаем карту для первого множества
	setA := make(map[int]bool)
	for _, x := range a {
		setA[x] = true
	}

	// Создаем карту для второго множества
	setB := make(map[int]bool)
	for _, x := range b {
		setB[x] = true
	}

	// Создаем слайс для хранения пересечения
	result := make([]int, 0)

	// Проверяем каждый элемент первого множества
	// на вхождение во второе множество
	for x := range setA {
		if setB[x] {
			result = append(result, x)
		}
	}

	return result
}

func main() {
	arr1 := []int{7, 2, 6, 4, 1, 8, 99, 2, 3}
	arr2 := []int{5, 15, 0, 3, 2, 7, 4, 99, 64}
	res := intersection(arr1, arr2)
	fmt.Println(res) // [3 4 5]
}
