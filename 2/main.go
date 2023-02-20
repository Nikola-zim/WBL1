package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// Функция подсчёта кваадратов
func calcPow2(evenList *[5]float64, wg *sync.WaitGroup) {
	for index, value := range evenList {
		evenList[index] = math.Pow(value, 2)
	}
	defer wg.Done()
}

func calcPow2new(hugeEvenList []float64, batchSize int) {
	start := time.Now()
	hugeEvenListLen := len(hugeEvenList)
	var wg1 sync.WaitGroup
	wg1.Add(hugeEvenListLen)
	stepsNum := hugeEvenListLen / batchSize
	//Разбиваем массив на части по размеру batchSize и раздаём горутинам
	for del := 0; del < stepsNum; del++ {
		go func(del int) {
			for idx := batchSize * del; idx < (del+1)*batchSize; idx++ {
				hugeEvenList[idx] = math.Pow(hugeEvenList[idx], 2)
				wg1.Done()
			}
		}(del)
	}
	//Работа с остатком
	batchRest := len(hugeEvenList) % batchSize //размер остатка
	restStartIndex := batchSize * stepsNum     //начальный индекс остатка
	for idx := restStartIndex; idx < restStartIndex+batchRest; idx++ {
		hugeEvenList[idx] = math.Pow(hugeEvenList[idx], 2)
		wg1.Done()
	}
	//Ожидание завершения горутин
	wg1.Wait()
	//Вывод части слайса
	fmt.Println(hugeEvenList[0:15])
	//Время выполнения
	duration := time.Since(start)
	fmt.Println("Время выполнения calcPow2new: ", duration.Milliseconds())

}

func main() {
	//Решение №1
	var wg sync.WaitGroup
	evenList := [...]float64{2, 4, 6, 8, 10}
	wg.Add(1)
	//Передаем массив и счётчик активных элементов
	go calcPow2(&evenList, &wg)
	//Ждём завершения горутины
	wg.Wait()
	fmt.Println("Результат calcPow2: ", evenList)

	//Решение №2
	//Создадим массив побольше
	//runtime.GOMAXPROCS(12)
	const size = 300000000     //размер нового массива
	const batchSize = 25000000 //размер слайса для обработки в одно горутине
	hugeEvenList := make([]float64, size, size)
	for i := 0; i < size; i++ {
		hugeEvenList[i] = float64(i)
	}
	calcPow2new(hugeEvenList, batchSize)

}
