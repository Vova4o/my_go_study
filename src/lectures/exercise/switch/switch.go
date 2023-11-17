//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

func age() int {
	return 5
}

func main() {
	switch p := age(); {
	case p > 18:
		fmt.Println("Adult")
	case p > 13:
		fmt.Println("Teenager")
	case p > 4:
		fmt.Println("Child")
	case p > 1:
		fmt.Println("Toddler")
	default:
		fmt.Println("Newborn")
	}

	switch age := 14; {
	case age == 0:
		fmt.Println("Newborn")
	case age >= 1 && age <= 3:
		fmt.Println("Toddler")
	case age < 13:
		fmt.Println("Child")
	case age < 18:
		fmt.Println("Teenager")
	default:
		fmt.Println("Adult")
	}
}
