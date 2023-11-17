package main

import "fmt"

func printRoute(title string, slice []string) {
	fmt.Println("")
	fmt.Println("---", title, "---")
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
}

func main() {
	// creating a slice
	route := []string{"Grocerry", "Department store", "Saloon"}

	printRoute("Title1", route)
	// add to a slice at the end
	route = append(route, "Home")
	printRoute("Title2", route)

	fmt.Println()
	fmt.Println(route[0], "visited")
	fmt.Println(route[1], "visited")
	// reslicing our existing slice
	route = route[2:]
	printRoute("Remaining locations", route)

}
