package main

import (
	"fmt"
)

func main() {
	var i int8
	var data int64
	fmt.Println("Введите номер бита для замены (0..64)")
	for {
		//Получаем значение с консоли
		_, err := fmt.Scan(&i)
		if i < 0 || i > 65 {
			fmt.Println("Неверные значения. Повторите ввод")
			continue
		}
		if err != nil {
			return
		}
		//Замена выбранного бита
		data ^= 1 << (i - 1)
		//Результат
		showRes(data)
	}
}

func showRes(data int64) {
	fmt.Print("Значение в десятичной системе: ", data, "\nВ битах: ")
	for i := 0; i < 64; i++ {
		if (data & (1 << i)) == 0 {
			print("0")
		} else {
			print("1")
		}
	}
	print("\n")
}
