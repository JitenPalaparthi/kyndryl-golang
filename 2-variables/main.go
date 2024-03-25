package main

import (
	"fmt"
	"reflect"
)

func main() {
	// standard way of declaring variables
	var num1 int = 1231
	var num2 uint8 = 123
	var num3 float32 = 3434.343
	var num4 float64 = 32432423.4234232343
	var str1 string = "Hello World"

	var (
		rune1 rune = 'A'
		rune2 rune = '池'
	)

	var byte1 byte = 123
	var byte2, byte3 uint8 = 134, 34

	fmt.Println(num1, num2, num3, num4, str1, rune1, rune2, byte1, byte2, byte3)

	// var declaration based on value
	var num5 = 100 // int
	var str2 = "Hello Kyndryl"
	var float2 = 123.343 // float64
	fmt.Println("flaot2:", float2, "Type of:", reflect.TypeOf(float2))

	// shorthand declaration
	num6 := 123       // int
	float3 := 123.343 // float64
	str3 := "Hello IBM folks"
	//age := 38 /// unit8 , we are using int 8 bytes --.7 bytes are wasted
	var age uint8 = 38
	ok := true

	complex1 := complex(12, 32.34) // complex128
	var r1, i1 float32 = 12.34, 34.56
	complex2 := complex(r1, i1)
	complex3 := 32.45 + 43i // compiler understands it as a complex number.
	fmt.Println(complex3, complex2, complex1, ok, age, str3, float3, num6, float2, str2, num5)
	fmt.Println(rune1, rune2)
}

/* numbers
int,uint, int8,int16,int32,int64, uint4,uint14,uint32,uint64
float32, float64
uintptr
rune -->int32
byte -->unit8
*/
// strings --> string

// boolean --> bool
// complex -- complex64 , complexc128
// interface{} or any
