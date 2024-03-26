package main

import "fmt"

func main() {

	//fmt.Println("hello", 1, true, "World",'A',12312.23423)
	fmt.Println(SumOf())
	fmt.Println(SumOf(10, 20))
	fmt.Println(SumOf(10, 20, 234, 45, 4, 6, 57, 7, 5, 6, 454, 56, 545, 66))
	arr1 := [5]int{10, 12, 13, 14, 15}
	fmt.Println(SumOf(arr1[:]...))
	slice1 := []int{10, 12, 13, 14, 15}
	fmt.Println(SumOf(slice1...))

}

// variadic parameter must the last parameter
// since it should be the last parmater , you cannot have more than one variadic
// variadics are only created for functions or methods.
// cannot create a variable as a variadic
func SumOf(nums ...int) (sum int) {
	for _, v := range nums {
		sum += v
	}
	return sum
}
