// Instructions

// Write a function that returns the concatenation of two string passed in arguments.
// Expected function

// func Concat(str1 string, str2 string) string {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.Concat("Hello!", " How are you?"))

// }

// And its output :

// $ go run .
// Hello! How are you?
// $

package main

func Concat(str1 string, str2 string) string {
	return str1 + str2
}