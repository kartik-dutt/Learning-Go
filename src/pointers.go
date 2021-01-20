package main

import "fmt"

func pass_by_value(a int) {
	a = 0
}

func pass_by_ref(a *int) {
	*a = 0
}

func main() {
	fmt.Println("Start")
	a := 100
	fmt.Print("Pass by Value changed value of ", a, " to ")
	pass_by_value(a)
	fmt.Println(a)
	pass_by_ref(&a)
	fmt.Println("Pass by ref", a)
}
