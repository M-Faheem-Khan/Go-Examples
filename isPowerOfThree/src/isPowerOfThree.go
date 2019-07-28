/**
 * @author Muhammad Faheem Khan
 * @email faheem5948@gmail.com
 * @create date 2019-07-28 02:11:11
 * @modify date 2019-07-28 02:11:11
 * @desc Validates if a given number is a Power of three
 */

package main

import "fmt"

func main() {
	//  Printing About program string
	fmt.Println("Validates if a given number is Power of Three or Not")
	// Variable
	var number int
	// Getting the number to validate
	fmt.Print("Number: ")
	fmt.Scanln(&number)
	// checking if the number is a valid power of three
	// while loop
	for i := 1; i != 0; {
		// if the mod 3 of the number is not 0 it is not a odd / easily divisible by 3 (whole number)
		if number%3 != 0 {
			i = 0 // break the loop
		} else {
			number = number / 3 // divide the number
		}
		// breaking the loop if the number is less than 3
		if number <= 2 {
			i = 0
		}
	}
	fmt.Println(number == 1) // print if the number is valid or not
}

// if n % 3 != 0: return False
// else: n=n/3
// return n == 1
