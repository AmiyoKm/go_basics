package mutex

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	count map[string]int
}

func (c *Counter) inc(key string) {
	c.mu.Lock()
	c.count[key]++
	c.mu.Unlock()
}

func (c *Counter) value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count[key]
}

func GoMutex() {
	c := Counter{count: make(map[string]int)}
	for range 1000 {
		go c.inc("key")
	}
	time.Sleep(time.Second)
	fmt.Println(c.value("key"))
}
