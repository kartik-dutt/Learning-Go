package main

import (
	"fmt"
	"time"
)

func printString(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s," : " ,i)
	}
}

func main() {
	printString("direct call")
	go printString("go routine")

	go printString("Some Text")
	fmt.Println("The go routine works on anonymous functions as well. It runs asynchronously")

	// Wait group or time.sleep must be used.
	time.Sleep(time.Second)
}
