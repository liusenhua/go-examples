package main

import "fmt"
import "time"

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		for i := 0; i < 1000000; i++ {
		}
		time.Sleep(1e3)
		results <- j * 2
	}
}

const (
	MAX_JOBS int = 10000
	NUM_WORKERS int = 16
)
func main() {

	jobs := make(chan int, MAX_JOBS)
	results := make(chan int, MAX_JOBS)

	for w := 1; w <= NUM_WORKERS; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= MAX_JOBS; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= MAX_JOBS; a++ {
		<-results
	}
}
