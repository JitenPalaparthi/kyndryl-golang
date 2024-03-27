package main

import (
	"fmt"
)

func main() {

	map1 := make(map[string]any, 5)

	map1["greet"] = "Hello Kyndryl!"
	map1["sum"] = func(a, b int) int {
		return a + b
	}
	var func1 Func = func(i1, i2 int) int {
		return i1 * i2
	}
	map1["func-type"] = func1

	map1["sub"] = func(a, b int) int {
		return a - b
	}(10, 20)

	for k, v := range map1 {
		//	fmt.Println("Key", k, "-->", reflect.TypeOf(v), "Kind-->", reflect.TypeOf(v).Kind())

		switch v.(type) {
		case string:
			fmt.Println("Key:", k, "-->", v)
		case int:
			fmt.Println("Key:", k, "-->", v)
		case func(int, int) int:
			//r := v.(func(int, int) int)(10, 20)
			r := v.(func(int, int) int)(10, 20)
			//fmt.Println(r)
			fmt.Println("Key:", k, "-->", r)
		case Func:
			r := v.(Func)(10, 20)
			fmt.Println("Key:", k, "-->", r)
		default:
			fmt.Println("unsupported type", k)
		}
	}

}

type Func func(int, int) int
