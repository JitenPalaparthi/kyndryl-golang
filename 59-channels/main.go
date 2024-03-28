package main

import (
	"fmt"
)

func main() {
	//var ch chan int     // declaration
	ch := make(chan int, 1) // unbuffered channel
	// done := make(chan bool)
	done := make(chan struct{}) // empty struct
	fmt.Println(cap(ch))

	go func() { // func1
		//time.Sleep(time.Second * 5)
		v1 := <-ch
		fmt.Println(v1 * v1)
		v2 := <-ch
		fmt.Println(v2 * v2)
		done <- struct{}{} // sending empty struct value
	}()
	go func() { // func1
		ch <- 1000 // why it is blocked?
		ch <- 1001
		fmt.Println("Why I am executed now?")
	}()
	fmt.Println(<-done)
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
