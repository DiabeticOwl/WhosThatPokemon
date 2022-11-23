package object

import (
	_ "image/png"
	"time"

	objRes "whosthatpokemon/object/resources"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
)

type typeOpts struct {
	// ID from the type JSON.
	// Extracted manually from: https://pokeapi.co/api/v2/type
	ID int
	// Position of the column. 1 or 2.
	columnPos int
	// X and Y.
	coordinates [2]float64
	// RGB color.
	Color [3]uint8
}

type typeBtn struct {
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

	typeOpt *typeOpts

	selected *bool

	typeText    *PkmFont
	SelTypeText *PkmFont

	lastClickAt *time.Time
}

var (
	typeFilterText *PkmFont
	TypesTextMap   = make(map[string]*typeOpts)
)

func init() {
	opt := &truetype.Options{Size: TxtStandardFontSize}

	text := "Pokemon Type:"
	typeFilterText = NewFont(objRes.SenFont, opt, text, 0, 80).(*PkmFont)

	// Visible description of the currently 18 types of pokemon.
	// Coordinates are crafted by giving an manual Y position with an offset of
	// 35 and then calculate the X position given the initialX defined at draw.
	// Types are distributed in two column per row format.
	// Colors extracted from: https://bulbapedia.bulbagarden.net/wiki/Type
	TypesTextMap = map[string]*typeOpts{
		"Normal": {
			ID:          1,
			columnPos:   1,
			coordinates: [2]float64{0, 90},
			Color:       [3]uint8{168, 168, 120},
		},
		"Fire": {
			ID:          10,
			columnPos:   2,
			coordinates: [2]float64{0, 90},
			Color:       [3]uint8{240, 128, 48},
		},
		"Fighting": {
			ID:          2,
			columnPos:   1,
			coordinates: [2]float64{0, 125},
			Color:       [3]uint8{192, 48, 40},
		},
		"Water": {
			ID:          11,
			columnPos:   2,
			coordinates: [2]float64{0, 125},
			Color:       [3]uint8{104, 144, 240},
		},
		"Flying": {
			ID:          3,
			columnPos:   1,
			coordinates: [2]float64{0, 160},
			Color:       [3]uint8{168, 144, 240},
		},
		"Grass": {
			ID:          12,
			columnPos:   2,
			coordinates: [2]float64{0, 160},
			Color:       [3]uint8{120, 200, 80},
		},
		"Poison": {
			ID:          4,
			columnPos:   1,
			coordinates: [2]float64{0, 195},
			Color:       [3]uint8{160, 64, 160},
		},
		"Electric": {
			ID:          13,
			columnPos:   2,
			coordinates: [2]float64{0, 195},
			Color:       [3]uint8{248, 208, 48},
		},
		"Ground": {
			ID:          5,
			columnPos:   1,
			coordinates: [2]float64{0, 230},
			Color:       [3]uint8{224, 192, 104},
		},
		"Psychic": {
			ID:          14,
			columnPos:   2,
			coordinates: [2]float64{0, 230},
			Color:       [3]uint8{248, 88, 136},
		},
		"Rock": {
			ID:          6,
			columnPos:   1,
			coordinates: [2]float64{0, 265},
			Color:       [3]uint8{248, 88, 136},
		},
		"Ice": {
			ID:          15,
			columnPos:   2,
			coordinates: [2]float64{0, 265},
			Color:       [3]uint8{152, 216, 216},
		},
		"Bug": {
			ID:          7,
			columnPos:   1,
			coordinates: [2]float64{0, 300},
			Color:       [3]uint8{168, 184, 32},
		},
		"Dragon": {
			ID:          16,
			columnPos:   2,
			coordinates: [2]float64{0, 300},
			Color:       [3]uint8{112, 56, 248},
		},
		"Ghost": {
			ID:          8,
			columnPos:   1,
			coordinates: [2]float64{0, 335},
			Color:       [3]uint8{112, 88, 152},
		},
		"Dark": {
			ID:          17,
			columnPos:   2,
			coordinates: [2]float64{0, 335},
			Color:       [3]uint8{112, 88, 72},
		},
		"Steel": {
			ID:          9,
			columnPos:   1,
			coordinates: [2]float64{0, 370},
			Color:       [3]uint8{184, 184, 208},
		},
		"Fairy": {
			ID:          18,
			columnPos:   2,
			coordinates: [2]float64{0, 370},
			Color:       [3]uint8{238, 153, 172},
		},
	}
}

