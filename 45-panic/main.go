package main

import "fmt"

func main() {
	func() { // Func1
		func() { // Func2
			var ptr1 *int
			*ptr1 = 100
			fmt.Println(*ptr1)
		}() // There is a panic in anonymous func.. But the caller of anonymous function is main
		fmt.Println("hey I am anonymous but I am also a caller")
	}()
	fmt.Println("Hello main")
}
