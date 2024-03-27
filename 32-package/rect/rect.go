package rect

import "fmt"

type Rect struct { // Rect is exported type
	L, B float32 // L and B are exported fields
	a, p float32 // a and p are unexported/restricted fields.
}

func (r *Rect) Area() float32 {
	r.a = r.L * r.B
	return r.a
}

func (r *Rect) Perimeter() float32 {
	r.p = 2 * (r.L + r.B)
	return r.p
}

func (r *Rect) display() { // this is unexported or restricted or similar to private
	fmt.Println("Length:", r.L)
	fmt.Println("Bredth:", r.B)
}

func (r *Rect) Display() { // this is exported or unrestricted or similar to public
	r.display()
}
