package main

import (
	"fmt"
	"time"
)

func main() {
	//var ch chan int     // declaration
	ch := make(chan int, 1) // unbuffered channel
	fmt.Println(cap(ch))
	// go func() {
	// 	ch <- 1000
	// }()

	// go func() { // func1
	// 	fmt.Println(<-ch)
	// }()

	go func() { // func1
		//time.Sleep(time.Second * 5)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	}()
	go func() { // func1
		ch <- 1000 // why it is blocked?
		ch <- 1001
		fmt.Println("Why I am executed now?")
	}()

	time.Sleep(time.Millisecond * 1)
}

// Channels: channel is used to communicate go routines
// Do not comunicate by sharing memory;share memory by communicating
// chan keyword is used to crate a channel
// make function is used to instantiate a channel
// channel can be nil
// Two kinds of channels
//	1- buffered channel
// 	2- unbuffered channel
// when a channel is full,[if buffered with one value, if unbuffered the number of values are equal to the size]
// the sender go routine is blocked until the receiver go routine receives the value
// the receiver go routine is blocked until the sender go routine sends the value
// Send a value to the channel: ch <- 100
// receive a value from the channel <-ch
// deadlock means there is sync problems between go routines while sending or receiving channels
