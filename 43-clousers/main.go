package main

import "fmt"

func main() {

	r1 := Calc(30, 20, func(i1, i2 int) int {
		return i1 + i2
	})

	fmt.Println("Add:", r1)

	r2 := Calc(50, 30, Sub)

	fmt.Println("Sub:", r2)

	Mul := func(a, b int) int {
		return a * b
	}

	r3 := Calc(50, 30, Mul)

	fmt.Println("Mul:", r3)

	var Div Func = func(a, b int) int {
		return a / b
	}

	r4 := Calc(50, 30, Div)

	fmt.Println("Div:", r4)

	r5 := Div.Mod(30, 4)
	fmt.Println("Mod:", r5)

	f := FuncRet(10, 20)

	r6 := f()
	fmt.Println("Func Ret", r6)
}

// Clousure
func Calc(a, b int, f func(int, int) int) int {
	return f(a, b) // execution
}

func Sub(a, b int) int {
	return a - b
}

type Func func(int, int) int

func (f Func) Mod(a, b int) int {
	return a % b
}

func FuncRet(a, b int) func() int {
	return func() int {
		return a + b
	}
}
