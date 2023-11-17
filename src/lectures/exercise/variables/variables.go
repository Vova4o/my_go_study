//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	// * Store your favorite color in a variable using the `var` keyword
	var color = "green"
	fmt.Println("My favorit color"+" is", color)

	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	year, age := 1980, 43
	fmt.Println("My age", age, ",I was born at", year)

	//* Store your first & last initials in two variables using block assignment
	var (
		firstInitial  = "V"
		secondInitial = "G"
	)
	fmt.Println("Initials=", firstInitial, secondInitial)

	//* Declare (but don't assign!) a variable for your age (in days),
	var myAgeInDays int
	myAgeInDays = age * 365
	fmt.Println("My age in days", myAgeInDays)
	//  then assign it on the next line by multiplying 365 with the age
	// 	variable created earlier
	//
}
