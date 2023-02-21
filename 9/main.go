package main

import (
	"fmt"
	"time"
)

const arrSize = 30000 //Размер массива исходных чисел
const bufSize = 1     //Размер буферизированного канала

func main() {
	now := time.Now()
	var arr [arrSize]int
	for idx := range arr {
		arr[idx] = idx
	}
	firstChan := make(chan int, bufSize)
	//Горутина для записи массива в канал
	go writeArr(firstChan, arr)
	secondChan := make(chan int, bufSize)
	//Горутина для чтения из firstChan, умножения и вывода
	go multiplyX(firstChan, secondChan)
	for val := range secondChan {
		fmt.Println(val)
	}
	fmt.Println("took time:", time.Now().Sub(now))
}

func writeArr(firstChan chan int, arr [arrSize]int) {
	for _, idx := range arr {
		firstChan <- arr[idx]
	}
	close(firstChan)
}

func multiplyX(firstChan chan int, secondChan chan int) {
	for val := range firstChan {
		secondChan <- val * 2
	}
	close(secondChan)
}
