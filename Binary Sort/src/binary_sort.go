/**
 * @author Muhammad Faheem Khan
 * @email faheem5948@gmail.com
 * @create date 2019-07-16 06:16:06
 * @modify date 2019-07-16 06:16:06
 * @desc Simple Binary Sort Implementation in go
 */

package main

// Import Statements
import "fmt"       // I/O
import "math/rand" // Random Numbers

func main() {
	// printing the about program line
	fmt.Println("Simple Implementation of Binary Sort")
	// defining array
	var sort_array [10]int
	// populating sort_array with random number
	for i := 0; i < 10; i++ {
		sort_array[i] = rand.Intn(100)
	}
	// Printing the array values
	fmt.Println(sort_array)
	// sorting array
	binarySort(sort_array)
}

// this function sort the array
func binarySort(arr [10]int) {
	// first loop
	for k := 0; k < len(arr); k++ {
		// second loop to go through all indexes in array
		for l := 0; l < (len(arr) - 1); l++ {
			// temp variables
			var a = arr[l]
			var b = arr[l+1]
			// switching index places
			if a > b {
				arr[l] = b
				arr[l+1] = a
			}
		}
	}
	// printing the arry
	fmt.Println(arr)
}
