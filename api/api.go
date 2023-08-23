package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type PokemonUrl struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type AllPokemonApiResponse struct {
	Count       int          `json:"count"`
	Next        string       `json:"next"`
	Previous    string       `json:"previous"`
	PokemonUrls []PokemonUrl `json:"results"`
}

type PokemonTypes struct {
	Slot int         `json:"slot"`
	Type PokemonType `json:"type"`
}

type PokemonType struct {
	Name string `json:"name"`
}

type Pokemon struct {
	Name  string         `json:"name"`
	Types []PokemonTypes `json:"types"`
}

type Api struct {
}

func (a Api) GetAllPokemon() []Pokemon {
	gen1PokemonUrl := "https://pokeapi.co/api/v2/pokemon?limit=151"
	res, _ := http.Get(gen1PokemonUrl)
	body, _ := io.ReadAll(res.Body)
	var r AllPokemonApiResponse
	json.Unmarshal(body, &r)
	var pokemons []Pokemon
	for _, p := range r.PokemonUrls {
		res, _ := http.Get(p.Url)
		body, _ := io.ReadAll(res.Body)
		var pkm Pokemon
		json.Unmarshal(body, &pkm)
		pokemons = append(pokemons, pkm)
	}
	return pokemons
}
