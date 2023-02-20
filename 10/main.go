package main

import (
	"fmt"
)

var arr = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 5.1, -5.1, 0}

// Нахождение нужной группы для заданной температуры
func findGroup(temp float64) int {
	switch {
	case temp < 0:
		return (int(temp)/10 - 1) * 10
	case temp >= 0:
		return int(temp) / 10 * 10
	default:
		return 0
	}
}

// Проходит по массиву температур и добавляет в map к нужной группе эту температуру
func makeGroups(arr []float64) {
	var res = make(map[int][]float64)
	for _, temp := range arr {
		//Добавление к нужному слайсу нужной температуры
		res[findGroup(temp)] = append(res[findGroup(temp)], temp)
	}
	for key, value := range res {
		fmt.Println("Key:", key, "Value:", value)
	}
}

func main() {
	makeGroups(arr)
}
