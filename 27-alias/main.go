package main

func main() {

	var num1 integer
	var num2 int = 100

	num1 = num2 // this is not converting to another type.. Both the types are same with a different name
	println(num1, num2)

}

// alias

type integer = int // in all ways int and integer are same
