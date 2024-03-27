package main

import (
	"demo/shape"
	"demo/shape/cuboid"
	"demo/shape/rect"
	"demo/shape/square"
	"fmt"
)

func main() {

	r1 := &rect.Rect{L: 12.23, B: 13.45}
	var s1 square.Square = 25.3
	r2 := &rect.Rect{L: 1223.23, B: 1323.45}
	var s2 square.Square = 2235.3
	c1 := &cuboid.Cuboid{L: 12.23, B: 13.45, H: 12.45}
	cr1 := Circle{R: 12.45}

	shapes := make([]shape.IShape, 6)
	shapes[0] = r1
	shapes[1] = s1
	shapes[2] = r2
	shapes[3] = s2
	shapes[4] = c1
	shapes[5] = &cr1
	// var iarea shape.IArea
	// iarea = r1
	// a1 := iarea.Area()
	// fmt.Println("Rect Area", a1)
	// iarea = s1
	// a2 := iarea.Area()
	// fmt.Println("Square Area", a2)
	for _, v := range shapes {
		Shape(v)
	}
}

func Shape(ishape shape.IShape) {
	a := ishape.Area()
	fmt.Println("Area of", ishape.Who(), a)
	p := ishape.Perimeter()
	fmt.Println("Perimeter of", ishape.Who(), p)
	fmt.Println()
}

const (
	PI = 3.14
)

type Circle struct {
	R float32
}

func (c *Circle) Area() float32 {
	return PI * c.R * c.R
}

func (c *Circle) Perimeter() float32 {
	return 2 * PI * c.R
}

func (c *Circle) Who() string {
	return "Circle"
}

// Dependency Injection

// class Rect: IArea
// class Rect implemnts IArea
