package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr1 [5]int // this is an array of 5 elements
	// index starts from 0.So it is 0-4

	fmt.Println(arr1) // [0,0,0,0,0]

	arr1[0] = 100
	arr1[1] = 101
	arr1[2] = 102
	arr1[3] = 103
	arr1[4] = 104
	fmt.Println(arr1)
	fmt.Println("Type of arr1:", reflect.TypeOf(arr1)) // [5]int. The type includes the length as well for array
	var arr2 [5]int
	arr2 = arr1
	fmt.Println(arr2)

	arr3 := [5]int{1, 2, 3, 4, 5}   // creating array in shorthand notation
	arr4 := [...]int{1, 2, 3, 4, 5} // length is evaluated at compile time and so [5]int
	fmt.Println(arr3)
	fmt.Println(arr4)
	sum := 0
	for i := 0; i < len(arr1); i++ {
		sum += arr1[i]
	}

	fmt.Println(sum)
	sum = 0
	for _, v := range arr1 {
		sum += v
	}
	fmt.Println(sum)

}
