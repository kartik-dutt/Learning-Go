package main

import "fmt"

func rec(a int) int {
	if a < 0 {
		return 0
	}
	fmt.Println(a)
	return rec(a - 1)
}

func main() {
	rec(5)
}
