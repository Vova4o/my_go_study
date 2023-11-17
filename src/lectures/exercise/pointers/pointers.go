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

const (
	Active   = true
	inactive = false
)

type SecurityTag bool

// * Create a structure to store items and their security tag state
type Item struct {
	name string
	tag  SecurityTag
}

//   - Security tags have two states: active (true) and inactive (false)
//
// * Create functions to activate and deactivate security tags using pointers
func activate(tag *SecurityTag) {
	*tag = Active
}

func deactivate(tag *SecurityTag) {
	*tag = inactive
}

// * Create a checkout() function which can deactivate all tags in a slice
func checkout(items []Item) {
	fmt.Println("checking out...")
	for i := 0; i < len(items); i++ {
		deactivate(&items[i].tag)
	}
}

func main() {

	//  - Create at least 4 items, all with active security tags
	jeans := Item{"Jeans", Active}
	short := Item{"Short", Active}
	pants := Item{"Pants", Active}
	jacket := Item{"Jacket", Active}

	fmt.Println(jeans, short, pants, jacket)

	//  - Store them in a slice or array
	listOfItems := []Item{jeans, short, pants, jacket}
	fmt.Println(listOfItems)

	deactivate(&listOfItems[0].tag)
	// tagsState(false, ilstOfItems)
	fmt.Println(listOfItems)

	checkout(listOfItems)
	fmt.Println("checked out", listOfItems)

}
