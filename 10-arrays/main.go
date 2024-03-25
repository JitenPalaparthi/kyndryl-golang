package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var arr1 [50]int // this is an array of 5 elements
	// index starts from 0.So it is 0-4
	for i := range arr1 {
		arr1[i] = rand.Intn(100)
	}
	fmt.Println(arr1)

	var arr2d1 [3][3]int
	count := 1
	for i := 0; i < len(arr2d1); i++ {
		for j := 0; j < len(arr2d1[i]); j++ {
			arr2d1[i][j] = count
			count++
		}
	}
	fmt.Println(arr2d1)

	arr2d2 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2d2)

	arr3d1 := [2][2][3]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {10, 11, 12}}}
	sum := 0
	for _, v1 := range arr3d1 {
		for _, v2 := range v1 {
			for _, v3 := range v2 {
				sum += v3
			}
		}
	}
	fmt.Println("Sum of arr3d1:", sum)
}
