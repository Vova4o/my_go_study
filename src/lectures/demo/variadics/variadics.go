package main

import "fmt"

// creating a function that takes unlimited ammount of params
func sum(nums ...int) int {
	sum := 0
	// itterate throught the params
	for _, n := range nums {
		// summ them
		sum += n
	}
	// return the summm
	return sum
}

func main() {

	// creating a slice
	a := []int{1, 2, 3}

	// creating another slice
	b := []int{4, 5, 6}

	// append one slice to another
	//however we hace to use ... at the end
	//cause the wont append otherwise
	all := append(a, b...)

	answer := sum(all...)

	fmt.Println(answer)

	answer = sum(1, 2, 3, 4, 5, 6)
	fmt.Println(answer)

}
