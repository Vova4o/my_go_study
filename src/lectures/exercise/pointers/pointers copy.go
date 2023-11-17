//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:

//* Perform the following:

//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

// * Create a structure to store items and their security tag state
//   - Security tags have two states: active (true) and inactive (false)
type Tags struct {
	item  string
	state bool
}

// * Create functions to activate and deactivate security tags using pointers
func tagsState(status bool, itemName []Tags, itemNum int) {
	itemName[itemNum].state = status
}

// * Create a checkout() function which can deactivate all tags in a slice
func checkout(items []Tags) {
	for key, _ := range items {
		item := items[key]
		item.state = false
		fmt.Println(item)
	}
}

func main() {

	//  - Create at least 4 items, all with active security tags
	var (
		jeans  = Tags{item: "Jeans", state: true}
		short  = Tags{item: "Short", state: true}
		pants  = Tags{item: "Pants", state: true}
		jacket = Tags{item: "Jacket", state: true}
	)
	fmt.Println(jeans, short, pants, jacket)

	//  - Store them in a slice or array
	listOfItems := []Tags{jeans, short, pants, jacket}
	fmt.Println(listOfItems)

	checkout(listOfItems)
	// tagsState(false, ilstOfItems)
	fmt.Println(listOfItems)

}
