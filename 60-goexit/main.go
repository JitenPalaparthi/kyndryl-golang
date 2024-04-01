package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// ch := make(chan int)

	// go SendSq(100, ch)

	go func() { //func1
		for i := 1; ; i++ {
			time.Sleep(time.Second * 1)
			fmt.Println(i * i)
			go func() {
				fmt.Println("This is another go routine inside func-1", i)
				runtime.Goexit()
				time.Sleep(time.Second * 5)
			}()
			if i >= 5 {
				runtime.Goexit() // Does not cause fatal if called in other than main
			}
		}
	}()

	fmt.Println("Values are sent to the channel")
	runtime.Goexit() // This is to exit the running go routine
	//time.Sleep(time.Second * 7)
} // since runtime.Goexit is called it will never come to the end of main

// runtime.Goexit is to exit the go routine in which runtime.Goexit is called
// if runtime.Goexit is called in main, it waits for all other go routines to complete
// their execution and fatal main go routine
// If runtime.Goexit is called in normal go routine(other than main), it exits that go routine

func SendSq(num int, ch chan int) {
	for i := 1; i <= num; i++ {
		ch <- i * i
	}
}

// Sender  -> Sends the value
// Receiver -> Receives the value

// Future
// Fanin
// FanOut
// Generator
// Context Implementation
// TimeAfter
