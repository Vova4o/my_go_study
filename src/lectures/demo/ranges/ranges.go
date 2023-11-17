package main

import "fmt"

func main() {
	slice := []string{"Hello", "world", "!"}

	// this statmet returns i the is position in a slice
	// element retuns the value corelated to to the position
	for i, element := range slice {
		fmt.Println(i, element, ":")
		// now we eterate throug the element it self to get every position
		// that wont be possible with regular for loop
		// but we can do it with range
		for _, ch := range element {
			fmt.Printf(" %q\n", ch)
		}
	}
}
