package challenge

import (
	"fmt"
	"strings"
	"time"
)

// Data represents data flowing through the pipeline
type Data struct {
	ID    int
	Text  string
	Value int
}

// Stage 1: Generate data
func generateData() <-chan Data {
	output := make(chan Data)

	go func() {
		defer close(output)

		texts := []string{
			"hello world",
			"go programming",
			"concurrency patterns",
			"pipeline example",
			"channel communication",
		}

		for i, text := range texts {
			data := Data{
				ID:   i + 1,
				Text: text,
			}
			fmt.Printf("🔄 Generated: %+v\n", data)
			output <- data
			time.Sleep(200 * time.Millisecond)
		}
	}()

	return output
}

// Stage 2: Transform text to uppercase and calculate length
func transformData(input <-chan Data) <-chan Data {
	output := make(chan Data)

	go func() {
		defer close(output)

		for data := range input {
			// Transform the data
			data.Text = strings.ToUpper(data.Text)
			data.Value = len(data.Text)

			fmt.Printf("🔄 Transformed: %+v\n", data)
			output <- data
			time.Sleep(150 * time.Millisecond)
		}
	}()

	return output
}

// Stage 3: Filter data (only keep items with Value > 10)
func filterData(input <-chan Data) <-chan Data {
	output := make(chan Data)

	go func() {
		defer close(output)

		for data := range input {
			if data.Value > 10 {
				fmt.Printf("✅ Filtered (kept): %+v\n", data)
				output <- data
			} else {
				fmt.Printf("❌ Filtered (dropped): %+v\n", data)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	return output
}

// Stage 4: Process final data
func processData(input <-chan Data) <-chan string {
	output := make(chan string)

	go func() {
		defer close(output)

		for data := range input {
			result := fmt.Sprintf("FINAL: ID=%d, TEXT='%s', LENGTH=%d",
				data.ID, data.Text, data.Value)

			fmt.Printf("🎯 Processed: %s\n", result)
			output <- result
			time.Sleep(100 * time.Millisecond)
		}
	}()

	return output
}

func PipelineChallenge() {
	fmt.Println("=== Pipeline Challenge ===")
	fmt.Println("Data Flow: Generate → Transform → Filter → Process")
	fmt.Println()

	// Create the pipeline
	stage1 := generateData()
	stage2 := transformData(stage1)
	stage3 := filterData(stage2)
	stage4 := processData(stage3)

	// Collect results
	var results []string
	for result := range stage4 {
		results = append(results, result)
	}

	fmt.Println("\n📋 Final Results:")
	for i, result := range results {
		fmt.Printf("%d. %s\n", i+1, result)
	}

	fmt.Printf("\nTotal items processed: %d\n", len(results))
}

// Advanced: Pipeline with error handling
type ProcessResult struct {
	Data  Data
	Error error
}

func generateDataWithErrors() <-chan Data {
	output := make(chan Data)

	go func() {
		defer close(output)

		// Include some "problematic" data
		items := []struct {
			id   int
			text string
		}{
			{1, "valid data"},
			{2, ""}, // Empty string - might cause issues
			{3, "another valid item"},
			{4, "short"}, // Short text
			{5, "this is a longer text item"},
		}

		for _, item := range items {
			data := Data{ID: item.id, Text: item.text}
			output <- data
			time.Sleep(300 * time.Millisecond)
		}
	}()

	return output
}

func processWithErrorHandling(input <-chan Data) <-chan ProcessResult {
	output := make(chan ProcessResult)

	go func() {
		defer close(output)

		for data := range input {
			result := ProcessResult{Data: data}

			// Simulate processing that might fail
			if data.Text == "" {
				result.Error = fmt.Errorf("empty text for ID %d", data.ID)
			} else {
				data.Text = strings.ToUpper(data.Text)
				data.Value = len(data.Text)
				result.Data = data
			}

			output <- result
			time.Sleep(200 * time.Millisecond)
		}
	}()

	return output
}

func PipelineWithErrorHandling() {
	fmt.Println("\n=== Pipeline with Error Handling ===")

	stage1 := generateDataWithErrors()
	stage2 := processWithErrorHandling(stage1)

	successCount := 0
	errorCount := 0

	for result := range stage2 {
		if result.Error != nil {
			fmt.Printf("❌ Error: %v\n", result.Error)
			errorCount++
		} else {
			fmt.Printf("✅ Success: %+v\n", result.Data)
			successCount++
		}
	}

	fmt.Printf("\nSummary: %d successful, %d errors\n", successCount, errorCount)
}
