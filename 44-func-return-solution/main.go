package main

import "fmt"

func main() {

	//var Add, Sub, Mul CalcFunc // create variables of CalcFunc
	var Add CalcFunc = func(nums ...any) func() any {
		return func() any {
			var result int
		outer:
			for _, v := range nums {
				switch v.(type) {
				case int:
					result = v.(int) + result
				default:
					fmt.Println("unsupported type")
					break outer
				}
			}
			return result
		}
	}

	f1 := Add(10, 12, 14, 16, 18, 20)
	fmt.Println("Add", f1())

	var Sub CalcFunc = Substact
	f2 := Sub(10, 12, 14, 16, 18, 20)
	fmt.Println("Sub", f2())
	var Mul CalcFunc = Multiply
	f3 := Mul(10, 12, 14, 16, 18, 20)
	fmt.Println("Mul", f3())

}

type CalcFunc func(nums ...any) func() any

func Substact(nums ...any) func() any {
	return func() any {
		var result int
	outer:
		for _, v := range nums {
			switch v.(type) {
			case int:
				result = v.(int) - result
			default:
				fmt.Println("unsupported type")
				break outer
			}
		}
		return result
	}
}

func Multiply(nums ...any) func() any {
	return func() any {
		var result int = 1
	outer:
		for _, v := range nums {
			switch v.(type) {
			case int:
				if v.(int) != 0 {
					result = v.(int) * result
				}
			default:
				fmt.Println("unsupported type")
				break outer
			}
		}
		return result
	}
}

// extend these functions to use for more types. Any must accept not only in but all numbers
