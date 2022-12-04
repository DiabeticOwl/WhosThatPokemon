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
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.design/x/clipboard"
)

const (
	gameWidth  int = 1000
	gameHeight int = 800
)

// Game is the main instance of ebiten.Game to be executed.
type Game struct {
	// Describes the width and height of the game.
	w int

	// Describes the height of the game.
	h int

	// Describes the internal time in which the game is currently at.
	// For more information please read the Update method: https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#Game
	tick *int

	// Points to the currently shown Pokemon value.
	pokemon *pokemon.Pokemon

	// Slice of object.Object type that stores first-party assets to draw and
	// update in the game.
	objects []object.Object

	// Points to the instance of messeji.InputField that is available to the
	// user.
	inpFld *messeji.InputField

	// Stores the user's inputted text through the messeji.InputField.
	inputtedPokemon string
}

// Update proceeds the game state by executing the "Update" method in each
// object within the game itself.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	*g.tick++

	for _, obj := range g.objects {
		obj.Update()
	}

	// Enables the user to type inside the InputField.
	g.inpFld.SetSingleLine(true)
	if err := g.inpFld.Update(); err != nil {
		panic(err)
	}
	if *object.SBtnClicked {
		g.inpFld.SetText("")
	}

	if ebiten.IsKeyPressed(ebiten.KeyControl) &&
		ebiten.IsKeyPressed(ebiten.KeyV) {
		g.inpFld.SetText(string(clipboard.Read(clipboard.FmtText)))
	}
	if ebiten.IsKeyPressed(ebiten.KeyEnter) ||
		ebiten.IsKeyPressed(ebiten.KeyNumpadEnter) {
		// Inputted text will be parsed accordingly in order to be properly used.
		cleanInput := strings.ToLower(g.inputtedPokemon)

		// Usage of mapsets in order to harness their Contains method.
		// https://pkg.go.dev/github.com/deckarep/golang-set/v2#Set
		pkmNames := mapset.NewSet[string]()
		cleanPkmNames := mapset.NewSet[string]()
		// Recollection of pokemon name's variants.
		for _, name := range g.pokemon.OtherNames {
			pkmNames.Add(strings.ToLower(name.Name))

			cleanedName := strings.ToLower(strings.Split(name.Name, "-")[0])
			cleanPkmNames.Add(cleanedName)
		}

		if pkmNames.Contains(cleanInput) || cleanPkmNames.Contains(cleanInput) {
			pokemon.Guessed = true

			actualSc, _ := strconv.Atoi(object.GameScore)
			actualSc += 1

			object.GameScore = fmt.Sprintf("%03d", actualSc)

			g.inpFld.SetText("")
		} else {
			pokemon.Guessed = false
		}
	}
	// Each 4 seconds, if the pokemon is guessed or the user surrenders by
	// clicking the corresponding button, another pokemon given the selected
	// filters will be chosen.
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

// Layout takes the outside size (e.g., the window size) and returns the
// (logical) screen size. It runs every time the outside size changes, updating
// the InputField accordingly.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	object.GameWidth, object.GameHeight = outsideWidth, outsideHeight

	if outsideWidth == g.w && outsideHeight == g.h {
		return outsideWidth, outsideHeight
	}

	// 475 being the standard width and height of the pokemon images.
	inpX := outsideWidth/2 - 475/2
	inpY := outsideHeight/2 + 475/2

	g.inpFld.SetRect(image.Rect(inpX, inpY, inpX+475, inpY+50))

	g.w, g.h = outsideWidth, outsideHeight

	return outsideWidth, outsideHeight
}

// Draw executes the "Draw" method in each object within the game.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
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

	footerX := g.w/2 - 80
	footerY := g.h/2 + g.pokemon.Image.Bounds().Dy()/2 + 140
	ebitenutil.DebugPrintAt(screen, "Made by DiabeticOwl", footerX, footerY)
}

func (g *Game) setUpInputField() {
	g.inpFld.SetSelectedFunc(func() (accept bool) {
		g.inputtedPokemon = g.inpFld.Text()

		return true
	})
}

func main() {
	rand.Seed(time.Now().UnixMilli())

	game := &Game{
		tick:    new(int),
		pokemon: pokemon.GetRandomPokemon(map[string][]string{}),
		objects: []object.Object{
			object.NewSurrenderBtn(),
			object.NewScore(),
		},
	}

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
