package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	//Ввод N секунд работы программы
	var n int
	fmt.Print("Введите время работы: ")
	_, err := fmt.Scanf("%v\n", &n)
	if err != nil {
		log.Fatal(err)
	}

	//Работа с каналами
	start := time.Now()

	ch := make(chan uint)
	go func() {
		for {
			fmt.Println(<-ch) // Вывод данных из канала
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	var counter uint
	for {
		//Проверка ограничения времени
		if time.Now().Sub(start) >= time.Duration(n)*time.Second {
			return
		}
		counter++
		ch <- counter // Отправка данных в канал
	}

}
