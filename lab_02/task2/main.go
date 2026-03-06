package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

type Triangle struct {
	Base   float64
	Height float64
	SideA  float64
	SideB  float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
func (t Triangle) Perimeter() float64 {
	return t.Base + t.SideA + t.SideB
}
func main() {
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 10, Height: 4},
		Triangle{Base: 4, Height: 6, SideA: 3, SideB: 5},
	}
	for _, s := range shapes {
		fmt.Println(s, s.Area(), s.Perimeter())
	}
}
