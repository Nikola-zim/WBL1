package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseStr(str *string) {
	// Создаём массив из слов в строке
	words := strings.Split(*str, " ")
	// Запись в обратном порядке
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	*str = strings.Join(words, " ")
	return
}

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	// Постоянный ввод и вывод результатов
	for {
		scanner.Scan()
		input = scanner.Text()
		reverseStr(&input)
		fmt.Println(input)
	}

}
