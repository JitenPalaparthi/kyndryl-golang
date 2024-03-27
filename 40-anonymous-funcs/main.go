package main

import "fmt"

func main() {

	// a simple hello world

	func() {
		fmt.Println("hello world")
	}()

	greet := func() {
		fmt.Println("hello world")
	}
	greet()

	a1, b1 := 10, 20
	func(a, b int) {
		println(a + b)
	}(a1, b1)

	c1 := func(a, b int) int {
		return a + b
	}(a1, b1)
	fmt.Println(c1)

	c2 := func(a, b int) int {
		return a + b
	}
	fmt.Println(c2(a1, b1))

	var c3 func(a, b, c int) int // if you want to assign a value to the variable what value can be assigned
	fmt.Println(c3)
	c3 = add
	r1 := c3(10, 20, 30)
	fmt.Println(r1)
	c3 = sub
	r2 := c3(30, 20, 10)
	fmt.Println(r2)

	// var num1 int
	// var slice1 []int
	// var arr1 [3]int
	// var map1 map[string]any
	// var c4 func(a, b, c int) int

}

func add(a1, b2, c3 int) int {
	return a1 + b2 + c3
}

func sub(a, b, c int) int {
	return a - b - c
}

func mul(a, b, c int) int {
	return a * b * c
}

// anonymous func
// a func that does not have a name is called anonymous func
//
