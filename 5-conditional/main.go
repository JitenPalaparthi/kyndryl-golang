package main

import (
	"os"
	"strconv"
)

func main() {

	var age uint8
	if len(os.Args) >= 2 {
		v, _ := strconv.Atoi(os.Args[1])
		age = uint8(v)
	}

	if age >= 18 {
		println("eligible for vote;Age is ", age)
	} else {
		println("not eligible for vote;Age is ", age)
	}
}
