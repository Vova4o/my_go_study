//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
// func greet(name str) str {
// 	fmt.Println("Good time of the day", name, "!")
// }

func greet(name string) {
	fmt.Println("Good time of the day", name)
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func message() string {
	return "Hello this is a message from a function!"
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func summOfThreeNum(a, b, c int) int {
	return a + b + c
}

// * Write a function that returns any number
func anyNumber() int {
	return 5
}

// * Write a function that returns any two numbers
func anyTwoNunmbers() (int, int) {
	return 2, 2
}

//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

func main() {
	greet("Vladimir")
	fmt.Println(message())
	//answer = 6
	answer1 := summOfThreeNum(1, 2, 3)
	// print = 6
	fmt.Println(answer1+2)

	fmt.Println(anyNumber())
	fmt.Println(anyTwoNunmbers())
	a, b := anyTwoNunmbers()
	// znachenie , err := anyTwoNunmbers()
	summOfNumbers := summOfThreeNum(a, b, anyNumber())
	fmt.Println(summOfNumbers)
}
