package main

import (
	"fmt"

	"github.com/JitenPalaparthi/mathshapes/shape"
	"github.com/JitenPalaparthi/mathshapes/shape/rect"
	"github.com/JitenPalaparthi/mathshapes/shape/square"
)

func main() {
	shape.What()
	r1 := rect.Rect{L: 1232.213, B: 112331.1231}
	fmt.Println("Area of Rect:", r1.Area())
	fmt.Println("Perimeter of Rect:", r1.Perimeter())

	var s1 square.Square = 25.25
	fmt.Println("Area of Square:", s1.Area())
	fmt.Println("Perimeter of Square:", s1.Perimeter())

}
