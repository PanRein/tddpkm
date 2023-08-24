package service_test

import (
	"testing"
)

func TestGetElectricPokemons(t *testing.T) {
	// TODO complete this
	var numberOfElectricPokemons int
	if numberOfElectricPokemons != 9 {
		t.Error("This pokedex is broken!")
	}
}
