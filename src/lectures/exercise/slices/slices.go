//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printContent(prep []Part) {
	fmt.Println("Accembly line contain:")
	for i := 0; i < len(prep); i++ {
		element := prep[i]
		fmt.Println(element)
	}
}

func main() {

	assemblyLine := []Part{"Belt", "Rubber", "Motor"}

	printContent(assemblyLine)

	assemblyLine = append(assemblyLine, "Cover", "Roller")

	fmt.Println("Add a few more parts")

	printContent(assemblyLine)

	assemblyLine = assemblyLine[3:]

	fmt.Println("Removing first three items.")

	printContent(assemblyLine)

}
