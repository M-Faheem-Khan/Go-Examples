/**
 * @author Muhammad Faheem Khan
 * @email faheem5948@gmail.com
 * @create date 2019-07-16 04:27:05
 * @modify date 2019-07-16 04:27:05
 * @desc Asks the user for a number to create a Factorial for the given number
 */

package main

import (
	"fmt"
)

func main() {
	// Printing about the program line
	fmt.Println("This program create factorial of a given number")
	// Declaring Variables
	var factorial int
	var product int = 1

	// Getting the number from user
	for j := 1; j != 0; {
		fmt.Println("Factorial Number has to be more than 0")
		fmt.Print("Factorial: ")
		fmt.Scanf("%d", &factorial)
		if factorial > 0 {
			j = 0
		}
	}

	for i := factorial; i != 0; i-- {
		product *= i
	}
	fmt.Println(product)
}
