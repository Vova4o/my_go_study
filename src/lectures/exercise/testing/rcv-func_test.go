//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import "testing"

func newPlayer() Player {
	return Player{
		name:      "test",
		health:    100,
		maxHealth: 100,
		energy:    100,
		maxEnegry: 500,
	}
}

func TestHealth(t *testing.T) {
	player := newPlayer()
	player.addHealth(999)
	if player.health > player.maxHealth {
		t.Fatalf("Incorrect health value: %v, want %v", player.health, player.maxHealth)
	}
	player.applyDamage(player.maxHealth + 1)
	if player.health < 0 {
		t.Fatalf("Health %v, minimum: 0", player.health)
	}
	if player.health > player.maxHealth {
		t.Fatalf("Health %v, maximum: 100", player.health)
	}

}
func TestEnergy(t *testing.T) {
	player := newPlayer()
	player.addEnergy(500)
	if player.energy > 500 {
		t.Fatalf("Energy %v, maximum: %v", player.energy, player.maxEnegry)
	}
	player.consumeEnegry(600)
	if player.energy < 0 {
		t.Fatalf("Energy %v, manimum: 0", player.energy)
	}
	if player.energy > player.maxEnegry {
		t.Fatalf("Energy %v, maximum: %v", player.energy, player.maxEnegry)
	}
}
