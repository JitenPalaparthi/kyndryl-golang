package main

import "fmt"

func main() {

	//fmt.Println("hello", 1, true, "World",'A',12312.23423)

	arr1 := [5]int{10, 11, 12, 13, 14}
	slice1 := []int{10, 11, 12, 13, 14}

	// slice2 := slice1
	slice2 := arr1[:] // arr1[:]

	fmt.Println("arr1", arr1)
	slice2[2] = 300
	fmt.Println("arr1 after change in slice2", arr1)

	slice3 := slice1[2:]

	slice4 := slice1[2:4] // from 2 to 4 but not 4

	slice5 := slice1[:3] // but not 3
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)
	fmt.Println("slice3", slice3)
	fmt.Println("slice4", slice4)
	fmt.Println("slice5", slice5)

	// 10 11 12 13 14 15

	slice1 = append(slice1[0:3], slice1[4:]...)
	//slice1=append(slice1,68,56,78)
	//slice1=append(slice1,slice3...)
	// 10, 11, 12,14,15
	fmt.Println("slice1", slice1)

}
