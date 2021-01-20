package main

import "fmt"

func set(buffer chan<- string, msg string) {
	buffer <- msg
}

func transfer(src <-chan string, dest chan<- string) {
	dest <- (<-src)
}

func main() {
	buffer1 := make(chan string, 1)
	buffer2 := make(chan string, 1)

	set(buffer1, "Kartik")
	transfer(buffer1, buffer2)
	fmt.Println(<-buffer2)
}
