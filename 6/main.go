package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func firstWay(canStop *bool, wg *sync.WaitGroup) {
	fmt.Print("Работа firstWay:")
	defer wg.Done()
	for {
		if !*canStop {
			fmt.Print(".")
			time.Sleep(1 * time.Second)
		} else {
			fmt.Printf("\nРабота firstWay завершена\n")
			return
		}
	}
}

func secondWay(quit chan bool, wg *sync.WaitGroup) {
	fmt.Printf("\n\nРабота secondWay:")
	defer wg.Done()
	for {
		select {
		case <-quit:
			fmt.Printf("\nРабота secondWay завершена\n")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Print(".")
		}
	}
}

func thirdWay(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Printf("\n\nРабота secondWay:")
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\nРабота thirdWay завершена\n")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Print(".")
		}
	}
}

func fourthWay(ch chan struct{}, wg *sync.WaitGroup) {
	fmt.Printf("\n\nРабота fourthWay:")
	defer wg.Done()
	for {
		select {
		case _, opened := <-ch:
			if !opened {
				fmt.Printf("\nРабота fourthWay завершена\n")
				return
			}
		default:
			time.Sleep(1 * time.Second)
			fmt.Print(".")
		}
	}
}

func main() {
	//WaitGroup нужен для корректного вывода
	wg := sync.WaitGroup{}
	wg.Add(1)

	// 1й Способ: дать горутине завершить своё выполнение
	var canStop bool
	go firstWay(&canStop, &wg)
	//Завершение по установке флага canStop
	time.Sleep(3 * time.Second)
	canStop = true
	wg.Wait()

	// 2й Способ: завершение с помощью каналов
	wg.Add(1)
	quit := make(chan bool)
	go secondWay(quit, &wg)
	time.Sleep(3 * time.Second)
	quit <- true
	wg.Wait()

	// 3й Способ: использование контекста
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go thirdWay(ctx, &wg)
	time.Sleep(3 * time.Second)
	cancel()
	wg.Wait()

	//4й Способ: завершение с помощью закрытия канала
	wg.Add(1)
	ch := make(chan struct{})
	go fourthWay(ch, &wg)
	time.Sleep(3 * time.Second)
	close(ch)
	wg.Wait()
}
