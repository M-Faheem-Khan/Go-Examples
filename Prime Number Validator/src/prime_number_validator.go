/**
 * @author Muhammad Faheem Khan
 * @email faheem5948@gmail.com
 * @create date 2019-07-17 23:19:24
 * @modify date 2019-07-17 23:19:24
 * @desc Checks if a number is prime or not
 */

package main

// imports
import "fmt"
import "math"

// main function
func main() {
	// printing the about line
	fmt.Println("Prime number validator")
	// Declaring the variable to store the users number in
	var number int
	fmt.Print("Number: ")
	fmt.Scanf("%f", &number)
	// printing if the numbers are prime or not
	if validateNumber(number) {
		fmt.Println("Number is Prime")
	} else {
		fmt.Println("Number is not Prime")
	}
}

// validates if the passed number is prime or not
func validateNumber(number int) bool {
	// getting the half of the number
	// we have to convert number to float64 because that what is required by math.Float
	rangeMax := math.Floor(float64(number) / float64(2))
	// looping through to see if the number if prime or not
	for i := 2; i <= int(rangeMax); i++ {
		// return false if divisible
		if number%i == 0 {
			return false
		}
	}
	// if the number if more than 1 -> true else false
	return number > 1
}
