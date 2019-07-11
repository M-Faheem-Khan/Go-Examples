/**
 * @author Muhammad Faheem Khan
 * @email faheem5948@gmail.com
 * @create date 2019-07-10 20:48:33
 * @modify date 2019-07-10 20:48:33
 * @desc Ask the user for two numbers and perform the four basic math opperations (+, -, *, /) on them
 */

package main

import (
	"fmt"
)

func main() {
	// Printing the about this program line
	fmt.Println("This is a simple program that performs the 4 basic math opperations of any two given numbers")
	// variables
	var a int
	var b int
	// getting inputs
	fmt.Println("Add a single space character in between the two numbers. Example: 10 20 ")
	fmt.Print("Numbers: ")
	fmt.Scanf("%d %d", &a, &b)
	// defining new variables that hold the output of the four basic math opperations
	var add = a + b
	var subtract = a - b
	var divide = a / b
	var multiply = a * b
	// printing the output to the screen
	fmt.Println(a, "+", b, "=", add)
	fmt.Println(a, "-", b, "=", subtract)
	fmt.Println(a, "*", b, "=", multiply)
	fmt.Println(a, "/", b, "=", divide)
}
