package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// `Rect`는 method.go에 정의되어있기에 go run method.go interface.go 명령어로 실행시켜야한다. main package 에 다 만들면 안되는 eu..
func main() {
	r := Rect{10, 10}
	c := Circle{10}

	showShapes(r, c)

	var x interface {} // Empty interface ( like Object in java )
	x = 1
	x = "Joonhak"
	letsPrint(x) // Joonhak

	_ := x.(int) // Type Assertion -> runtime error ( panic )
}

func showShapes(shapes ...Shape) {
	for _, s := range shapes {
		fmt.Println(s.area(), s.perimeter())
	}
}

func letsPrint(x interface{}) {
	fmt.Println(x)
}