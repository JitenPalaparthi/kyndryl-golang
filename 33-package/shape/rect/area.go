package rect

func (r *Rect) Area() float32 {
	r.a = r.L * r.B
	return r.a
}
