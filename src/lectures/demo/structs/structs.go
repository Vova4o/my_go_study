package main

import "fmt"

// creating a new type Passanger that contains all inside
type Passanger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

// Creating a new type Bus
type Bus struct {
	FrontSeat Passanger
}

func main() {

	//One line create and assign variable
	casey := Passanger{"Casey", 1, false}
	fmt.Println(casey)

	// we create bill, but assign only the fields that we need
	// the rest of the feelds will be default
	var (
		bill = Passanger{Name: "Bill", TicketNumber: 2}
		ella = Passanger{Name: "Ella", TicketNumber: 3}
	)
	fmt.Println(bill, ella)

	// creating uninitiolised variable with all feelds default
	var heidi Passanger
	// assign the values to the preveously created var
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4
	fmt.Println(heidi)

	casey.Boarded = true
	bill.Boarded = true

	//we check if the condition true then print out the message
	if bill.Boarded {
		fmt.Println("Bill has boarded the bus")
	}
	// different way to print out the message
	if casey.Boarded {
		fmt.Println(casey.Name, "has boarded the bus")
	}

	heidi.Boarded = true
	bus := Bus{heidi}
	fmt.Println(bus)
	// This dot sructure shows us how nested data can be
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")
}
