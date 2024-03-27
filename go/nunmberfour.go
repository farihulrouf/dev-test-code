package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)

const (
	baseURL = "https://api.data.gov.sg/v1/graduates"
)

var (
	concurrentLimit = flag.Int("concurrent_limit", 2, "Limit of concurrent processes")
	outputDir       = flag.String("output", ".", "Directory to save CSV files")
)

type Graduate struct {
	Year  int
	Major string
	Count int
}

func fetchData(year int, ch chan<- Graduate) {
	url := fmt.Sprintf("%s?year=%d", baseURL, year)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data for year", year, ":", err)
		return
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	reader.FieldsPerRecord = -1

	// Skip header
	_, _ = reader.Read()

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading record for year", year, ":", err)
			continue
		}

		major := record[0]
		count, _ := strconv.Atoi(record[1])

		ch <- Graduate{Year: year, Major: major, Count: count}
	}
}

func saveToCSV(graduateData []Graduate, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"Year", "Major", "Count"})

	// Write data
	for _, g := range graduateData {
		writer.Write([]string{strconv.Itoa(g.Year), g.Major, strconv.Itoa(g.Count)})
	}

	return nil
}

func main() {
	flag.Parse()

	years := []int{2020, 2021, 2022}

	var wg sync.WaitGroup
	ch := make(chan Graduate)

	// Start workers
	for i := 0; i < *concurrentLimit; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for g := range ch {
				filename := fmt.Sprintf("%s/%d.csv", *outputDir, g.Year)
				if err := saveToCSV([]Graduate{g}, filename); err != nil {
					fmt.Println("Error saving CSV:", err)
				}
			}
		}()
	}

	// Fetch data
	for _, year := range years {
		go fetchData(year, ch)
	}

	// Close channel after all data fetched
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Wait for goroutines to finish
	wg.Wait()
	fmt.Println("Data processing completed.")
}
