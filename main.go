package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func (rectangle Rectangle) area() int {
	return rectangle.Width * rectangle.Height
}

func main() {
	rect1 := Rectangle{Width: 10, Height: 20}

	fmt.Println(rect1.area())
}
