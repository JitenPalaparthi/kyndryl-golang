package main

import "fmt"

func main() {

	fmt.Println("Main started")
	defer println("Main ends-2")
	defer println("Main ends-1")
	func() {
		fmt.Println("loop starts")
		func() {
			defer fmt.Println("loop ends")
		}()
		println("Print even numbers")
		for i := 1; i <= 100; i++ {
			if i%2 == 0 {
				print(i, " ")
			}
		}
		println()
	}()
	println("Print odd numbers")
	for i := 1; i <= 100; i++ {
		if i%2 == 1 {
			print(i, " ")
		}
	}
	println()
}
