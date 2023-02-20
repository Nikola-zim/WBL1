package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)
	res := new(big.Int)

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter first number: ")
		a.SetString(readLine(reader), 10)

		fmt.Print("Enter second number: ")
		b.SetString(readLine(reader), 10)

		fmt.Print("Enter operation (+, -, *, /): ")
		operation := readLine(reader)

		switch operation {
		case "+":
			res.Add(a, b)
		case "-":
			res.Sub(a, b)
		case "*":
			res.Mul(a, b)
		case "/":
			res.Div(a, b)
		default:
			fmt.Println("Invalid operation")
			return
		}
		fmt.Println("Result: ", res)
	}
}

func readLine(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
