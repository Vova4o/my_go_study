//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Coordinates struct {
	X int
	Y int
}

type Rectangle struct {
	BottomLeft Coordinates
	TopRight   Coordinates
}

func perimeterCalc(prop Rectangle) int {
	return (prop.TopRight.X-prop.BottomLeft.X)*2 + (prop.TopRight.Y-prop.BottomLeft.Y)*2
}

func areaCalc(prop Rectangle) int {
	return ((prop.TopRight.X - prop.BottomLeft.X) * (prop.TopRight.Y - prop.BottomLeft.Y))
}

func main() {
	pointOne := Coordinates{X: 3, Y: 3}
	fmt.Println(pointOne)
	pointTwo := Coordinates{X: 9, Y: 9}
	fmt.Println(pointTwo)
	rectangle := Rectangle{pointOne, pointTwo}
	fmt.Println(rectangle)

	fmt.Println("Perimeter", perimeterCalc(rectangle))
	fmt.Println("Area", areaCalc(rectangle))
}
