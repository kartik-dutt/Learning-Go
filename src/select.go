package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		fmt.Println("Hello")
		time.Sleep(3 * time.Second)
		c1 <- "one"
	}()

	go func() {
		fmt.Println("Hello2")
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	fmt.Println("Starts here!")
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Recieved", msg1)
		case msg2 := <-c2:
			fmt.Println("recieved", msg2)
		}
	}

	// This is how concurrency work imo :
	// -----------
	//main-------------
	// |		|			|
	// |		 th1		th2
	// | 		 |		|
	// | <-close 	|
	// |           <-close
}
