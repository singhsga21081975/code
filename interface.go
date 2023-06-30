package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func main() {
	rectangle := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}

	var shape Shape
	shape = &rectangle

	fmt.Printf("Rectangle - Area: %.2f, Perimeter: %.2f\n", shape.Area(), shape.Perimeter())

	shape = &circle // Use a pointer to Circle
	fmt.Printf("Circle - Area: %.2f, Perimeter: %.2f\n", shape.Area(), shape.Perimeter())
}
