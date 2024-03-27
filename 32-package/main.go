package main

import (
	"fmt"
	"shapes/rect"
	"time"
)

func main() {
	fmt.Println(time.Now())
	r1 := new(rect.Rect)
	r1.L = 123.34
	r1.B = 123.56
	fmt.Println("Area of Rect r1:", r1.Area())
	fmt.Println("Perimeter of Rect r1:", r1.Perimeter())
	// r1.display() // cannot call
	r1.Display() // can call
}

// There are 3 types of packages in Go
// 1- Standard packages
// 			GOROOT env .
//			src/ pkg/ bin/
// 2- User defined packages
// 3- Third party packages

/*
- Go does not have access modifiers or specifiers
- Go has a package system
- every new package you create should have a directory.
- The practice is the name of the pacakge is the name of its immediate directory
- In the package any type or identifer starts with Uppercase is considerd as exported type or filed
- Exported is analogous to public

- To import User Defined Packages, the module name is the root directory
*/
