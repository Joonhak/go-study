package main

import (
	"basic/mypackage"
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

func (c Circle) Area() float64 {
	return math.Pow(c.Radius, 2) * math.Pi
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	r := mypackage.Rect{10, 10}
	c := Circle{10}

	showShapes(r, c)

	var x interface{} // Empty interface ( like Object in java )
	x = 1
	x = "Joonhak"
	letsPrint(x) // Joonhak

	_ = x.(int) // Type Assertion -> runtime error ( panic )
}

func showShapes(shapes ...Shape) {
	for _, s := range shapes {
		fmt.Println(s.Area(), s.Perimeter())
	}
}

func letsPrint(x interface{}) {
	fmt.Println(x)
}
