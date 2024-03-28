package main

import "fmt"

func main() {
	r := recover()
	fmt.Println(r) // it is nil .. it gets a value only when there is a panic
	func() {       // func1
		fn := "Jiten"
		fullName := FullName(&fn, nil)
		fmt.Println(fullName)
	}()
	println("Continue main")
	func() { //func2
		println("Print even numbers")
		for i := 1; i <= 100; i++ {
			if i%2 == 0 {
				print(i, " ")
			}
		}
		println()
	}()
}

func OnPanic() {
	if r := recover(); r != nil { // if recover not nil means there is a panic
		fmt.Println("recovering from a panic:", r)
		fmt.Println("Log all details")
	}
}

func FullName(fn, ln *string) *string {
	defer OnPanic()
	if fn == nil || *fn == "" {
		panic("firstname is nil or empty")
	}
	if ln == nil || *ln == "" {
		panic("lastname is nil or empty")
	}
	*fn = *fn + *ln
	return fn
}

// recover
