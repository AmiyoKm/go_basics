package main

import (
	"fmt"
	"sync"
	"syscall"
	"time"
)

func sayHello(wg *sync.WaitGroup, count map[int]int) {
	defer wg.Done()
	time.Sleep(50 * time.Microsecond)
	fmt.Printf("Goroutine running on OS thread: %d\n", syscall.Gettid())
	if _, ok := count[syscall.Gettid()]; !ok {
		count[syscall.Gettid()] = 1
	} else {
		count[syscall.Gettid()] += 1
	}

}

func main() {
	var wg sync.WaitGroup
	num := 200
	wg.Add(num)

	count := make(map[int]int)

	for i := 0; i < num; i++ {
		go sayHello(&wg, count)
	}
	wg.Wait()
	for key, val := range count {
		fmt.Println(key, " : ", val)
	}
}
