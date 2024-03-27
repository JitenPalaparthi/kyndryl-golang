package square

type Square float32

func (s Square) Area() float32 {
	return float32(s * s)
}

func (s Square) Perimeter() float32 {
	return float32(4 * s)
}

func (r Square) Who() string { // this is exported or unrestricted or similar to public
	return "Square"
}
