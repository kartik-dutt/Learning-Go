package main

import "fmt"

func main() {
	const n = 32
	fmt.Println(n)

	// Constant expressions perform arithmetic with arbitrary precision.
	const d = 3e20 / n
	fmt.Println(d)
}
