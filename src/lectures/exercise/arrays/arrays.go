//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:

//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

// - Products must include the price and the name
type Product struct {
	name  string
	price int
}

func printLastItem(prep [4]Product) {
	for i := (len(prep) - 1); i >= 0; i-- {
		item := prep[i]
		if item.name != "" {
			fmt.Println(item.name, "last item on the list")
		}
	}
}

func printTotanItemsInBasket(prep [4]Product) {
	k := 0
	for i := 0; i < len(prep); i++ {
		item := prep[i]
		if item.name != "" {
			k++
		}
	}
	fmt.Println("Total items in shoping cart", k)
}

func totalPrice(prep [4]Product) {
	k := 0
	for i := 0; i < len(prep); i++ {
		k += prep[i].price
	}
	fmt.Println("Total cost of the items in the cart", k)
}

func main() {

	// Using an array, create a shopping list with enough room
	//  for 4 products

	//How to add multiple items to an array...
	// shoppingList := [4]Product{
	// 	{"Banana", 10},
	// 	{"Milk", 20},
	// 	{"Beef", 30},
	// }

	var shoppingList = [4]Product{}

	// Insert 3 products into the array
	shoppingList[0].name = "Apple"
	shoppingList[0].price = 10
	shoppingList[1].name = "Banana"
	shoppingList[1].price = 20
	shoppingList[2].name = "Beef"
	shoppingList[2].price = 30

	printLastItem(shoppingList)
	printTotanItemsInBasket(shoppingList)
	totalPrice(shoppingList)

	fmt.Println("Add one more item to the list")

	shoppingList[3].name = "Milk"
	shoppingList[3].price = 33

	printLastItem(shoppingList)
	printTotanItemsInBasket(shoppingList)
	totalPrice(shoppingList)

}
