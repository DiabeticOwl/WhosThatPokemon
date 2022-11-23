package object

import (
	"time"
	objRes "whosthatpokemon/object/resources"
	"whosthatpokemon/resources/images/objectImages"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
)

type Object interface {
	Draw(screen *ebiten.Image)
	Update()
}

const (
	// Used for delimiting how long a click can take. If the user presses for
	// more than 200 milliseconds then the game will see it as an unique click.
	Debouncer = 200 * time.Millisecond

	TxtSmallFontSize    = 18
	TxtStandardFontSize = 25
)

var (
	SBtnClicked           *bool
	GameHeight, GameWidth int
	GameScore             *string
	SelectedGenerations   = make(map[int]struct{})
	SelectedTypes         = make(map[int]struct{})
	standardFnt           *PkmFont
	standardBoldFnt       *PkmFont

	filterBtnImg    *ebiten.Image
	filterBtnSelImg *ebiten.Image
)

func init() {
	SBtnClicked = new(bool)

	opt := &truetype.Options{Size: TxtStandardFontSize}
	standardFnt = NewFont(objRes.SenFont, opt, "", 0, 0).(*PkmFont)
	standardBoldFnt = NewFont(objRes.BoldSenFont, opt, "", 0, 0).(*PkmFont)

	img := decodeLocalImage(objectImages.DefaultBtn_png)
	filterBtnImg = ebiten.NewImageFromImage(img)
	img = decodeLocalImage(objectImages.DefaultBtnSelected_png)
	filterBtnSelImg = ebiten.NewImageFromImage(img)
}
