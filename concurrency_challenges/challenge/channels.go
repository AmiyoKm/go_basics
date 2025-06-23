package challenge

import (
	"fmt"
	"sync"
)

func Send(c chan string, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 5 {
		c <- fmt.Sprintf("Goroutine %d: Ping %d\n", id, i)
	}
}

func Channels() {
	ch := make(chan string, 2)
	var wg sync.WaitGroup

	wg.Add(2)
	go Send(ch, 1, &wg)
	go Send(ch, 2, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Print(v)
	}
}
