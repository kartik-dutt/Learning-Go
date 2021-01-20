package main

import (
	"fmt"
	"sync"
	"time"
)

func myFn(wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	fmt.Println("Finished Executing GoRoutine")
	wg.Done()
}

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Starting worker ", i)
	fmt.Println("Done with ", i)
}

func main() {
	fmt.Println("Go WaitGroups!")
	var wg sync.WaitGroup
	wg.Add(1)
	go myFn(&wg)
	// Select requires channels.
	// Time out wastes extra time.
	// time.Sleep(time.Second)
	wg.Wait()
	fmt.Println("Here!")
	wg.Add(1)
	go myFn(&wg)
	wg.Wait()
	fmt.Println("Finished execution of the program")

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}
