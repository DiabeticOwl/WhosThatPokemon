package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

// Information from the corpora of objects in the API.
type PokemonInfo struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var (
	pkmInfos = make(map[string]PokemonInfo)

	wg sync.WaitGroup
)

func ExtractPokemonInfo() {
	// Limit is used for the entire corpora of data within the API.
	pkmInfoUrl := "https://pokeapi.co/api/v2/pokemon/?limit=10000"
	pkmSpecInfoUrl := "https://pokeapi.co/api/v2/pokemon-species/?limit=10000"
	pkmGeneInfoUrl := "https://pokeapi.co/api/v2/generation/?limit=100"
	pkmTypeInfoUrl := "https://pokeapi.co/api/v2/type/?limit=100"
	pkmGameInfoUrl := "https://pokeapi.co/api/v2/version-group/?limit=100"

	pkmInfos = map[string]PokemonInfo{
		"pokemon":         getInfoJson(pkmInfoUrl),
		"pokemon-species": getInfoJson(pkmSpecInfoUrl),
		"generation":      getInfoJson(pkmGeneInfoUrl),
		"pokemon-type":    getInfoJson(pkmTypeInfoUrl),
		"pokemon-game":    getInfoJson(pkmGameInfoUrl),
	}

	wg.Add(len(pkmInfos))
	go extractPokemon()
	go extractSpecies()
	go extractGeneration()
	go extractPokemonType()
	go extractPokemonGame()
	wg.Wait()
}

func extractPokemon() {
	pkms := make(map[int]Pokemon)
	var i int
	for _, result := range pkmInfos["pokemon"].Results {
		// Bulk timeout. One second each 100 extractions.
		if i%100 == 0 {
			time.Sleep(1 * time.Second)
		}
		i++
		var pkm Pokemon

		jsonResp, err := http.Get(result.URL)
		if err != nil {
			panic(err)
		}
		json.NewDecoder(jsonResp.Body).Decode(&pkm)

		artUrl := pkm.Sprites.Other.OfficialArtwork.FrontDefault
		if artUrl == "" {
			artUrl = pkm.Sprites.Other.Home.FrontDefault
		}
		// If there is no artwork available.
		if artUrl == "" {
			continue
		}

		pkm.ImageUrl = artUrl

		pkms[pkm.ID] = pkm

		fmt.Printf("Pokemon info of \"%s\" retrieved.\n", pkm.Name)
	}

	jsonF, err := os.Create("../json/pokemonDetails.json")
	if err != nil {
		panic(err)
	}
	defer jsonF.Close()

	jsonSlice, err := json.Marshal(pkms)
	if err != nil {
		panic(err)
	}
	jsonF.Write(jsonSlice)

	wg.Done()
}

func extractSpecies() {
	species := make(map[string]Species)
	var i int
	for _, result := range pkmInfos["pokemon-species"].Results {
		// Bulk timeout. One second each 100 extractions.
		if i%100 == 0 {
			time.Sleep(1 * time.Second)
		}

		i++
		var spc Species
		jsonResp, err := http.Get(result.URL)
		if err != nil {
			panic(err)
		}
		json.NewDecoder(jsonResp.Body).Decode(&spc)

		// Note: If name is not enough, use the entire species URL.
		species[spc.Name] = spc

		fmt.Printf("Species info of \"%s\" retrieved.\n", spc.Name)
	}

	jsonF, err := os.Create("../json/pokemonSpecies.json")
	if err != nil {
		panic(err)
	}
	defer jsonF.Close()

	jsonSlice, err := json.Marshal(species)
	if err != nil {
		panic(err)
	}
	jsonF.Write(jsonSlice)

	wg.Done()
}

func extractGeneration() {
	genes := make(map[string]Generation)
	for _, result := range pkmInfos["generation"].Results {
		var gene Generation

		jsonResp, err := http.Get(result.URL)
		if err != nil {
			panic(err)
		}
		json.NewDecoder(jsonResp.Body).Decode(&gene)

		// Note: If name is not enough, use the entire generation URL.
		genes[gene.Name] = gene

		fmt.Printf("\"%s\" retrieved.\n", gene.Name)
	}

	jsonF, err := os.Create("../json/generationDetails.json")
	if err != nil {
		panic(err)
	}
	defer jsonF.Close()

	jsonSlice, err := json.Marshal(genes)
	if err != nil {
		panic(err)
	}
	jsonF.Write(jsonSlice)

	wg.Done()
}

func extractPokemonType() {
	pkmTypes := make(map[string]PokemonType)
	for _, result := range pkmInfos["pokemon-type"].Results {
		// Ignoring unofficial pokemon types.
		if result.Name == "unknown" || result.Name == "shadow" {
			continue
		}

		var pkmType PokemonType

		jsonResp, err := http.Get(result.URL)
		if err != nil {
			panic(err)
		}
		json.NewDecoder(jsonResp.Body).Decode(&pkmType)

		pkmTypes[pkmType.Name] = pkmType

		fmt.Printf("\"%s\" type retrieved.\n", pkmType.Name)
	}

	jsonF, err := os.Create("../json/typeDetails.json")
	if err != nil {
		panic(err)
	}
	defer jsonF.Close()

	jsonSlice, err := json.Marshal(pkmTypes)
	if err != nil {
		panic(err)
	}
	jsonF.Write(jsonSlice)

	wg.Done()
}

func extractPokemonGame() {
	pkmGames := make(map[string]PokemonGame)
	for _, result := range pkmInfos["pokemon-type"].Results {
		var pkmGame PokemonGame

		jsonResp, err := http.Get(result.URL)
		if err != nil {
			panic(err)
		}
		json.NewDecoder(jsonResp.Body).Decode(&pkmGame)

		pkmGames[pkmGame.Name] = pkmGame

		fmt.Printf("\"%s\" type retrieved.\n", pkmGame.Name)
	}

	jsonF, err := os.Create("../json/gameDetails.json")
	if err != nil {
		panic(err)
	}
	defer jsonF.Close()

	jsonSlice, err := json.Marshal(pkmGames)
	if err != nil {
		panic(err)
	}
	jsonF.Write(jsonSlice)

	wg.Done()
}
