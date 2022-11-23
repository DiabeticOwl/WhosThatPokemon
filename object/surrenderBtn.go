package object

import (
	"bytes"
	"image"
	"time"

	_ "golang.org/x/image/webp"

	"github.com/hajimehoshi/ebiten/v2"

	"whosthatpokemon/resources/images/objectImages"
)

type surrenderBtn struct {
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

	btnText *PkmFont

	lastClickAt *time.Time
}

var (
	surrenderBtnImg *ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(objectImages.SurrenderBtn_webp))
	if err != nil {
		panic(err)
	}

	surrenderBtnImg = ebiten.NewImageFromImage(img)
}

func (sBtn surrenderBtn) Update() {
	// sBtn.btnText.Update()
	cursorX, cursorY := ebiten.CursorPosition()

	btn := &button{
		x:           sBtn.x,
		y:           sBtn.y,
		w:           &sBtn.w,
		h:           &sBtn.h,
		lastClickAt: sBtn.lastClickAt,
	}

	buttonIsClicked(cursorX, cursorY, btn, func() {
		*SBtnClicked = true
	})
}

func (sBtn surrenderBtn) Draw(screen *ebiten.Image) {
	sW, sH := screen.Size()

	// 475 being the standard width and height of the pokemon images.
	*sBtn.x = float64(sW)/2 - 475./2
	*sBtn.y = float64(sH)/2 + 300

	// Note: Ebitengine won't display another image that is filled with a color
	// after the screen is already filled with another.
	// Tested this with `ebiten.NewImage()`, and by creating a new image from
	// `image.Image` and then `ebiten.NewImageFromImage()`.

	opts := &ebiten.DrawImageOptions{}
	// opts.GeoM.Scale(1, 1)
	opts.GeoM.Translate(*sBtn.x, *sBtn.y)

	screen.DrawImage(surrenderBtnImg, opts)

	sBtn.btnText.Text = "I give up."
	sBtn.btnText.X = int(*sBtn.x + float64(sBtn.w)/2.8)
	sBtn.btnText.Y = int(*sBtn.y + float64(sBtn.h)/1.5)
	sBtn.btnText.Draw(screen)
}

func NewSurrenderBtn() Object {
	return &surrenderBtn{
		btnText:     standardFnt,
		x:           new(float64),
		y:           new(float64),
		w:           475,
		h:           50,
		lastClickAt: &time.Time{},
	}
}
