package main

import (
    "fmt"
    "strconv"
    "strings"
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
