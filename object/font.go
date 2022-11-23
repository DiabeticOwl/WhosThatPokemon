package object

import (
	"image/color"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

var (
	SenFnt font.Face
)

type PkmFont struct {
	Fnt  font.Face
	Text string
	// Coordinates where the text will be drawn.
	X int
	Y int
}

func (f *PkmFont) Update() {}
func (f *PkmFont) Draw(screen *ebiten.Image) {
	// Note: text.Draw doesn't show any color different than color.Black and
	// color.White.
	// Note: Ebiten will panic if the font brings any color other than
	// color.Black.
	text.Draw(screen, f.Text, f.Fnt, f.X, f.Y, color.Black)
}

func NewFont(ttf []byte, opts *truetype.Options, text string, x, y int) Object {
	tt, _ := truetype.Parse(ttf)
	fnt := truetype.NewFace(tt, opts)

	return &PkmFont{
		Fnt:  fnt,
		Text: text,
		X:    x,
		Y:    y,
	}
}
