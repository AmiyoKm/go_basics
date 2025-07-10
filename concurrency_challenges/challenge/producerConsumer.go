package challenge

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Item struct {
	ID    int
	Value string
}

func producer(id int, buffer chan<- Item, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		item := Item{
			ID:    id*100 + i,
			Value: fmt.Sprintf("Producer-%d-Item-%d", id, i),
		}

		fmt.Printf("Producer %d creating item %d\n", id, item.ID)
		buffer <- item
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func consumer(id int, buffer <-chan Item, wg *sync.WaitGroup) {
	defer wg.Done()

	for item := range buffer {
		fmt.Printf("Consumer %d processing item %d: %s\n", id, item.ID, item.Value)
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	}
}

func ProducerConsumer() {
	const bufferSize = 10
	const numProducers = 3
	const numConsumers = 2

	buffer := make(chan Item, bufferSize)
	var producerWG, consumerWG sync.WaitGroup

	// Start producers
	for i := 1; i <= numProducers; i++ {
		producerWG.Add(1)
		go producer(i, buffer, &producerWG)
	}

	// Start consumers
	for i := 1; i <= numConsumers; i++ {
		consumerWG.Add(1)
		go consumer(i, buffer, &consumerWG)
	}

	// Close buffer when all producers are done
	go func() {
		producerWG.Wait()
		close(buffer)
	}()

	// Wait for all consumers to finish
	consumerWG.Wait()
	fmt.Println("All items processed!")
}
