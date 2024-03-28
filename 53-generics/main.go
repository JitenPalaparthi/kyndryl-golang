package main

import (
	"fmt"
	"reflect"
)

func main() {
	// to find biggest numbers among integers
	a, b := 347, 344
	c := MaxInts(a, b)
	fmt.Println(c)
	fmt.Println(MaxAny(12, 13))
	fmt.Println(MaxAny(12.4234, 13.23423))
	fmt.Println(MaxAny(12, 13.23423))
	fmt.Println(Max(12.34, 5))
	fmt.Println(Sum(23.343, 433343432234))
	fmt.Println(Sum("23.343", "433343432234"))

}

func MaxAny(i, j any) any { // i and j are numeric types
	if reflect.TypeOf(i) == reflect.TypeOf(j) {
		switch i.(type) {
		case int:
			if i.(int) > j.(int) {
				return i
			}
			return j
		case int8:
			if i.(int8) > j.(int8) {
				return i
			}
			return j
		case int16:
			if i.(int16) > j.(int16) {
				return i
			}
			return j
		case int32:
			if i.(int32) > j.(int32) {
				return i
			}
			return j
		case int64:
			if i.(int64) > j.(int64) {
				return i
			}
			return j
		case uint:
			if i.(uint) > j.(uint) {
				return i
			}
			return j
		case uint8:
			if i.(uint8) > j.(uint8) {
				return i
			}
			return j
		case uint16:
			if i.(uint16) > j.(uint16) {
				return i
			}
			return j
		case uint32:
			if i.(uint32) > j.(uint32) {
				return i
			}
			return j
		case uint64:
			if i.(uint64) > j.(uint64) {
				return i
			}
			return j
		case float32:
			if i.(float32) > j.(float32) {
				return i
			}
			return j
		case float64:
			if i.(float64) > j.(float64) {
				return i
			}
			return j
		}
	}
	return nil
}

func MaxInts(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// func [T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float64 | float32](a T, b T) T {
func Max[T Compare](a T, b T) T {
	// a1, b1 := true, false
	// if a1 > b1 {

	// }
	if a > b {
		return a
	}
	return b
}

type Compare interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float64 | float32
}
type Sumable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float64 | float32 | string
}

func Sum[S Sumable](a, b S) S {
	return a + b
}
