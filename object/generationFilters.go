package object

import (
	_ "image/png"
	"time"

	objRes "whosthatpokemon/object/resources"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
)

type generationBtn struct {
	// Both x and y needs to be pointers for the update function to work as it
	// will take 0 values otherwise.

	// Delimits the position in the X axis where the button will
	// start it's drawing.
	x *float64
	// Delimits the position in the Y axis where the button will
	// start it's drawing.
	y *float64
	// Button's Width.
	w int
	// Button's height.
	h int

	selected *bool
	// Generation's number.
	Number  int
	genText *PkmFont

	lastClickAt *time.Time
	btnRepr     *button
}

var (
	genFilterText   Object
	GenerationsText = []string{
		"First Generation",
		"Second Generation",
		"Third Generation",
		"Forth Generation",
		"Fifth Generation",
		"Sixth Generation",
		"Seventh Generation",
		"Eight Generation",
		"Ninth Generation",
	}
)

func init() {
	opt := &truetype.Options{Size: TxtStandardFontSize}
	genFilterText = NewFont(objRes.SenFont, opt, "Generation Filter:", 20, 35)
}

func (geneBtn generationBtn) Update() {
	cursorX, cursorY := ebiten.CursorPosition()

	buttonIsClicked(cursorX, cursorY, geneBtn.btnRepr, func() {
		if !*geneBtn.selected {
			SelectedGenerations[geneBtn.Number] = struct{}{}
		} else {
			delete(SelectedGenerations, geneBtn.Number)
		}

		*geneBtn.selected = !*geneBtn.selected
	})
}

func (geneBtn generationBtn) Draw(screen *ebiten.Image) {
	*geneBtn.x = 20.
	// 25 = Height of the button. 10 for padding-bottom.
	*geneBtn.y = float64(50 + (25+10)*(geneBtn.Number-1))

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(*geneBtn.x, *geneBtn.y)

	if *geneBtn.selected {
		screen.DrawImage(filterBtnSelImg, opts)
	} else {
		screen.DrawImage(filterBtnImg, opts)
	}

	genFilterText.Draw(screen)
	geneBtn.genText.Draw(screen)
}

func createGenText(genNum int) *PkmFont {
	// 25 = Height of the button. 10 for padding-bottom.
	geneTextHeight := 70 + (25+10)*genNum

	opt := &truetype.Options{Size: TxtSmallFontSize}
	genText := GenerationsText[genNum]
	generationText := NewFont(objRes.SenFont, opt, genText, 40, geneTextHeight)

	return generationText.(*PkmFont)
}

func InitializeGenerationBtns() []Object {
	var genBtns []Object
	for i := range GenerationsText {
		genText := createGenText(i)

		newGenBtn := &generationBtn{
			x:           new(float64),
			y:           new(float64),
			w:           200,
			h:           25,
			Number:      i + 1,
			lastClickAt: &time.Time{},
			selected:    new(bool),
			genText:     genText,
		}
		newGenBtn.btnRepr = &button{
			x:           newGenBtn.x,
			y:           newGenBtn.y,
			w:           &newGenBtn.w,
			h:           &newGenBtn.h,
			lastClickAt: newGenBtn.lastClickAt,
		}
		genBtns = append(genBtns, newGenBtn)
	}

	return genBtns
}
