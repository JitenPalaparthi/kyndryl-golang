package main

import "fmt"

func main() {

	var num11 uint8 = 128

	// var num2 int = num1 //it works in many , but not in Go

	var num12 int = int(num11) // not ok in go: (int)num1

	fmt.Println(num11, num12)

	fmt.Printf("num11: %d type of num11:%T, num12:%v type of num12: %T\n", num11, num11, num12, num12)

	var (
		num1 int     = 123123
		num2 int8    = 123
		num3 float32 = 232.343
		num4 float64 = 3232.3433443
		num5 int64   = 34234234234234
		num6 any     = 231231
		num7 any     = num4
	)

	sum := float64(num1) + float64(num2) + float64(num3) + num4 + float64(num5) + float64(num6.(int)) + num7.(float64)
	fmt.Printf("Sum:%.5f\n", sum)

	var num10 int = 12312
	var any1 any
	//var any2 any // nil

	any1 = num10         // similar to boxing in other programming languages
	num111 := any1.(int) // unboxing in other programming languages
	any1 = num2
	num112 := any1.(int8)
	// any1 = num3
	// any1 = num4
	// any1 = num5
	// any1 = "Hello"
	// any1 = false
	fmt.Println(num111, num112)

	// type inference

	var num114 int     // 0
	var str2 string    //""
	var float3 float64 //0
	var ok bool        // false
	var any3 any       // nil
	fmt.Println(num114, str2, float3, ok, any3)

}
