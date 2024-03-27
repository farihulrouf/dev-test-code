package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

// GraduatesResponse represents the structure of the response from the API
type GraduatesResponse struct {
	Year  int    `json:"year"`
	Major string `json:"major"`
	Count int    `json:"count"`
}

var (
	concurrentLimit = flag.Int("concurrent_limit", 2, "concurrent limit for fetching data")
	outputDir       = flag.String("output", "./csv", "output directory for CSV files")
)

func main() {
	flag.Parse()

	// Create the output directory if it doesn't exist
	if err := os.MkdirAll(*outputDir, os.ModePerm); err != nil {
		log.Fatal("Error creating output directory:", err)
	}

	// Fetch data from the API
	resp, err := http.Get("https://api.data.gov.sg/v1/graduates")
	if err != nil {
		log.Fatal("Error fetching data from API:", err)
	}
	defer resp.Body.Close()

	// Decode JSON response
	var graduates []GraduatesResponse
	if err := json.NewDecoder(resp.Body).Decode(&graduates); err != nil {
		log.Fatal("Error decoding JSON response:", err)
	}

	// Initialize channels for processing
	input := make(chan GraduatesResponse, len(graduates))
	output := make(chan map[string]int, len(graduates))
	done := make(chan bool)

	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < *concurrentLimit; i++ {
		wg.Add(1)
		go worker(input, output, &wg)
	}

	// Feed data to input channel
	go func() {
		for _, grad := range graduates {
			input <- grad
		}
		close(input)
	}()

	// Collect results from output channel
	go func() {
		wg.Wait()
		close(output)
		done <- true
	}()

	// Process results and save to CSV files
	for result := range output {
		year := result["year"]
		major := result["major"]
		count := result["count"]

		filename := fmt.Sprintf("%s/%d.csv", *outputDir, year)
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println("Error opening file:", err)
			continue
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		if err := writer.Write([]string{strconv.Itoa(year), major, strconv.Itoa(count)}); err != nil {
			log.Println("Error writing to CSV:", err)
		}
	}

	<-done
}

func worker(input <-chan GraduatesResponse, output chan<- map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()

	for grad := range input {
		// Process data
		result := map[string]int{
			"year":  grad.Year,
			"major": hash(grad.Major), // You can implement a hash function for majors if needed
			"count": grad.Count,
		}

		// Send result to output channel
		output <- result
	}
}

// hash function for majors (placeholder)
func hash(s string) int {
	return len(s) // Placeholder hash function
}
