package main

import "fmt"

func main() {

	fn, ln := "Jiten", "p"
	fullName := FullName(&fn, &ln)

	fmt.Println(fullName)

	// uncomment the below code  to see panic
	// fullName = FullName(nil, &ln)
	// fmt.Println(fullName)

	//fmt.Println(100 / 0) // divided by zero // The compiler finds the error
	var num int
	fmt.Println(100 / num) // here the compiler cannot determine the error

}

func FullName(fn, ln *string) *string {
	if fn == nil || *fn == "" {
		panic("firstname is nil or empty")
	}
	if ln == nil || *ln == "" {
		panic("lastname is nil or empty")
	}
	//fulln := new(string)
	//*fulln = *fn + *ln
	*fn = *fn + *ln
	return fn
	//return fulln
}
