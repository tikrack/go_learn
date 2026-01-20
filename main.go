package main

type A interface {
	area() int
}

type Square struct {
	Width  int
	Height int
}

func (square Square) area() int {
	return square.Width * square.Height
}

type Rectangle struct {
	Width  int
	Height int
}

func (rectangle Rectangle) area() int {
	return rectangle.Width * rectangle.Height
}

func tenMultiply(a A) int {
	return a.area() * 10
}

func main() {
	rect := Rectangle{Width: 10, Height: 10}

	println(tenMultiply(rect))
}
