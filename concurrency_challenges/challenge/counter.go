package challenge

import (
	"fmt"
	"sync"
)

var count int = 0

func add(mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	count++
	mu.Unlock()
}

func Counter() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	for range 100000 {
		wg.Add(1)
		go add(&mu , &wg)
	}

	wg.Wait()
	fmt.Println("Count : ", count)
}
