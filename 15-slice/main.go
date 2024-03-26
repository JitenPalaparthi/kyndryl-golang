package main

import "fmt"

func main() {
	//slice1 := []int{10, 12, 13, 14, 53}
	// slice1 := make([]int, 5, 10)
	// slice1[0], slice1[1], slice1[2], slice1[3], slice1[4] = 10, 12, 13, 14, 53
	// slice2 := slice1 // shallow copy
	// slice1 = append(slice1, 15, 16, 17, 18, 19)
	// IncrSliceBy(slice2, 10)
	// fmt.Println(slice1)
	// fmt.Println(slice2)
	slice3 := make([]int, 5, 10)
	slice3[0], slice3[1], slice3[2], slice3[3], slice3[4] = 1, 2, 3, 4, 5
	slice3 = AddIncrToSlice(slice3, 6)
	fmt.Println(slice3)

	slice3 = AddIncrToSlice(slice3, 7)
	fmt.Println(slice3)

	slice3 = AddIncrToSlice(slice3, 8)
	fmt.Println(slice3)

	slice3 = AddIncrToSlice(slice3, 9)
	fmt.Println(slice3)

	slice3 = AddIncrToSlice(slice3, 10)
	fmt.Println(slice3)

	slice3 = AddIncrToSlice(slice3, 11)
	fmt.Println(slice3)
	/*
		slice1:
		len:10
		cap:10
		ptr:0x110

		slice2:
		len:5
		cap:10
		ptr:0x110
	*/

	slice4 := []int{}
	var slice5 []int
	fmt.Println(len(slice4), cap(slice4), slice4)
	if slice4 == nil {
		fmt.Println("yes nil")
	}
	fmt.Println(len(slice5), cap(slice5), slice5)
	if slice5 == nil {
		fmt.Println("yes nil")
	}
}

func IncrSliceBy(arr []int, inc int) { // generally arrays and normal variables are pass by value
	for i, v := range arr {
		arr[i] = v + inc
	}
}

// slice, 12 --> 144 should be appended to the slice
// before that it should square all elements in the slice and append
func AddIncrToSlice(slice []int, num int) []int {
	for i, v := range slice {
		slice[i] = v + 1
	}
	return append(slice, num+1) // The rule is when ever you use append .. return the slice
	// better to follow when slice is used as input param, return it
}
