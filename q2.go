//A function that calls itself is called a recursive function in programming.

package main

import "fmt"

func generateOutput(n int) {
	// Base case: if n is less than 2, exit the function.
	if n < 2 {
		return
	}

	// Recursive:
	// 1. Call the function recursively with n/2.
	generateOutput(n / 2)
	// 2. Print n.
	fmt.Println(n)
}

func main() {
	generateOutput(9)
}
