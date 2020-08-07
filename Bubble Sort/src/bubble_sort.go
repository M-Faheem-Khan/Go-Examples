// Description: This programs sorts an given array of integers into ascending order using Bubble Sort
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("=== Bulbble Sort ===\n") // header
	// Declaring + init variable
	index := 0
	unSortedSlice := make([]int, 0) // slice to hold integer

	// asking the user how many integers they would like to sort
	fmt.Printf("How many integers would you like to add[2 - 10]? ") // prompt
	loops := readInteger()                                          // reading integer

	// increment index until  loop-1 is reached
	for index < loops {
		fmt.Printf(": ")                                     // prompt
		unSortedSlice = append(unSortedSlice, readInteger()) // appending integer to slice
		index++                                              // incrementing index
	}

	fmt.Println("UnSorted: ", unSortedSlice) // printing unsorted slice
	BubbleSort(&unSortedSlice)               // sorting slice using bubble sort
	fmt.Println("Sorted:   ", unSortedSlice) // printing sorted slice

}

// BubbleSort sorts a given slice using Bubble Sort
func BubbleSort(s *[]int) {
	for l := 0; l < len(*s)-1; l++ {
		for i := range *s {
			if (*s)[i] > (*s)[l+1] {
				temp := (*s)[l+1]
				(*s)[l+1] = (*s)[i]
				(*s)[i] = temp
			}
		}
	}
}

// readInteger: Reads integers from cli
func readInteger() int {
	reader := bufio.NewReader(os.Stdin) // creating reader to read from stdin

	// Keep asking until the user enters a valid integer
	for {
		userInput, err := reader.ReadString('\n') // reading until '\n'

		// Possible Error: Unable to Read from Stdin
		if err != nil {
			continue
		}

		userInput = userInput[:len(userInput)-2] // removing '\r\n'
		userInputInteger, err := strconv.Atoi(userInput)

		// Error: Invalid Integer
		if err != nil {
			continue
		}
		return userInputInteger
	}
}
