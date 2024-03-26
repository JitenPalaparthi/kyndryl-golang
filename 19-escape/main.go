package main

import (
	"math/rand"
)

// This demo is for escape analysis. Run the below command(with -gcflags) to see what gets allocated in heap
// go run -gcflags="-m" main.go

var map1 map[string]string

const (
	MAX = 100
	MIN = 10
)

func main() {

	var num1 int = 100

	slice1 := make([]int, 0)
	println(num1)
	for _, v := range slice1 {
		print(v, " ")
	}
	println()
	for i := 1; i <= 1000000000; i++ {
		slice1 = append(slice1, rand.Intn(20000))
	}

	slice2 := slice1

	slice2 = append(slice2, 100)
	// for _, v := range slice1 {
	// 	print(v, " ")
	// }
	println()

	slice3 := make([]any, 10)

	slice3[0] = "hello World"
	slice3[1] = true
	slice3[2] = 3432.32432
	//println(slice3)
	slice4 := []any{"Hello World", true, 21321.1312, 0, 123, complex(12, 23.4)}
	println(slice4)

	ptr4 := Square(num1) // these kind of allocations are generally in the heap
	println(*ptr4)

	//var ptr5 *int = &num1
	{
		//ptr5 = nil
		ptr5 := new(int) // 8 bytes , 0 give address to ptr5

		*ptr5 = 256

		println(*ptr5)
		//fmt.Println(*ptr5)

		SquareOut(12, ptr5)

		println(*ptr5)
	}
}

func Square(num int) *int {
	sq1 := num * num
	sq3 := new(int)

	return sq3
	return &sq1
}

func SquareOut(num int, sq2 *int) {
	*sq2 = num * num
}

// does not escape
// escapes to heap
// move to heap
