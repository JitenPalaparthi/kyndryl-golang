package main

func main() {

	// 1- example

	str := "Hello"
	for _, v := range str {
		defer print(string(v))
	}
	println()

	//2- Example
	a := new(int)
	*a = 10
	defer println("defer exec a:", *a)
	*a = 20
	println("normal exec a:", *a)
}

/*
o
l
l
e
H
*/
