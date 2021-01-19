package main

import "fmt"

func increment() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	incr := increment()
	fmt.Println(incr())
	fmt.Println(incr())

	cur := increment()
	fmt.Println(cur())
	fmt.Println(incr())
}
