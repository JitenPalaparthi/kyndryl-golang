package main

import "fmt"

func main() {

	var num1 int = 10 // 8 bytes --10

	// var ptr1 *int = &num1

	// // int num = 10;
	// // *int ptr = &num

	// fmt.Println("Value of num1", num1, "Value of num1 using pointer:", *ptr1, "Value of num1 at address:", *&num1)
	// Incr(num1) // call by value
	// fmt.Println(num1)
	// IncrPtr(ptr1)  // call by value
	// IncrPtr(&num1) // call by value
	// fmt.Println(num1)

	// var ptr2 **int = &ptr1

	// var ptr3 ***int = &ptr2

	// fmt.Println(***ptr3)

	ptr4 := Square(num1) // these kind of allocations are generally in the heap
	fmt.Println(*ptr4)

	//var ptr5 *int = &num1
	{
		//ptr5 = nil
		ptr5 := new(int) // 8 bytes , 0 give address to ptr5

		*ptr5 = 256

		fmt.Println(*ptr5)

		SquareOut(12, ptr5)

		fmt.Println(*ptr5)
	}

}

// stores the memory addres of a variable and also address or a memory allocated by new function
// directly you cant do pointer arthemetic
// poiners hold addresses

func Incr(num int) {
	num++
}

func IncrPtr(num *int) {
	*num++
}

func Square(num int) *int {
	sq := num * num
	return &sq
}

func SquareOut(num int, sq *int) {
	*sq = num * num
}
