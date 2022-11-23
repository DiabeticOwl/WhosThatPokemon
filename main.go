package main

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"whosthatpokemon/object"
	"whosthatpokemon/pokemon"

	"code.rocketnine.space/tslocum/messeji"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.design/x/clipboard"
)

const (
	gameWidth  int = 1000
	gameHeight int = 800
)

type Game struct {
	w, h    *int
	tick    *int
	pokemon *pokemon.Pokemon

	objects []object.Object

	inpFld          *messeji.InputField
	inputtedPokemon *string
	score           *string
}

func (g *Game) Update() error {
	*g.tick++

	for _, obj := range g.objects {
		obj.Update()
	}

	g.inpFld.SetSingleLine(true)
	if err := g.inpFld.Update(); err != nil {
		panic(err)
	}
	if *object.SBtnClicked {
		g.inpFld.SetText("")
	}

	if ebiten.IsKeyPressed(ebiten.KeyControl) && ebiten.IsKeyPressed(ebiten.KeyV) {
		g.inpFld.SetText(string(clipboard.Read(clipboard.FmtText)))
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		cleanInput := strings.ToLower(*g.inputtedPokemon)

		pkmNames := mapset.NewSet[string]()
		cleanPkmNames := mapset.NewSet[string]()
		for _, name := range g.pokemon.OtherNames {
			pkmNames.Add(strings.ToLower(name.Name))

			cleanedName := strings.ToLower(strings.Split(name.Name, "-")[0])
			cleanPkmNames.Add(cleanedName)
		}

		if pkmNames.Contains(cleanInput) || cleanPkmNames.Contains(cleanInput) {
			pokemon.Guessed = true

			actualSc, _ := strconv.Atoi(*g.score)
			actualSc += 1

			*object.GameScore = fmt.Sprintf("%03d", actualSc)

			g.inpFld.SetText("")
		} else {
			pokemon.Guessed = false
		}
	}
	if *g.tick%240 == 0 && (pokemon.Guessed || *object.SBtnClicked) {
		pokemon.Guessed = false
		*object.SBtnClicked = false

		pkmOptions := make(map[string][]string)
		if len(object.SelectedGenerations) > 0 {
			var selectedGene []string
			for gen := range object.SelectedGenerations {
				genNameId := pokemon.PokemonGenerations[gen]
				selectedGene = append(selectedGene, genNameId)
			}

			pkmOptions["Generation"] = selectedGene
		}
		if len(object.SelectedTypes) > 0 {
			var selectedTyp []string
			for typ := range object.SelectedTypes {
				typNameId := pokemon.PokemonTypes[typ]
				selectedTyp = append(selectedTyp, typNameId)
			}

			pkmOptions["Type"] = selectedTyp
		}

		g.pokemon = pokemon.GetRandomPokemon(pkmOptions)
	}

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	object.GameWidth, object.GameHeight = outsideWidth, outsideHeight

	if outsideWidth == *g.w && outsideHeight == *g.h {
		return outsideWidth, outsideHeight
	}

	// 475 being the standard width and height of the pokemon images.
	inpX := outsideWidth/2 - 475/2
	inpY := outsideHeight/2 + 475/2

	g.inpFld.SetRect(image.Rect(inpX, inpY, inpX+475, inpY+50))

	*g.w, *g.h = outsideWidth, outsideHeight

	return outsideWidth, outsideHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{235, 235, 235, 1})

	g.pokemon.Draw(screen)

	// Note: Since these objects doesn't overlap between themselves the slice
	// of objects can become a map in order to track which object is being
	// targeted and perform accordingly.
	for _, obj := range g.objects {
		obj.Draw(screen)
	}

	g.inpFld.Draw(screen)
}

func (g *Game) setUpInputField() {
	g.inpFld.SetSelectedFunc(func() (accept bool) {
		*g.inputtedPokemon = g.inpFld.Text()

		return true
	})
}

func main() {
	rand.Seed(time.Now().UnixMilli())

	game := &Game{
		w:               new(int),
		h:               new(int),
		tick:            new(int),
		pokemon:         pokemon.GetRandomPokemon(map[string][]string{}),
		inputtedPokemon: new(string),
		score:           new(string),
		objects: []object.Object{
			object.NewSurrenderBtn(),
			object.NewScore(),
		},
	}

	*game.score = fmt.Sprintf("%03d", 0)
	object.GameScore = game.score

	game.objects = append(game.objects, object.InitializeGenerationBtns()...)
	game.objects = append(game.objects, object.InitializeTypeBtns()...)

	// After the fonts has been instantiated.
	game.inpFld = object.NewInputField()
	game.setUpInputField()

	ebiten.SetWindowSize(gameWidth, gameHeight)
	ebiten.SetWindowTitle("Who's that pokemon?")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
