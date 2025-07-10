package challenge

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// squareWorker processes numbers and returns their squares
func squareWorker(id int, input <-chan int, output chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range input {
		result := num * num
		fmt.Printf("Worker %d: %d^2 = %d\n", id, num, result)
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		output <- result
	}
}

// fanOut distributes work to multiple workers
func fanOut(input <-chan int, numWorkers int) <-chan int {
	output := make(chan int)
	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go squareWorker(i, input, output, &wg)
	}

	// Close output when all workers are done
	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

// fanIn merges multiple channels into one
func fanIn(channels ...<-chan int) <-chan int {
	merged := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				merged <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func FanInFanOut() {
	// Create input channel and send numbers
	input := make(chan int, 10)
	go func() {
		defer close(input)
		for i := 1; i <= 20; i++ {
			input <- i
		}
	}()

	// Fan-out to 3 workers
	output1 := fanOut(input, 3)

	// Create another set of workers for demonstration
	input2 := make(chan int, 5)
	go func() {
		defer close(input2)
		for i := 21; i <= 25; i++ {
			input2 <- i
		}
	}()
	output2 := fanOut(input2, 2)

	// Fan-in results from both sets
	results := fanIn(output1, output2)

	// Collect all results
	var allResults []int
	for result := range results {
		allResults = append(allResults, result)
	}

	fmt.Printf("All results: %v\n", allResults)
	fmt.Printf("Total results: %d\n", len(allResults))
}
