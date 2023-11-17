package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {

	switch p := price(); {
	case p < 2:
		fmt.Println("Cheap item.")
	case p < 10:
		fmt.Println("Moderate priced")
	default:
		fmt.Println("Expancive item")
	}

	ticket := Economy

	switch ticket {
	case Economy:
		fmt.Println("Economy seating")
	case Business:
		fmt.Println("Business seating")
	case FirstClass:
		fmt.Println("FirstClass seating")
	default:
		fmt.Println("I guess you be standing")
	}
}
