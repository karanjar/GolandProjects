package newmethods

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * 0.5
}

//interface

type Shape interface {
	Area() float64
	Circumference() float64
}
type Round struct {
	Radius float64
}

func (r Round) Area() float64 {
	return math.Pi * r.Radius * r.Radius
}
func (r Round) Circumference() float64 {
	return math.Pi * r.Radius * 2
}
func Measure(s Shape) {
	fmt.Println(s, "*************************")
	fmt.Println(s.Area())
	fmt.Println(s.Circumference())
}
