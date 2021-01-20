package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Start Program")
	// Ways to declare a struct.
	a := person{name: "Kartik", age: 21}
	fmt.Println(a)
	fmt.Println(&a)
	newPerson := func(name string, age int) *person {
		return &person{name: name, age: age}
	}
	b := newPerson("Kartik", 21)
	fmt.Println(b)
}
