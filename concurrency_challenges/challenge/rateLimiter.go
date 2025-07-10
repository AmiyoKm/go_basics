package challenge

import (
	"fmt"
	"time"
)

// RateLimiter implements a token bucket rate limiter
func RateLimiter() {
	// Create a rate limiter that allows 3 requests per second
	limiter := make(chan struct{}, 3)

	// Fill the bucket initially
	for i := 0; i < 3; i++ {
		limiter <- struct{}{}
	}

	// Refill the bucket every second
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for range ticker.C {
			select {
			case limiter <- struct{}{}:
				// Token added
			default:
				// Bucket is full, skip
			}
		}
	}()

	// Simulate 10 requests
	for i := 1; i <= 10; i++ {
		go func(reqID int) {
			<-limiter // Wait for a token
			fmt.Printf("Request %d processed at %s\n", reqID, time.Now().Format("15:04:05"))
			time.Sleep(100 * time.Millisecond) // Simulate work
		}(i)
	}

	time.Sleep(5 * time.Second) // Wait for requests to complete
}
