package pokemon

import (
	"bytes"
	"compress/gzip"
	"fmt"
	_ "image/png"
	"io"
	"strings"
	"whosthatpokemon/object"
	objRes "whosthatpokemon/object/resources"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	pkmsMap           PokemonMap
	pkmsSpeciesMap    SpeciesMap
	pkmsGenerationMap GenerationMap
	pkmsTypeMap       PokemonTypeMap
	pkmsGameMap       PokemonGameMap
	Guessed           bool
	hintText          *object.PkmFont
	actualName        *object.PkmFont

	PokemonGenerations = make(map[int]string)
	PokemonTypes       = make(map[int]string)

	LastShownPokemon Pokemon
)

func init() {
	gReader, _ := gzip.NewReader(bytes.NewReader(pkmJson))
	pkmJson, _ = io.ReadAll(gReader)
	unmarshalJson(pkmJson, &pkmsMap)
	gReader, _ = gzip.NewReader(bytes.NewReader(spcJson))
	spcJson, _ = io.ReadAll(gReader)
	unmarshalJson(spcJson, &pkmsSpeciesMap)
	gReader, _ = gzip.NewReader(bytes.NewReader(genJson))
	genJson, _ = io.ReadAll(gReader)
	unmarshalJson(genJson, &pkmsGenerationMap)
	gReader.Close()
	gReader, _ = gzip.NewReader(bytes.NewReader(typeJson))
	genJson, _ = io.ReadAll(gReader)
	unmarshalJson(genJson, &pkmsTypeMap)
	gReader.Close()
	gReader, _ = gzip.NewReader(bytes.NewReader(gameJson))
	genJson, _ = io.ReadAll(gReader)
	unmarshalJson(genJson, &pkmsGameMap)
	gReader.Close()

	for genName, gen := range pkmsGenerationMap {
		PokemonGenerations[gen.ID] = genName
	}
	for typeName, typ := range pkmsTypeMap {
		PokemonTypes[typ.ID] = typeName
	}

	opt := &truetype.Options{Size: object.TxtSmallFontSize}
	hintText = object.NewFont(objRes.SenFont, opt, "", 0, 0).(*object.PkmFont)
	actualName = object.NewFont(objRes.SenFont, opt, "", 0, 0).(*object.PkmFont)
}

func (p *Pokemon) Draw(screen *ebiten.Image) {
	sW, sH := screen.Size()
	pkmW, pkmH := p.SilImage.Size()
	pkmWScale, pkmHScale := 1., 1.

	// Scaling image to a size of 475x475.
	if pkmW > 475 {
		pkmWScale = 475. / float64(pkmW)
	}
	if pkmH > 475 {
		pkmHScale = 475. / float64(pkmH)
	}

	pkmXPos := float64(sW)/2 - float64(pkmW)*pkmWScale/2
	pkmYPos := float64(sH)/2 - float64(pkmH)*pkmHScale/2

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(pkmXPos, pkmYPos)
	opts.GeoM.Scale(pkmWScale, pkmHScale)

	imgToShow := p.SilImage
	if *object.SBtnClicked || Guessed {
		imgToShow = p.Image
		LastShownPokemon = *p
	}

	pNameSplit := strings.Split(p.Name, "-")
	pName := pNameSplit[0]
	actNameX := int(pkmXPos + 5)
	actNameY := int(pkmYPos - 80)
	if !*object.SBtnClicked && !Guessed {
		actualName.Text = ""
	} else if *object.SBtnClicked {
		actualName.Text = "The actual pokemon name was: " + pName
		actualName.Text += "\nAlso known as: \n"
		for i, name := range p.OtherNames {
			sep := " "
			if i%2 == 0 {
				sep = "\n"
			}
			actualName.Text += name.Language + ": " + name.Name + sep
		}
	} else if Guessed {
		actualName.Text = "You guessed it! The pokemon name was: " + pName
		actualName.Text += "\nAlso known as: \n"
		for i, name := range p.OtherNames {
			sep := " "
			if i%2 == 0 {
				sep = "\n"
			}
			actualName.Text += name.Language + ": " + name.Name + sep
		}
	}
	actualName.X, actualName.Y = actNameX, actNameY
	actualName.Draw(screen)

	screen.DrawImage(imgToShow, opts)

	hintX := int(pkmXPos)
	hintY := int(pkmYPos + 475 + 130)
	if strings.Contains(p.Name, "-") {
		hintsItems := pNameSplit[1:]
		hint := strings.Join(hintsItems, "-")
		hint = fmt.Sprintf("Hint: %s", hint)

		hintText.Text = hint
	} else {
		hint := "Hint: "
		hint += object.GenerationsText[p.Generation.ID-1]

		hintText.Text = hint
	}
	hintText.X, hintText.Y = hintX, hintY
	hintText.Draw(screen)
}
