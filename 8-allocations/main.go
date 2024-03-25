package main

import "fmt"

var num int // This is a global variable. It is stored in data segment
// there are many isuse with global variables. data rases. Multiple goroutines may
// use them concurrently by an accident or deliberately.

func main() {
	// constans are evaluated at compile time
	var a1 int = 60
	var ok = true
	const C1 int = 7
	const C2 = 24 * C1 // can do this. Becase C1 is also a constant
	//const C3 = C2 * a1 // this is not allowed. a1 is not constant.
	// a1 will not be created until you run your application
	fmt.Println(a1, C1, C2)
	{
		str1 := "Hello Kyndryl folks!" // assume this is allocated in Heap. Assume means in Go allocation in stack or heap is decided by Runtime.
		var b1 int = 70
		{
			var c1 = a1 + b1
			fmt.Println(c1)
		}
		fmt.Println(str1)
		// free str1 .. You dont need to do this. This is done by GC
	}
	fmt.Println(ok)

}

/*---------------------------------------------------------------------------------------
- application contains a process.
- process contains some memory
- memory is divided logically into code segment, data segment, stack and heap
- All global variables are stored in data segment
- All constants are stored in code segment
- In general all variables are stored in Stack segment.Allocations and deallocations
	- are based on the scope operator
- Few types are stored in Heap memory.Heap is managed by GC.Allocations and deallocations
	- are managed by runtime and GC
---------------------------------------------------------------------------------------*/
