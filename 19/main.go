package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	var res strings.Builder // для ускорения работы

	for {
		res.Reset()
		_, err := fmt.Scan(&input)
		//Постоянный ввод строки
		if err != nil {
			fmt.Println("Неверный ввод")
			continue
		}

		//Поэлементная запись в res вводимой строки в обратном порядке
		for i := len([]rune(input)) - 1; i >= 0; i-- {
			res.WriteRune([]rune(input)[i])
		}

		fmt.Printf("%s\n", res.String())
	}

}
