package main

import "fmt"

func main() {
	// creating an empty map
	// after map the firs is the KEY and the second is VALUE
	shopingList := make(map[string]int)

	//adding some stuff to a map
	//adding it by chousing map and giving a key and a value
	//-----------KEY-----Value
	shopingList["eggs"] = 11
	shopingList["milk"] = 1
	shopingList["bread"] += 1
	// shopingList["cerial"] = 2

	//adding more to an existing key
	shopingList["eggs"] += 1
	fmt.Println(shopingList)

	//deleting from shopping list
	delete(shopingList, "milk")

	fmt.Println("Milk deleted, new list:", shopingList)

	//getting acces to a value of eggs by et key
	fmt.Println("nned", shopingList["eggs"], "eggs")

	//looking for cerial
	cerial, found := shopingList["cerial"]
	fmt.Println("Need cerial?")
	if !found {
		fmt.Println("nope")
	} else {
		fmt.Println("yep", cerial, "boxes")
	}
}
