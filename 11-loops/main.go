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
	i := 0
	for ; i < len(arr1); i++ {
		fmt.Print(arr1[i], " ")
	}
	fmt.Println()
	i = 0
	for i < len(arr1) {
		fmt.Print(arr1[i], " ")
		i++
	}
	fmt.Println()
	i = 0
	for {
		if i >= len(arr1) {
			break
		}
		fmt.Print(arr1[i], " ")
		i++
	}
	fmt.Println()

	for i := 0; i < len(arr1); i++ {
		if arr1[i]%2 == 0 {
			fmt.Print(arr1[i], " ")
		}
	}
	fmt.Println()

	for i := 0; i < len(arr1); i++ {
		if arr1[i]%2 != 0 {
			continue
		} else {
			fmt.Print(arr1[i], " ")
		}
	}
	fmt.Println()

	//when ever the value of i is equal to the value inside the array
	//break loops
	arr2d1 := [5][5]int{{6, 34, 54, 65, 94}, {12, 14, 54, 3, 65}}
	//arr2d2 := [5][5]int{{6, 34, 54, 65, 94}, {12, 14, 54, 3, 65}}

	for i := 0; i < len(arr2d1); i++ {
		for j := 0; j < len(arr2d1[i]); j++ {
			if i == arr2d1[i][j] {
				break
			} else {
				println(i, "-->", arr2d1[i][j])
			}
		}
	}

out:
	//out := false
	for i := 1; i <= 10; i = i + 2 {
		// if out {
		// 	break
		// }
		fmt.Println(i, "outer looop")
		for j := 3; j <= 10; j = j + 1 {
			fmt.Println(i, ":", j)
			if i > j {
				//out = true
				//break
				break out
			}
		}
	}
	fmt.Println("I am dont from main")

	a, b := 10, 20
	// t:= b
	// b= a
	// a = t

	a, b = b, a // swap
	a2, b2, c2 := 10, 20, 30
	a2, b2, c2 = b2, c2, a2

	fmt.Println(a, b)
	fmt.Println(a2, b2, c2)

}
