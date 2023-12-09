package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(prop [4]Room) {
	for i := 0; i < len(prop); i++ {
		room := prop[i]
		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is durty")
		}
	}
}

func main() {

	rooms := [...]Room{
		{name: "Office"},
		{name: "Warehouce"},
		{name: "Reception"},
		{name: "Ops"},
	}

	checkCleanliness(rooms)

	fmt.Println("Performing some cleaning")

	rooms[2].cleaned = true
	rooms[3].cleaned = true

	checkCleanliness(rooms)
}
