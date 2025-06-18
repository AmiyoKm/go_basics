package channel

import (
	"fmt"
	"sync"
	"time"
)

func Channel2() {
	c := make(chan int, 2)
	c <- 10
	c <- 20
	println(&c)
	println(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1

	for range n {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func Channel3() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci2(c, quit chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case c <- x:
			fmt.Println("I am chosen at random")
			x, y = y, x+y
		case <-quit:
			fmt.Println("QUIT")
			return
		}
	}
}

func Channel4() {
	var wg sync.WaitGroup
	c := make(chan int)
	quit := make(chan int)
	wg.Add(1)
	go func() {
		for range 10 {
			fmt.Println(<-c)
		}
		quit <- 1

	}()
	go fibonacci2(c, quit, &wg)
	wg.Wait()
}

func SelectChannel() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("TICK")
		case <-boom:
			fmt.Println("BOOM")
			return
		default:
			fmt.Println("     .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
