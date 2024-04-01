package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := GetServer("Srver1", time.Millisecond*1)
	ch2 := GetServer("Srver2", time.Millisecond*1)
	ch3 := GetServer("Srver3", time.Millisecond*1)
	timeout := time.After(time.Millisecond * 1)

	for {
		select {
		case s := <-ch1:
			fmt.Println(s)
			return
		case s := <-ch2:
			fmt.Println(s)
			return
		case s := <-ch3:
			fmt.Println(s)
			return
		case <-timeout:
			fmt.Println("timedout")
			return
		default:
			//fmt.Println("This is default----")
		}
	}
}

func GetServer(server string, tm time.Duration) <-chan string {
	srv := make(chan string)
	go func() {
		time.Sleep(tm)
		srv <- server
		close(srv)
	}()
	return srv
}

// When you send a value to the channel it receives the value
// When you close the channel also it recived the value
