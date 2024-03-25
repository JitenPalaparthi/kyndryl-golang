package main

import "fmt"

func main() {
	var any1 any = 3434.343
	switch any1.(type) { // type switch
	case int:
		fmt.Println(any1, "is int type")
	case float64:
		fmt.Println(any1, "is float64 type")
	case string:
		fmt.Println(any1, "is string type")
	case nil:
		fmt.Println(any1, "is nil")
	default:
		fmt.Println("not int,float64 or string")
	}
}
