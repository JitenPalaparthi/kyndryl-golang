package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello Main--Start")
	go Greet() // This is asked to run as a saperate go routine/ green thread
	// defer Greet()
	go func() {
		fmt.Println("Hello Kyndryl!")
	}()
	go func() {
		for i := 1; i <= 10; i++ {
			print(i, " ")
		}
	}()

	fmt.Println("Hello Main--End")
	time.Sleep(time.Millisecond * 1)
}

func Greet() {
	fmt.Println("Hello World")
}

// Go routine is nothing but a green thread
// This is managed and scheduled by the go scheduler

// 1. Go create a go routine just prepend go keyworkd infront of a function
// 2. main is also a go routine
// 3. No go routine waits for other go routine to complete its execution which is by design
// 4. The order of execution of go routines is not in the developers hand by design.
// 5. Cannot prioratise go routines
