package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	lambda := func(str string) {
		messages <- str
	}

	lambda("Messages 1")
	lambda("Message 2")

	// FIFO
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
