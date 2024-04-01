package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	ch := make(chan int)
	sum := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Millisecond * 300)
			fmt.Println(time.Now().Unix())
		}
	}()
	go sender(10, ch)
	go receive(ch, sum)

	go fmt.Println("Sum of numbers:", <-sum)

	runtime.Goexit() // for simplicity using this .. use waitgroups instead
}

func sender(num int, ch chan<- int) {
	for i := 1; i <= num; i++ {
		time.Sleep(time.Second * 1)
		ch <- i
	}
	close(ch)
}

func receive(ch <-chan int, sum chan<- int) {
	s := 0
	for v := range ch { // when using range loop always close the sender
		s = s + v
	}
	sum <- s
}

//chan<- send only go rountine
// <-chan receive only go routine

// 1- Future
// 2- Generator
