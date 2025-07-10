package challenge

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// simulateWork simulates a task that might take variable time
func simulateWork(ctx context.Context, id int, duration time.Duration) <-chan string {
	result := make(chan string, 1)

	go func() {
		defer close(result)

		select {
		case <-time.After(duration):
			result <- fmt.Sprintf("Task %d completed successfully", id)
		case <-ctx.Done():
			result <- fmt.Sprintf("Task %d cancelled: %v", id, ctx.Err())
		}
	}()

	return result
}

// TimeoutChallenge demonstrates various timeout scenarios
func TimeoutChallenge() {
	fmt.Println("=== Timeout Challenge ===")

	// Scenario 1: Task completes within timeout
	fmt.Println("\n1. Task that completes in time:")
	ctx1, cancel1 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel1()

	result1 := simulateWork(ctx1, 1, 1*time.Second)
	fmt.Println(<-result1)

	// Scenario 2: Task times out
	fmt.Println("\n2. Task that times out:")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel2()

	result2 := simulateWork(ctx2, 2, 3*time.Second)
	fmt.Println(<-result2)

	// Scenario 3: Multiple tasks with different timeouts
	fmt.Println("\n3. Multiple tasks with different behaviors:")

	tasks := []struct {
		id       int
		timeout  time.Duration
		workTime time.Duration
	}{
		{3, 2 * time.Second, 1 * time.Second}, // Will complete
		{4, 1 * time.Second, 3 * time.Second}, // Will timeout
		{5, 3 * time.Second, 2 * time.Second}, // Will complete
	}

	for _, task := range tasks {
		ctx, cancel := context.WithTimeout(context.Background(), task.timeout)
		result := simulateWork(ctx, task.id, task.workTime)
		fmt.Println(<-result)
		cancel()
	}
}

// RaceConditionChallenge simulates a race between multiple operations
func RaceConditionChallenge() {
	fmt.Println("\n=== Race Condition Challenge ===")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Start multiple competing tasks
	channels := make([]<-chan string, 5)
	for i := 0; i < 5; i++ {
		duration := time.Duration(rand.Intn(4)+1) * time.Second
		channels[i] = simulateWork(ctx, i+1, duration)
	}

	// Use select to get the first result
	completed := 0
	for completed < 5 {
		select {
		case result := <-channels[0]:
			if result != "" {
				fmt.Printf("First channel result: %s\n", result)
				channels[0] = nil
			}
		case result := <-channels[1]:
			if result != "" {
				fmt.Printf("Second channel result: %s\n", result)
				channels[1] = nil
			}
		case result := <-channels[2]:
			if result != "" {
				fmt.Printf("Third channel result: %s\n", result)
				channels[2] = nil
			}
		case result := <-channels[3]:
			if result != "" {
				fmt.Printf("Fourth channel result: %s\n", result)
				channels[3] = nil
			}
		case result := <-channels[4]:
			if result != "" {
				fmt.Printf("Fifth channel result: %s\n", result)
				channels[4] = nil
			}
		case <-ctx.Done():
			fmt.Println("Overall timeout reached!")
			return
		}
		completed++
	}
}
