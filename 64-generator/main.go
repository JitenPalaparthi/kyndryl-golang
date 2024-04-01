package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// go func() {
	// 	for i := 1; i <= 10; i++ {
	// 		time.Sleep(time.Millisecond * 300)
	// 		fmt.Println(time.Now().Unix())
	// 	}
	// }()
	ch := sender(3)
	//sum :=

	fmt.Println("Sum of numbers:", <-receive(ch))

	runtime.Goexit() // for simplicity using this .. use waitgroups instead
}

func sender(num int) (ch chan int) {
	ch = make(chan int)
	go func() {
		for i := 1; i <= num; i++ {
			time.Sleep(time.Second * 1)
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func receive(ch <-chan int) (sum chan int) {
	sum = make(chan int)
	s := 0
	go func() {
		for v := range ch { // when using range loop always close the sender
			fmt.Println(v)
			s = s + v
		}
		sum <- s
		close(sum)
	}()
	return sum
}

//chan<- send only go rountine
// <-chan receive only go routine

// 1- Future
// 2- Generator
