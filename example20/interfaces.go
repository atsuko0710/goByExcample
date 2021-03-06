package main

import (
	"fmt"
	"math"
)

type geometry interface {
    area() float64
}


type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
}

func main() {
	r := rect{height: 3, width: 4}
	measure(r)

	c := circle{radius: 2}
	measure(c)
}
