package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	i := 1
	for {
		if i >= 4 {
			break
		}

		i++
		fmt.Println(i)
	}

	for i <= 10 {
		fmt.Println(i)
		i++
	}
}
