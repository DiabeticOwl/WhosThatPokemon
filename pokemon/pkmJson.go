package pokemon

import "github.com/hajimehoshi/ebiten/v2"

type PokemonName struct {
	Language string
	Name     string
}
type PokemonMap map[int]*Pokemon
type Pokemon struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	SpeciesName struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	ImageUrl string `json:"ImageUrl"`

	OtherNames []PokemonName
	Generation *Generation

	Image    *ebiten.Image
	SilImage *ebiten.Image
}

type SpeciesMap map[string]*Species
type Species struct {
	Generation struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"generation"`
	ID          int    `json:"id"`
	IsBaby      bool   `json:"is_baby"`
	IsLegendary bool   `json:"is_legendary"`
	IsMythical  bool   `json:"is_mythical"`
	Name        string `json:"name"`
	Names       []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Varieties []struct {
		IsDefault bool `json:"is_default"`
		Pokemon   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"varieties"`
}

type GenerationMap map[string]*Generation
type Generation struct {
	ID         int `json:"id"`
	MainRegion struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"main_region"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonSpecies []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon_species"`
}

type PokemonTypeMap map[string]*PokemonType
type PokemonType struct {
	GameIndices []struct {
		GameIndex  int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"game_indices"`
	Generation struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"generation"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Pokemon []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		Slot int `json:"slot"`
	} `json:"pokemon"`
}

type PokemonGameMap map[string]*PokemonGame
type PokemonGame struct {
	Generation struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"generation"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Order     int    `json:"order"`
	Pokedexes []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokedexes"`
	Regions []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"regions"`
	Versions []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"versions"`
}
