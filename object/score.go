package object

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type score struct{}

func NewScore() Object {
	return &score{}
}

func (sc *score) Update() {}

func (sc *score) Draw(screen *ebiten.Image) {
	score := fmt.Sprintf("Score: %s", *GameScore)

	// The standard pokemon size is 475x475.
	pkmX := float64(GameWidth)/2 + 475./2

	// 10 = Padding-left.
	scX := int(pkmX + 5.)
	scY := 35

	text.Draw(screen, score, standardBoldFnt.Fnt, scX, scY, color.Black)
}
