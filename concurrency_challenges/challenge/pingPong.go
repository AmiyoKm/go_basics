package challenge

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func ping(ch chan int, id int) {
	for {
		time.Sleep(500 * time.Millisecond)
		val := <-ch
		fmt.Printf("Go routine %d pinged %d:\n", id, val)
		ch <- val + 1
	}
}

func PingPong() {
	ch := make(chan int)
	wg.Add(2)

	go ping(ch, 1)
	go ping(ch, 2)

	ch <- 1
	wg.Wait()
}
