package cuboid

type Cuboid struct { // Rect is exported type
	L, B, H float32 // L and B are exported fields
	a, p    float32 // a and p are unexported/restricted fields.
}

func (r *Cuboid) Area() float32 {
	r.a = r.L * r.B * r.H
	return r.a
}

func (r *Cuboid) Perimeter() float32 {
	r.p = 4 * (r.L + r.B + r.H)
	return r.p
}
func (r *Cuboid) Who() string { // this is exported or unrestricted or similar to public
	return "Cuboid"
}
