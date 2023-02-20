package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	var res strings.Builder

	for {
		res.Reset()
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Неверный ввод")
			continue
		}
		//fmt.Printf("Length of given string is %d\n", len([]rune(input)))
		for i := len([]rune(input)) - 1; i >= 0; i-- {
			res.WriteRune([]rune(input)[i])
		}

		fmt.Printf("%s\n", res.String())
	}

}
