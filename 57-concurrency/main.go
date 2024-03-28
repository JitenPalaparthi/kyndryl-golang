package main

import (
	"fmt"
	"sync"
)

//var num *int

func main() {
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	num := new(int) // shared variable
	//var v int = 1
	//	num = &v
	for i := 1; i <= 100; i++ {
		wg.Add(1) // wg.Counter++
		//go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		go func(wg *sync.WaitGroup) {
			//mu.Lock()
			Incr(mu, num)
			//mu.Unlock()
			wg.Done()
			//}(wg, mu)
		}(wg)
	}
	wg.Wait()
	fmt.Println("The ultimate result", *num)
}

func Incr(mu *sync.Mutex, num *int) {
	println(*num)
	mu.Lock()
	*num++
	mu.Unlock()
}

// func Decr(mu *sync.Mutex, num *int) {
// 	mu.Lock()
// 	*num--
// 	mu.Unlock()
// }
