package main

import (
	"fmt"
	"strings"
)

func isUnique1(s *string) bool {
	//переводим все символы в нижний регистр и помещаем в слайс из рун
	data := []rune(strings.ToLower(*s))
	l := len(data)
	//перебор в цикле всех пар символов из строки
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if data[j] == data[i] {
				return false
			}
		}
	}
	return true
}

func isUnique2(s *string) bool {
	data := strings.ToLower(*s)
	m := make(map[rune]bool)
	for _, ch := range data {
		// Если по текущему ключу символ не встречался, записываем элемент в map и продолжаем цикл
		if _, ok := m[ch]; !ok {
			m[ch] = false
		} else {
			// Случай, если символ уже был
			return false
		}
	}
	return true
}

func main() {
	var input string
	for {
		fmt.Print("Введите строку: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			return
		}
		fmt.Println(input+" -", isUnique1(&input))
		fmt.Println(input+" -", isUnique2(&input))
	}
}
