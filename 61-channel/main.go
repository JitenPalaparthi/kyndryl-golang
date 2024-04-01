package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)

	// go SendSqU( ch)
	// go ReceiveSqU(ch)

	go SendSq(10, ch)
	go ReceiveSq(ch)

	fmt.Println("Values are sent to the channel")
	runtime.Goexit() // This is to exit the running go routine
} // since runtime.Goexit is called it will never come to the end of main

func SendSq(num int, ch chan int) {
	for i := 1; i <= num; i++ {
		ch <- i * i
		time.Sleep(time.Millisecond * 100)
	}
	close(ch) // close the channel once sending is done
}

func ReceiveSq(ch chan int) {
	for {
		v, ok := <-ch // value, and whether the channel is closed or not; If ok== false , ch is closed
		if ok {
			fmt.Println(v)
		} else {
			return
		}
	}
}

func ReceiveSqRange(ch chan int) {
	for v := range ch { // Range loop can iterate until the channel is closed
		fmt.Println(v)
	}
}

func SendSqU(ch chan int) {
	for i := 1; ; i++ {
		ch <- i * i
		time.Sleep(time.Millisecond * 100)
	}
	//close(ch) // close the channel once sending is done
}

func ReceiveSqU(ch chan int) {
	for {
		v := <-ch
		fmt.Println(v)
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
