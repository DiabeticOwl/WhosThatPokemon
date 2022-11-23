//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -compress -input ../json/pokemonDetails.json -package pokemon -output ../../pokemon/pokemonDetails.go -var pkmJson
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -compress -input ../json/pokemonSpecies.json -package pokemon -output ../../pokemon/pokemonSpecies.go -var spcJson
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -compress -input ../json/generationDetails.json -package pokemon -output ../../pokemon/generationDetails.go -var genJson
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -compress -input ../json/typeDetails.json -package pokemon -output ../../pokemon/typeDetails.go -var typeJson
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -compress -input ../json/gameDetails.json -package pokemon -output ../../pokemon/gameDetails.go -var gameJson
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -input ../fonts/Sen/Sen-Regular.ttf -package font -output ../../object/resources/senRegularFont.go -var SenFont
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -input ../fonts/Sen/Sen-ExtraBold.ttf -package font -output ../../object/resources/senBoldFont.go -var BoldSenFont
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -input ../images/objectImages/surrenderButton.webp -package objectImages -output ../images/objectImages/surrenderButton.go -var SurrenderBtn_webp
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -input ../images/objectImages/defaultButton.png -package objectImages -output ../images/objectImages/defaultButton.go -var DefaultBtn_png
//go:generate go run github.com/hajimehoshi/file2byteslice/cmd/file2byteslice -input ../images/objectImages/defaultButtonSelected.png -package objectImages -output ../images/objectImages/defaultButtonSelected.go -var DefaultBtnSelected_png

package main
