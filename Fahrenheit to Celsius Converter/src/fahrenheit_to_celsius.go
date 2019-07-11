/**
 * @author Muhammad Faheem Khan
 * @email faheem5948@gmail.com
 * @create date 2019-07-10 21:09:50
 * @modify date 2019-07-10 21:09:50
 * @desc Asks the user for the temprature in Fahrenheit and prints to the screen the celsius converted temprature
 */

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Fahrenheit to Celsius Converter")
	var temp float32
	fmt.Print("Fahrenheit: ")
	fmt.Scanf("%f", &temp)
	var celsius = float32(temp-32) * float32(5) / float32(9)
	fmt.Printf("Celsius: %f", celsius)
}
