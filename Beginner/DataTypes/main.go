package main

import "fmt"

// To build binaries --> go build, else go run
func main() {
	var x uint8 = 8
	var y int64 = 1220
	// fmt.Println(x + y) // Causes an error, Nice!
	fmt.Println(int(x) + int(y))
	fmt.Println(x * 2)
	fmt.Println(y * y)

	var maxFloat32 float32
	maxFloat32 = 16777216
	fmt.Println(maxFloat32 == maxFloat32+10)
	var amazing bool = true

	var myName string
	myName = "KD"
	fmt.Println(myName, amazing, " GO")
}
