package main

import "time"

func main() {

	println("start of main")
	defer func() {
		println("end of main")
		time.Sleep(time.Second * 5)
	}()
	num := 0
	println(100 / num)
	//panic("it is me panic it")
	//
}

// panic always looks for defers to execute before panic
