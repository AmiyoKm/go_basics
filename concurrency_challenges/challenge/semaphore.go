package challenge

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Semaphore limits concurrent access to a resource
type Semaphore struct {
	permits chan struct{}
}

func NewSemaphore(maxConcurrency int) *Semaphore {
	return &Semaphore{
		permits: make(chan struct{}, maxConcurrency),
	}
}

func (s *Semaphore) Acquire() {
	s.permits <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.permits
}

// Database simulates a database with limited connections
type Database struct {
	semaphore *Semaphore
	name      string
}

func NewDatabase(name string, maxConnections int) *Database {
	return &Database{
		semaphore: NewSemaphore(maxConnections),
		name:      name,
	}
}

func (db *Database) Query(userID int, query string) {
	fmt.Printf("User %d waiting for DB connection to %s...\n", userID, db.name)

	// Acquire a connection (blocks if all connections are in use)
	db.semaphore.Acquire()
	defer db.semaphore.Release()

	fmt.Printf("User %d connected to %s, executing: %s\n", userID, db.name, query)

	// Simulate query execution time
	queryTime := time.Duration(rand.Intn(1000)+500) * time.Millisecond
	time.Sleep(queryTime)

	fmt.Printf("User %d finished query on %s (took %v)\n", userID, db.name, queryTime)
}

func SemaphoreChallenge() {
	fmt.Println("=== Semaphore Challenge ===")
	fmt.Println("Database with max 3 concurrent connections")

	// Create a database that allows only 3 concurrent connections
	db := NewDatabase("ProductionDB", 3)

	var wg sync.WaitGroup

	// Simulate 10 users trying to access the database concurrently
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(userID int) {
			defer wg.Done()

			queries := []string{
				"SELECT * FROM users",
				"UPDATE user_profile SET last_login = NOW()",
				"INSERT INTO activity_log VALUES (...)",
				"DELETE FROM temp_data WHERE created < ...",
			}

			query := queries[rand.Intn(len(queries))]
			db.Query(userID, query)
		}(i)
	}

	wg.Wait()
	fmt.Println("All database operations completed!")
}

// DownloadManager demonstrates semaphore for limiting concurrent downloads
func DownloadManager() {
	fmt.Println("\n=== Download Manager ===")
	fmt.Println("Max 2 concurrent downloads")

	downloadSemaphore := NewSemaphore(2)

	files := []string{
		"large-file-1.zip",
		"document.pdf",
		"video.mp4",
		"archive.tar.gz",
		"database-backup.sql",
		"images.zip",
	}

	var wg sync.WaitGroup

	for i, file := range files {
		wg.Add(1)
		go func(fileID int, fileName string) {
			defer wg.Done()

			fmt.Printf("Starting download %d: %s\n", fileID+1, fileName)

			downloadSemaphore.Acquire()
			defer downloadSemaphore.Release()

			fmt.Printf("⬇️  Downloading %s...\n", fileName)

			// Simulate download time
			downloadTime := time.Duration(rand.Intn(2000)+1000) * time.Millisecond
			time.Sleep(downloadTime)

			fmt.Printf("✅ Completed %s (took %v)\n", fileName, downloadTime)
		}(i, file)
	}

	wg.Wait()
	fmt.Println("All downloads completed!")
}
