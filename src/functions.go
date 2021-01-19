package main

import "fmt"

func add(a int, b int) {
	fmt.Println(a + b)
}

func addNos(a, b int, c, d string) int {
	fmt.Println(c + d)
	return a + b
}

func main() {
	add(2, 3)
	c := addNos(2, 3, "Hello", "!")
	fmt.Println(c)
}
