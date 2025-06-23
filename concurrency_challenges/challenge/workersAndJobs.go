package challenge

import (
	"fmt"
	"sync"
	"time"
)

func jobsWorker(id int, wg *sync.WaitGroup, jobs chan int) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(500 * time.Millisecond)
	}
}

func WorkersJobs() {
	var wg sync.WaitGroup
	jobs := make(chan int, 10)

	for w := 1; w <= 5; w++ {
		wg.Add(1)
		go jobsWorker(w, &wg, jobs)
	}

	for i := range 10 {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}
