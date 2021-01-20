package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
		wg.Done()
	}
}

func main() {
	const numJobs = 6
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup
	for w := 1; w <= 6; w++ {
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		wg.Add(1)
		jobs <- j
	}

	wg.Wait()
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
