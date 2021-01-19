package main

import "fmt"

func add_multiply(a, b int) (int, int, float32) {
	return a + b, a * b, float32(a) / float32(b)
}

func main() {
	fmt.Println("Start")
	fmt.Println(add_multiply(2, 3))
}
