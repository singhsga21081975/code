package main

import (
	"fmt"
	"time"
)

var done = make(chan struct {})
var msgs = make(chan int)

func produce() {
	for i := 0; i < 10; i++ {
		msgs <- i
	}
	fmt.Println("Before closing channel")
	close(msgs)
	fmt.Println("Before passing true to done")
	done <- struct{}{}
}

func consume() {
	for {
		msg := <-msgs

		time.Sleep(100 * time.Millisecond)
		fmt.Println("Consumer: ", msg)
	}
}

func main() {
	go produce()
	go consume()
	<-done
	fmt.Println("After calling DONE")
}
