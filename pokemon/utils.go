package pokemon

import (
	"encoding/json"
	"image"
	"math/rand"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"whosthatpokemon/object"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	AlreadyShownPokemon = mapset.NewSet[string]()
	desiredLanguages    = mapset.NewSet[string]()
)

func init() {
	desiredLanguages.Add("roomaji")
	desiredLanguages.Add("fr")
	desiredLanguages.Add("de")
	desiredLanguages.Add("it")
	desiredLanguages.Add("en")
}

func decodeImageFromWeb(imgUrl string) (*ebiten.Image, *ebiten.Image) {
	res, err := http.Get(imgUrl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	img, _, err := image.Decode(res.Body)
	if err != nil {
		panic(err)
	}

	silImg := object.ChangeImageColor(img, [3]uint8{0, 0, 0})
	actualImg := ebiten.NewImageFromImage(img)
	silhouetteImg := ebiten.NewImageFromImage(silImg)

	return actualImg, silhouetteImg
}

func unmarshalJson(jsonF []byte, target interface{}) {
	if err := json.Unmarshal(jsonF, &target); err != nil {
		panic(err)
	}
}

func prepSelectedPkm(spc *Species) *Pokemon {
	AlreadyShownPokemon.Add(spc.Name)

	randPkm := pkmsMap[spc.ID]
	randPkm.Generation = pkmsGenerationMap[spc.Generation.Name]

	for _, name := range spc.Names {
		if desiredLanguages.Contains(name.Language.Name) {
			cleanedName := strings.ReplaceAll(name.Name, "♂", "")
			cleanedName = strings.ReplaceAll(cleanedName, "♀", "")

			pkmName := PokemonName{
				Language: name.Language.Name,
				Name:     cleanedName,
			}
			if len(randPkm.OtherNames) == desiredLanguages.Cardinality() {
				break
			}
			randPkm.OtherNames = append(randPkm.OtherNames, pkmName)
		}
	}

	if randPkm.Image == nil {
		url := randPkm.ImageUrl
		randPkm.Image, randPkm.SilImage = decodeImageFromWeb(url)
	}

	return randPkm
}

func GetRandomPokemon(pkmOptions map[string][]string) *Pokemon {
	*object.SBtnClicked = false

	// Note: Games and Regions filters can be implemented.
	optsCount := len(pkmOptions)
	if optsCount > 0 {
		validGenSpcs := mapset.NewSet[string]()
		validTpySpcs := mapset.NewSet[string]()

		// For each selected generation.
		for _, genName := range pkmOptions["Generation"] {
			selGen := pkmsGenerationMap[genName]
			for _, spc := range selGen.PokemonSpecies {
				validGenSpcs.Add(spc.Name)
			}
		}

		// For each selected type.
		for _, typName := range pkmOptions["Type"] {
			selTyp := pkmsTypeMap[typName]
			for _, pkmStr := range selTyp.Pokemon {
				pkmID, _ := strconv.Atoi(filepath.Base(pkmStr.Pokemon.URL))
				// Checking that the pokemon exists in our dataset.
				if pkm, ok := pkmsMap[pkmID]; ok {
					pkmSpc := pkm.SpeciesName.Name

					validTpySpcs.Add(pkmSpc)
				}
			}
		}

		// Intersects or union the valid species depending on the user selection.
		var validSpcs mapset.Set[string]
		if validGenSpcs.Cardinality() > 0 && validTpySpcs.Cardinality() > 0 {
			validSpcs = validGenSpcs.Intersect(validTpySpcs)
		} else {
			validSpcs = validGenSpcs.Union(validTpySpcs)
		}
		// Remove species already shown.
		validSpcs = validSpcs.Difference(AlreadyShownPokemon)

		// If there is no valid species left use the already shown and try again.
		if validSpcs.Cardinality() == 0 {
			validSpcs = AlreadyShownPokemon
			AlreadyShownPokemon = mapset.NewSet[string]()
		}

		// Pick and return a random species.
		i := rand.Intn(validSpcs.Cardinality())
		spc := pkmsSpeciesMap[validSpcs.ToSlice()[i]]
		return prepSelectedPkm(spc)
	}

	for _, pkm := range pkmsMap {
		if pkm.Image == nil {
			pkm.Image, pkm.SilImage = decodeImageFromWeb(pkm.ImageUrl)
		}

		spc := pkmsSpeciesMap[pkm.SpeciesName.Name]
		return prepSelectedPkm(spc)
	}

	return new(Pokemon)
}
