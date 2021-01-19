package main

import "fmt"

func main() {
	ht := make(map[int]int)
	for i := 0; i < 3; i++ {
		ht[i] = i * i
	}

	for key, value := range ht {
		fmt.Println(key, " -> ", value)
	}

	arr := make([]string, 3)
	arr[0] = "Hello"
	arr[1] = "World"
	arr[2] = "!"

	for key, idx := range arr {
		fmt.Println(key, idx)
	}
}
