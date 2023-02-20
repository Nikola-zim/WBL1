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
	for idx, _ := range arr {
		arr[idx] = idx
	}
	//Горутина для записи массива в канал
	firstChan := make(chan int, bufSize)
	go writeArr(firstChan, arr)
	secondChan := make(chan int, bufSize)
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
