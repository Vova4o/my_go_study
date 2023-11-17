package main

import "fmt"

// single parking space
// with defenition ocupied or not
type Space struct {
	occupied bool
}

// Contain all of the parking spaces
type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {

	lot := ParkingLot{spaces: make([]Space, 4)}
	fmt.Println("Initial state of lot:", lot)

	lot.occupySpace(1)
	occupySpace(&lot, 2)

	fmt.Println("After two cars came in:", lot)

	lot.vacateSpace(1)

	fmt.Println("After one car came left:", lot)

}
