package shape

type IArea interface {
	Area() float32
}

type IPerimeter interface {
	Perimeter() float32
}

type IWho interface {
	Who() string
}

type IShape interface {
	IArea
	IPerimeter
	IWho
}
