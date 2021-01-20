package main

import (
	"fmt"
	"time"
)

func main() {
	// Example where web page is downloading, get the data from the server.
	// Add optional parameter 1 i.e. channel buffer size, it will automatically be blocked after it's buffer is full.
	// <-chan ==> Read Only, -> chan ===> Write only.
	done := make(chan bool, 1)
	lambda := func() {
		fmt.Println("Hello world!")
		time.Sleep(2 * time.Second)
		done <- true // Send a value to channel like this.
	}

	go lambda()

	fmt.Println("Do the remaining task.")
	// Go main routine waits untill the value is filed into the done channel.
	// <-done
	// Alternatively ...
	// Exceed buffer size. done <- true
	fmt.Println("Done with first go routine")
}
