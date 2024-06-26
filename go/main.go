package main

import (
    "fmt"
    "strconv"
    "strings"
	"encoding/csv"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

// A. Convert string to array of strings
func stringToArray(s string) []string {
    arr := make([]string, len(s))
    for i, char := range s {
        arr[i] = string(char)
    }
    return arr
}

// B. Convert string to array of integers (show only numbers)
func stringToIntArray(s string) []int {
    var arr []int
    for _, char := range s {
        if num, err := strconv.Atoi(string(char)); err == nil {
            arr = append(arr, num)
        }
    }
    return arr
}

// C. Convert string to array of strings (show only characters)
func stringToCharArray(s string) []string {
    var arr []string
    for _, char := range s {
        if char >= 'A' && char <= 'Z' {
            arr = append(arr, string(char))
        }
    }
    return arr
}


//number 2
func calculator(input string) bool {
    // Split the input string based on operators
    parts := strings.Split(input, "+-*/:")

    // Initialize result with the first number
    result, err := strconv.Atoi(parts[0])
    if err != nil {
        return false // Invalid input
    }

    // Iterate over remaining parts
    for _, part := range parts[1:] {
        // Extract operator and operand
        operator := input[strings.IndexAny(input, "+-*/:"):strings.IndexAny(input, "+-*/:")+1]
        operand, err := strconv.Atoi(part)
        if err != nil {
            return false // Invalid input
        }

        // Perform arithmetic operation
        switch operator {
        case "+":
            result += operand
        case "-":
            result -= operand
        case "*":
            result *= operand
        case "/":
            if operand == 0 {
                return false // Division by zero
            }
            result /= operand
        case ":":
            if operand == 0 {
                return false // Division by zero
            }
            result %= operand
        }
    }

    // The result should be true
    return result == 1
}



//answer number 3

func generateJSON(name []string, age []int, hobby []string) {
	// Generate JSON for each person
	for i := 0; i < len(name); i++ {
		person := struct {
			Name  string `json:"name"`
			Age   int    `json:"age"`
			Hobby string `json:"hobby"`
		}{
			Name:  name[i],
			Age:   age[i],
			Hobby: hobby[i],
		}

		// Marshal the person struct into JSON
		jsonData, err := json.Marshal(person)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}

		// Print JSON data
		fmt.Println(string(jsonData))
		if i < len(name)-1 {
			fmt.Println(",")
		}
	}
}

//Number four 

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


func fetchData(url string) ([]GraduatesResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var graduates []GraduatesResponse
	err = json.NewDecoder(resp.Body).Decode(&graduates)
	if err != nil {
		return nil, err
	}

	return graduates, nil
}

func processData(graduates []GraduatesResponse) {
	// Initialize channels for processing
	input := make(chan GraduatesResponse, len(graduates))
	output := make(chan map[string]string, len(graduates))
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
		saveToCSV(result)
	}

	<-done
}

func worker(input <-chan GraduatesResponse, output chan<- map[string]string, wg *sync.WaitGroup) {
	defer wg.Done()

	for grad := range input {
		// Process data
		result := map[string]string{
			"year":  strconv.Itoa(grad.Year),
			"major": hash(grad.Major),
			"count": strconv.Itoa(grad.Count),
		}

		// Send result to output channel
		output <- result
	}
}

func hash(s string) string {
	// Simple hash function, just for demonstration
	return strconv.Itoa(len(s))
}

func saveToCSV(data map[string]string) {
	year, _ := strconv.Atoi(data["year"])
	count, _ := strconv.Atoi(data["count"])

	filename := fmt.Sprintf("%s/%d.csv", *outputDir, year)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{data["year"], data["major"], strconv.Itoa(count)}); err != nil {
		log.Println("Error writing to CSV:", err)
	}
}
//end function number four



//number five A
func displayBarChart(data []int) {
	maxValue := findMax(data)

	for i := maxValue; i > 0; i-- {
		line := ""
		for _, value := range data {
			if value >= i {
				line += "| "
			} else {
				line += "  "
			}
		}
		fmt.Println(line)
	}

	// Horizontal axis with data values
	fmt.Println(data)
}

//number five A
func findMax(data []int) int {
	max := data[0]
	for _, value := range data {
		if value > max {
			max = value
		}
	}
	return max
}




//number five B


func displayBarChartB(data []int) {
	maxValue := findMax(data)

	for i := maxValue; i > 0; i-- {
		line := ""
		for _, value := range data {
			if value >= i {
				line += "| "
			} else {
				line += "  "
			}
		}
		fmt.Println(line)
	}

	// Horizontal axis with data values
	fmt.Println(data)
}

//number five B

func findMaxB(data []int) int {
	max := data[0]
	for _, value := range data {
		if value > max {
			max = value
		}
	}
	return max
}

func insertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = key

		// Display visualization step
		displayBarChart(data)
		fmt.Println()
	}
}


func main() {
    inputString := "1AB23C5678D"

    // number 1A.
    resultA := stringToArray(inputString)
    fmt.Println("Result A:", resultA)

    // number 1B.
    resultB := stringToIntArray(inputString)
    fmt.Println("Result B:", resultB)

    // number 1C.
    resultC := stringToCharArray(inputString)
    fmt.Println("Result C:", resultC)
   

    //Number 2

    input := "5+5+5*5:5"
    fmt.Println(calculator(input))



    //answer number 3
    // Define arrays
	names := []string{"brian", "habib", "malik"}
	ages := []int{25, 25, 24}
	hobbies := []string{"hiking", "touring", "traveling"}

	// Call function to generate JSON for each person
	generateJSON(names, ages, hobbies)



	//call function number four 
	flag.Parse()

	// Create the output directory if it doesn't exist
	if err := os.MkdirAll(*outputDir, os.ModePerm); err != nil {
		log.Fatal("Error creating output directory:", err)
	}

	graduates, err := fetchData("https://api.data.gov.sg/v1/graduates")
	if err != nil {
		log.Fatal("Error fetching data from API:", err)
	}

	processData(graduates)

	//end number four



    //number five a
    // Input array
	data := []int{1, 4, 5, 6, 8, 2}

	// Displaying vertical bar chart
	displayBarChart(data)



    //answer number five B
    datab := []int{1, 4, 5, 6, 8, 2}

	// Display original array
	fmt.Println("Original array:")
	displayBarChartB(datab)
	fmt.Println()

	// Sorting array using insertion sort with visualization steps
	fmt.Println("Sorting using Insertion Sort with visualization steps:")
	insertionSort(datab)
}
