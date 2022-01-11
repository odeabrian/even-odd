package main

import "fmt"

func main() {

	// Declare empty slice
	n := []int{}

	// Add the numbers 0 through 10 to the empty slice
	for i := 0; i < 11; i++ {
		n = append(n, i)
	}

	// Loop through each element in the slice, using the modulus operator check if there is a remainder
	// If no remainder, then the value is even, else the value is odd
	for i := range n {
		if i%2 == 0 {
			fmt.Println(n[i], "is even")
		} else {
			fmt.Println(n[i], "is odd")
		}
	}

}
