package service

import (
	"pokedex/api"
)

type PokemonTypes struct {
	Slot int
	Type PokemonType
}

type PokemonType struct {
	Name string
}

type Pokemon struct {
	Name  string
	Types []PokemonTypes
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func getAllPokemon() []Pokemon {
	a := api.Api{}
	allPokemons := a.GetAllPokemon()
	var r []Pokemon
	for _, p := range allPokemons {

		var types []PokemonTypes
		for _, t := range p.Types {
			types = append(types, PokemonTypes{
				Slot: t.Slot,
				Type: PokemonType{
					Name: t.Type.Name,
				},
			})
		}

		r = append(r, Pokemon{
			Name:  p.Name,
			Types: types,
		})
	}
	return r
}
