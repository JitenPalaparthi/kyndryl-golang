package main

import (
	"demo/calc"
	"fmt"
)

func main() {
	r, err := calc.Add(10, 12.34)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sum:", r)
	}
	r, err = calc.Add(true, 12.34)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sum:", r)
	}

	calc.Sub(-123, 2)

	r2, err := calc.Add("Hello ", 45)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r2)
	}

}
