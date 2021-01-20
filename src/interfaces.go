package main

import "fmt"

type shapes interface {
	area() float32
}

type rectangle struct {
	width, height float32
}

// Don't use pointer here.
func (r rectangle) area() float32 {
	return r.width * r.height
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return 3.14 * c.radius * c.radius
}

func GetArea(s shapes) {
	fmt.Println(s.area())
}

func main() {
	fmt.Println("Start")
	c := circle{radius: 3}
	r := rectangle{width: 2, height: 2}
	GetArea(c)
	GetArea(r)
}
