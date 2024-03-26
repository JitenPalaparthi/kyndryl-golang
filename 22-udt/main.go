package main

import "fmt"

func main() {

	c1 := ColorCode{int: 101, string: "Red", string1: "15:50:40"}
	fmt.Println(c1.int)
	fmt.Println(c1.string)
	fmt.Println(c1.string1)
	c2 := new(ColorCode)
	fmt.Println(*c2)

}

// a struct without field names
type ColorCode struct {
	int
	string
	string1
}

type string1 = string
