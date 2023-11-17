//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
)

var numberOfRoll int
var numberOfDice int
var numberOfSides int
var summOfScore int

func main() {
	// rand.Seed(100)
	rand.Seed(time.now().Unix)
	numberOfRoll = 2
	numberOfDice = 2
	numberOfSides = 6
	// we perfom number of roll that we are going to do and pass it to each dice
	for i := 1; i <= numberOfRoll; i++ {
		fmt.Println("Dice", i)
		// No we take each dice and roll it
		for i := 1; i <= numberOfDice; i++ {
			// now we need to figure oute number for side
			// but we can do it with random function
			// cause that the end of equation
			fmt.Println("Roll", i)
			summOfScore = summOfScore + rand.Intn(numberOfSides)
			// fmt.Println("Random for umber of Sides", rand.Intn(numberOfSides))
			fmt.Println(summOfScore)
		}
	}
	//* Print additional information in these cirsumstances:
	//  - "Snake eyes": when the total roll is 2, and total dice is 2
	//  - "Lucky 7": when the total roll is 7
	//  - "Even": when the total roll is even
	//  - "Odd": when the total roll is odd

	// var even bool = summOfScore%2 == 0

	switch summOfScore {
	case 2:
		fmt.Println("Snake eyes", summOfScore)
	case 7:
		//  - "Lucky 7": when the total roll is 7
		fmt.Println("Lucky 7", summOfScore)
	case summOfScore % 2:
		fmt.Println("Even", summOfScore)
	default:
		fmt.Println("Odd", summOfScore)
	}
	fmt.Println(summOfScore)
}
