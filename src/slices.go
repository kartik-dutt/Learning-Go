package main

import "fmt"

func main() {
	// We use builtin make method.
	a := make([]int, 3)
	fmt.Println(a)

	a = append(a, 3, 4, 5)
	fmt.Println(a)

	fmt.Println(a[2:])
	fmt.Println(a[2:4])

	t := make([]int, 3)
	copy(a, t)
	fmt.Println(t)

	twoD := make([][]int, 0)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD = append(twoD, make([]int, innerLen))
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
