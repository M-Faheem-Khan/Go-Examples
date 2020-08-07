// Pseudorandom Number Generation: Linear Congruential Method
// x = (ax + c) % m
// https://en.wikipedia.org/wiki/Linear_congruential_generator

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	fmt.Println("=== Linear Congruential Method ===")
	a := readInteger("a: ")    // multiplier
	seed := readInteger("x: ") // seed integer X(0) = 3
	m := readInteger("m: ")    // mod to limit the integer
	c := rand.Intn(m - 1)      // randomly generated between 0 <= x < m
	x := seed                  // x is used to check when the numbers starts to repeat
	counter := 0               // after how many unique integers will the numbers start to repeat

	for {
		x = (a*x + c) % m
		fmt.Printf("%d\n", x)

		// If the number start to repeat
		if x == seed {
			break
		}
		counter = counter + 1 // incrementing counter
	}

	// printing number of unique integers before sequence repitition
	fmt.Printf("There are %d Unique Numbers before repeating sequence", counter)

}

// Read integer from cli
func readInteger(prompt string) int {
	reader := bufio.NewReader(os.Stdin)            // read from stdin
	fmt.Printf("%s", prompt)                       // printin g
	userInput, _ := reader.ReadString('\n')        // read until '\n'
	userInput = userInput[:len(userInput)-2]       // removing '\r\n'
	userInputInteger, _ := strconv.Atoi(userInput) // casting from string to integer
	return userInputInteger                        // returning integer
}
