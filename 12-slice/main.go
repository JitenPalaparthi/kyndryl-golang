package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var slice1 []int // only declared but not instantiated
	fmt.Println(slice1)
	if slice1 == nil { // this means , slice has been declared but not instantiated
		slice1 = make([]int, 5, 10) // instantiating a slice
	}
	fmt.Println(slice1) // type inference works

	for i := range slice1 {
		slice1[i] = rand.Intn(100)
	}
	// [0 0 0 0 0]
	fmt.Println(slice1, "Address:", &slice1[0], "Length:", len(slice1), "Capacity:", cap(slice1))
	slice1 = append(slice1, 121)
	fmt.Println(slice1, "Address:", &slice1[0], "Length:", len(slice1), "Capacity:", cap(slice1))
	slice1 = append(slice1, 122)
	fmt.Println(slice1, "Address:", &slice1[0], "Length:", len(slice1), "Capacity:", cap(slice1))
	slice1 = append(slice1, 123, 124, 125, 126, 127)
	fmt.Println(slice1, "Address:", &slice1[0], "Length:", len(slice1), "Capacity:", cap(slice1))

	slice2 := slice1 // shallow copy.. Only the headers are copied
	slice3 := make([]int, len(slice2))
	copy(slice3, slice2) // deep copy

}
