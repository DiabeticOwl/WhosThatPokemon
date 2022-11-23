package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getInfoJson(url string) PokemonInfo {
	fmt.Printf("Getting the urls of %s.\n", url)

	var pkmUrls PokemonInfo
	jsonResp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	json.NewDecoder(jsonResp.Body).Decode(&pkmUrls)

	return pkmUrls
}
