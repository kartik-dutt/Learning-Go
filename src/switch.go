package main

import (
	"fmt"
	"time"
)

func main() {
	// Single switch case is same as c++.
	// multiple condition.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	case time.Tuesday:
		fmt.Println("It's Tuesday")
	default:
		fmt.Println("It's a weekday")
	}
	// Type (can only be used inside switch)
	getType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	getType("hello")
}
