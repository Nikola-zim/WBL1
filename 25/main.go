package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	sleep1(2 * time.Second)
	fmt.Println(time.Now().Sub(now))

	now = time.Now()
	sleep2(2 * time.Second)
	fmt.Println(time.Now().Sub(now))

}

// Использование метода after
func sleep1(tm time.Duration) {
	<-time.After(tm)
}

// Использование метода NewTimer
func sleep2(tm time.Duration) {
	timer := time.NewTimer(tm)
	_ = <-timer.C
}
