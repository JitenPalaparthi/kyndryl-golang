package main

import (
	"fmt"
	"math/rand"
	"unsafe"
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

	slice4 := make([]byte, 4)
	slice4 = append(slice4, 10, 12, 13, 14, 15)
	// slice4[0]=10
	// slice4[1]=12
	fmt.Println(slice4, len(slice4), cap(slice4), unsafe.Sizeof(slice4)) // 16 bytes
	slice5 := make([]string, 2)
	slice5 = append(slice5, "hello", "World", "How", "are", "you")
	fmt.Println(slice5, len(slice5), cap(slice5), unsafe.Sizeof(slice5)) // 16 bytes
	fmt.Println(unsafe.Sizeof(slice5[0]))
	fmt.Println(len(slice5[3]))

	str := "hello kyndryl folks..welcome to Golang."
	fmt.Println(len(str) * 4)

	str1 := "Hello World!"
	str2 := "Hello 您好"

	for i := 0; i < len(str1); i++ {
		fmt.Print(string(str1[i]), "-")
	}
	// H-e-l-l-o- -W-o-r-l-d-!-
	// in str1 , we are very clear that all these are ascii chars so each carries 1 byte
	fmt.Println()

	for i := 0; i < len(str2); i++ {
		fmt.Print(string(str2[i]), "-")
	}
	// H-e-l-l-o--// we dont know what it prints
	fmt.Println()

	for _, v := range str2 {
		fmt.Print(string(v), "-")
	}
	fmt.Println()

	// utf8 architecture

	// [0 0 0 0 10 12 13 14]
}
