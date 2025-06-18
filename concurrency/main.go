package main

import (
	"concurrency/channel"
	"concurrency/mutex"
	"fmt"
	"sync"
	"time"
)

func say(s string, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	for i := range 10 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)
	}
}

func main() {
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go say("Hello", &wg)
	// say("World", nil)
	// //go say("Bruh" , &wg)
	// wg.Wait()
	// channel.ChannelExec()
	// channel.Channel2()
	// channel.Channel3()
	//channel.Channel4()
	channel.SelectChannel()
	mutex.GoMutex()
}
