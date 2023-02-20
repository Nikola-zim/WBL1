package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"
)

const delay = 1000 * time.Millisecond

func main() {
	numbWork := flag.Int("n", 1, "number of workers")
	flag.Parse()
	signalChan := make(chan os.Signal, 1) //Канал для ожидания прерывания
	cleanupDone := make(chan bool)
	c := make(chan interface{}) //Канал с произвольными данными
	publisher(c, *numbWork)
	workers(c, *numbWork)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			cleanupDone <- true
			close(c)
		}
	}()
	<-cleanupDone
}

func publisher(c chan interface{}, numbWork int) {
	var sleepingTime int
	//Задержка времени с расчётом колличество воркеров и константного времни
	sleepingTime = int(delay) / numbWork * 2
	go func() {
		for {
			c <- "Some Data"
			time.Sleep(time.Duration(sleepingTime))
		}
	}()
}

func workers(c chan interface{}, numbWork int) {
	for i := 0; i <= numbWork; i++ {
		go func() {
			for {
				fmt.Println(<-c)
			}
		}()
	}
}
