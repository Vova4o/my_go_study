//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name

type Player struct {
	name              string
	health, maxHealth uint
	energy, maxEnegry uint
}

//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.

func (player *Player) addHealth(ammount uint) {
	player.health += ammount
	if player.health > player.maxHealth {
		player.health = player.maxHealth
	}
	fmt.Println(player.name, "Add", ammount, "healts ->", player.health)
}

func (player *Player) addEnergy(ammount uint) {
	player.energy += ammount
	if player.energy > player.maxEnegry {
		player.energy = player.maxEnegry
	}
	fmt.Println(player.name, "Add", ammount, "energy ->", player.energy)
}

func (player *Player) applyDamage(ammount uint) {
	if player.health-ammount > player.health {
		player.health = 0
	} else {
		player.health -= ammount
	}
	fmt.Println(player.name, "Damage", ammount, "->", player.health)
}

func (player *Player) consumeEnegry(ammount uint) {
	if player.energy-ammount > player.energy {
		player.energy = 0
	} else {
		player.energy -= ammount
	}
	fmt.Println(player.name, "consumed Enegrgy", ammount, "now left with:", player.energy)
}

//  - Print out the statistic change within each function
//  - Execute each function at least once

func main() {
	player := Player{
		name:      "PicaChu",
		health:    100,
		maxHealth: 100,
		energy:    500,
		maxEnegry: 500,
	}

	player.applyDamage(99)
	player.addHealth(20)
	player.consumeEnegry(20)
	player.addEnergy(10)

	player.applyDamage(100)
	player.consumeEnegry(550)
}
