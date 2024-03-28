package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var ok bool

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU()) // 8
	wg := new(sync.WaitGroup)
	fmt.Println("Hello Main--Start")
	//wg.Add(4)
	wg.Add(1)
	go GreetW(wg)
	wg.Add(1) // wg.Counter++
	go func(wg *sync.WaitGroup) {
		Greet()   // This is asked to run as a saperate go routine/ green thread
		wg.Done() // wg.Counter--
	}(wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		time.Sleep(time.Second * 5)
		fmt.Println("Hello Kyndryl!")
		ok = true
		wg.Done()
	}(wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		time.Sleep(time.Second * 8)
		fmt.Println("Hello Kyndryl!---5")
		wg.Done()
	}(wg)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		for i := 1; i <= 10; i++ {
			print(i, " ")
		}
		wg.Done()
	}(wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		for !ok {
			fmt.Println("Number of go routines", runtime.NumGoroutine())
			fmt.Println("Number of CPU:", runtime.NumCPU())
			fmt.Println("Number of CGO:", runtime.NumCgoCall())
			time.Sleep(time.Millisecond * 100)
		}
		wg.Done()
	}(wg)

	fmt.Println("Hello Main--End")

	wg.Wait() // waiting until the counter becomes zero

}

// func myFatal(str string) {
// 	println(str)
// 	os.Exit(1)
// }

func Greet() {
	fmt.Println("Hello World")
}
func GreetW(wg *sync.WaitGroup) {
	fmt.Println("Hello World-2")
	wg.Done()
}

// Go routine is nothing but a green thread
// This is managed and scheduled by the go scheduler

// 1. Go create a go routine just prepend go keyworkd infront of a function
// 2. main is also a go routine
// 3. No go routine waits for other go routine to complete its execution which is by design
// 4. The order of execution of go routines is not in the developers hand by design.

// WaitGroup
// Count
// Wait untile the counter becomes 0