func (typeBtn typeBtn) Update() {
	cursorX, cursorY := ebiten.CursorPosition()

	btn := &button{
		x:           typeBtn.x,
		y:           typeBtn.y,
		w:           &typeBtn.w,
		h:           &typeBtn.h,
		lastClickAt: typeBtn.lastClickAt,
	}

	buttonIsClicked(cursorX, cursorY, btn, func() {
		if !*typeBtn.selected {
			SelectedTypes[typeBtn.typeOpt.ID] = struct{}{}
		} else {
			delete(SelectedTypes, typeBtn.typeOpt.ID)
		}

		*typeBtn.selected = !*typeBtn.selected
	})
}

func (typeBtn typeBtn) Draw(screen *ebiten.Image) {
	sW, _ := screen.Size()
	initialX := float64(sW)/2 + 475./2 + 5

	// Calculating how much the image needs to be scaled in order to be resized
	// to the object's width and height.
	imgSize := filterBtnImg.Bounds().Size()
	// Multiple float64 conversion in order to maintain the precision.
	imgSizeX := float64(imgSize.X)
	imgSizeY := float64(imgSize.Y)
	imgXScale := float64(typeBtn.w) / imgSizeX
	imgYScale := float64(typeBtn.h) / imgSizeY

	if typeBtn.typeOpt.columnPos%2 == 0 {
		typeBtn.typeOpt.coordinates[0] = initialX + imgSizeX/2 + 10
	} else {
		typeBtn.typeOpt.coordinates[0] = initialX
	}

	col := typeBtn.typeOpt.Color
	coordX := typeBtn.typeOpt.coordinates[0]
	coordY := typeBtn.typeOpt.coordinates[1]

	// For the clicking function.
	*typeBtn.x = coordX
	*typeBtn.x = coordX

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(imgXScale, imgYScale)
	opts.GeoM.Translate(coordX, coordY)

	typeSelImg := ChangeImageColor(filterBtnImg, col)
	typeSel := ebiten.NewImageFromImage(typeSelImg)
	screen.DrawImage(typeSel, opts)

	if *typeBtn.selected {
		typeBtn.SelTypeText.X = int(coordX + 10)
		typeBtn.SelTypeText.Y = int(coordY + 18)
		typeBtn.SelTypeText.Draw(screen)
	} else {
		typeBtn.typeText.X = int(coordX + 10)
		typeBtn.typeText.Y = int(coordY + 18)
		typeBtn.typeText.Draw(screen)
	}

	typeFilterText.X = int(initialX)
	typeFilterText.Draw(screen)
}

func createTypeText(text string, fnt []byte) *PkmFont {
	opt := &truetype.Options{Size: TxtSmallFontSize}
	typeText := NewFont(fnt, opt, text, 0, 0)

	return typeText.(*PkmFont)
}

func InitializeTypeBtns() []Object {
	var genBtns []Object
	for typKey, typOpt := range TypesTextMap {
		coordY := typOpt.coordinates[1]

		typText := createTypeText(typKey, objRes.SenFont)
		selTypText := createTypeText(typKey, objRes.BoldSenFont)

		newGenBtn := &typeBtn{
			x:           new(float64),
			y:           new(float64),
			w:           100,
			h:           25,
			lastClickAt: &time.Time{},
			selected:    new(bool),
			typeOpt:     typOpt,
			typeText:    typText,
			SelTypeText: selTypText,
		}
		*newGenBtn.y = coordY
		genBtns = append(genBtns, newGenBtn)
	}

	return genBtns
}
