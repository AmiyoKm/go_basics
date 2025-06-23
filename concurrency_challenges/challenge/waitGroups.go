package challenge

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d is working \n", id)
}

func WG() {
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("ALL workers done")
}
