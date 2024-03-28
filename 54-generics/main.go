package main

import (
	"fmt"
	"unsafe"
)

func main() {

	r1 := Rect[int, string]{} //
	r1.L = 101
	r1.B = 102
	fmt.Println(unsafe.Sizeof(r1))
	fmt.Println(r1.Area())

	r2 := new(Rect[float32, string])
	r2.L = 101.234
	r2.B = 102.324
	fmt.Println(unsafe.Sizeof(r2))
	fmt.Println(r2.Area())

	r3 := new(Rect[int16, string])
	r3.L = 101
	r3.B = 102
	//
	fmt.Println(unsafe.Sizeof(r3))
	fmt.Println(r3.Area())

}

type Rect[T Compare, S string] struct {
	L, B T
}

func (r *Rect[T, S]) Area() S { // return S which is string does not make sense but for demo purppse
	//var ret S
	ret := fmt.Sprint(r.L * r.B)
	return S(ret)
}

func (r *Rect[T, S]) Perimeter() T {
	return 2 * (r.L + r.B)
}

type Compare interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float64 | float32
}
