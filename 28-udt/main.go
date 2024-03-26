package main

import "fmt"

func main() {

	r1 := &Rect{Length: 12.34, Bredth: 14.56}
	a1 := Area(*r1)
	fmt.Println("Area of r1:", a1)

	a2 := r1.Area()
	fmt.Println("Area of r1:", a2)
	p2 := r1.Perimeter()
	fmt.Println("Perimeter of r1:", p2)
	fmt.Println("Area in r1:", r1.A)
	fmt.Println("Perimeter in r1:", r1.P)
	fmt.Println("Length:", r1.Length)

	a3 := r1.AreaP()
	fmt.Println("Area of r1:", a3)
	p3 := r1.PerimeterP()
	fmt.Println("Perimeter of r1:", p3)
	fmt.Println("Area in r1:", r1.A)
	fmt.Println("Perimeter in r1:", r1.P)

}

type Rect struct {
	Length, Bredth float32
	A, P           float32
}

// Function
func Area(r Rect) float32 {
	return r.Length * r.Bredth
}

// Method
func (r Rect) Area() float32 { // even the receivers have value type and then reference type
	r.A = r.Length * r.Bredth
	return r.A
}

func (myrest Rect) Perimeter() float32 { // generally single letter is given as a name of the receiver. But you can give any valid identifier as a name
	myrest.P = 2 * (myrest.Length + myrest.Bredth)
	return myrest.P
}

// pointer receiver.. is a call by reference
// just it is a pointer receiver , you dont need to create a pointer struct
// automatically same reference is used
func (r *Rect) AreaP() float32 {
	r.A = r.Length * r.Bredth
	return r.A
}

func (r *Rect) PerimeterP() float32 { // generally single letter is given as a name of the receiver. But you can give any valid identifier as a name
	r.P = 2 * (r.Length + r.Bredth)
	return r.P
}
