/**
 * @author Muhammad Faheem Khan
 * @email faheem5948@gmail.com
 * @create date 2019-07-10 20:36:34
 * @modify date 2019-07-10 20:36:34
 * @desc Asks the user for their name and prints `Hello ${user}`
 */

package main

import (
	"fmt"
)

func main() {
	// Declaring string type variable userName, that will hold the users name
	var userName string
	// Asking the user for their name
	fmt.Print("Name: ")
	// Reading the console for name
	fmt.Scanf("%s", &userName)
	// printing Hello + userName to the screen
	fmt.Println("Hello", userName)
}
