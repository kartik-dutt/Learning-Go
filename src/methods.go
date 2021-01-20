package main

import "fmt"

type circle struct {
	radius float32
}

// Function inside a circle structure.
// They can't be declared inside the structure.
func (c circle) area() float32 {
	return 3.14 * 2 * c.radius * c.radius
}

func (c *circle) perimeter() float32 {
	return 3.14 * 2 * c.radius
}

func main() {
	c := circle{radius: 1.0 / 3.14}
	fmt.Println(c.area())
	fmt.Println(c.perimeter())
}
