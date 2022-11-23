package main

import (
	"fmt"
	"os"
)

func main() {
	ExtractPokemonInfo()

	genFile, err := os.Create("generate.go")
	if err != nil {
		panic(err)
	}
	defer genFile.Close()

	mainCmd := "//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice "

	// Json Resources.
	pkmJson := mainCmd + "-compress "
	pkmJson += "-input ../json/pokemonDetails.json -package pokemon "
	pkmJson += "-output ../../pokemon/pokemonDetails.go -var pkmJson"
	fmt.Fprintln(genFile, pkmJson)

	spcJson := mainCmd + "-compress "
	spcJson += "-input ../json/pokemonSpecies.json -package pokemon "
	spcJson += "-output ../../pokemon/pokemonSpecies.go -var spcJson"
	fmt.Fprintln(genFile, spcJson)

	genJson := mainCmd + "-compress "
	genJson += "-input ../json/generationDetails.json -package pokemon "
	genJson += "-output ../../pokemon/generationDetails.go -var genJson"
	fmt.Fprintln(genFile, genJson)

	typeJson := mainCmd + "-compress "
	typeJson += "-input ../json/typeDetails.json -package pokemon "
	typeJson += "-output ../../pokemon/typeDetails.go -var typeJson"
	fmt.Fprintln(genFile, typeJson)

	gameJson := mainCmd + "-compress "
	gameJson += "-input ../json/gameDetails.json -package pokemon "
	gameJson += "-output ../../pokemon/gameDetails.go -var gameJson"
	fmt.Fprintln(genFile, gameJson)

	// Objects Resources.
	senFont := mainCmd + "-input ../fonts/Sen/Sen-Regular.ttf -package font "
	senFont += "-output ../../object/resources/senRegularFont.go "
	senFont += "-var SenFont"
	fmt.Fprintln(genFile, senFont)

	boldSenFont := mainCmd + "-input ../fonts/Sen/Sen-ExtraBold.ttf -package font "
	boldSenFont += "-output ../../object/resources/senBoldFont.go "
	boldSenFont += "-var BoldSenFont"
	fmt.Fprintln(genFile, boldSenFont)

	surrenderBtn := mainCmd + "-input ../images/objectImages/surrenderButton.webp "
	surrenderBtn += "-package objectImages "
	surrenderBtn += "-output ../images/objectImages/surrenderButton.go "
	surrenderBtn += "-var SurrenderBtn_webp"
	fmt.Fprintln(genFile, surrenderBtn)

	defaultBtn := mainCmd + "-input ../images/objectImages/defaultButton.png "
	defaultBtn += "-package objectImages "
	defaultBtn += "-output ../images/objectImages/defaultButton.go "
	defaultBtn += "-var DefaultBtn_png"
	fmt.Fprintln(genFile, defaultBtn)

	defaultBtnSel := mainCmd + "-input ../images/objectImages/defaultButtonSelected.png "
	defaultBtnSel += "-package objectImages "
	defaultBtnSel += "-output ../images/objectImages/defaultButtonSelected.go "
	defaultBtnSel += "-var DefaultBtnSelected_png"
	fmt.Fprintln(genFile, defaultBtnSel)

	fmt.Fprintln(genFile, "\npackage main")
}
