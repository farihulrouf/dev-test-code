package main

import (
    "fmt"
    "strconv"
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

func main() {
    inputString := "1AB23C5678D"

    // A.
    resultA := stringToArray(inputString)
    fmt.Println("Result A:", resultA)

    // B.
    resultB := stringToIntArray(inputString)
    fmt.Println("Result B:", resultB)

    // C.
    resultC := stringToCharArray(inputString)
    fmt.Println("Result C:", resultC)
}
