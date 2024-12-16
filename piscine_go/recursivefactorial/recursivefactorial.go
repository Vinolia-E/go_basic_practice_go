// Instructions

// Write a recursive function that returns the factorial of the int passed as parameter.

// Errors (non possible values or overflows) will return 0.

// for is forbidden for this exercise.
// Expected function

// func RecursiveFactorial(nb int) int {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	arg := 4
// 	fmt.Println(piscine.RecursiveFactorial(arg))
// }

// $ go run .
// 24
// $

package main

import "fmt"

func RecursiveFactorial(nb int) int {
	if nb == 0 {
		return 1
	}
	recurse := nb - 1
	return nb * RecursiveFactorial(recurse)
}

func main() {
	arg := 4
	fmt.Println(RecursiveFactorial(arg))
}
