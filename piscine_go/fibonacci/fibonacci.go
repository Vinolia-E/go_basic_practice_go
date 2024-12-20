// Instructions

// Write a recursive function that returns the value at the position index in the fibonacci sequence.

// The first value is at index 0.

// The sequence starts this way: 0, 1, 1, 2, 3 etc...

// A negative index will return -1.

// for is forbidden for this exercise.
// Expected function

// package piscine

// func Fibonacci(index int) int {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	arg1 := 4
// 	fmt.Println(piscine.Fibonacci(arg1))
// }

// And its output :

// $ go run .
// 3
// $

package main

import "fmt"

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	} else {
		return Fibonacci(index-1) + Fibonacci(index-2)
	}
}

func main() {
	arg1 := 4
	fmt.Println(Fibonacci(arg1))
}
