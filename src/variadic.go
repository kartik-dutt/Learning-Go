package main

import "fmt"

func PrintAll(nums ...interface{}) {
	for _, num := range nums {
		fmt.Println(num)
	}
}

func PrintAllOfInt(nums ...int) {
	for _, num := range nums {
		fmt.Println(num)
	}
}

func main() {
	fmt.Println("Hello")
	// A slice must be passed as ...
	arr := []int{1, 2, 3}
	PrintAll(2, 3, arr, "Hello")
	// Either pass a slice or a integer, not both
	PrintAllOfInt(arr...)
}
