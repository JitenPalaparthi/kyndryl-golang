package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var func1 Func = func(i1, i2 int) int {
		return i1 + i2
	}

	r1 := func1(10, 20)
	r2 := func1.Call(10, 20)
	r3 := func1.ToString(10, 20)

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3, "Type:", reflect.TypeOf(r3))
}

type Func func(int, int) int

//var Func1 func(int, int) int

func (f Func) Call(a, b int) int {
	return f(a, b)
}

func (f Func) ToString(a, b int) string {
	return strconv.Itoa(f(a, b))
}
