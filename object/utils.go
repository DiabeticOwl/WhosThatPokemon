package object

import (
	"bytes"
	"image"
	"image/color"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type button struct {
	// Both x and y needs to be pointers for the update function to work as it
	// will take 0 values otherwise.

	// Delimits the position in the X axis where the button will
	// start it's drawing.
	x *float64
	// Delimits the position in the Y axis where the button will
	// start it's drawing.
	y *float64
	// Button's Width.
	w *int
	// Button's height.
	h           *int
	lastClickAt *time.Time
}

func decodeLocalImage(obj []byte) image.Image {
	img, _, err := image.Decode(bytes.NewReader(obj))
	if err != nil {
		panic(err)
	}

	return img
}

// buttonIsClicked will estimate how close a click on the game was using the
// formula for distance between points. Then it will execute the passed function
// if the click was close enough to dstCond.
func buttonIsClicked(cursorX, cursorY int, btn *button, toExec func()) {
	// Calculate whether a given click was unique or not and if it as within
	// the borders of the button.
	now := time.Now()
	leftButtonPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if leftButtonPressed && now.Sub(*btn.lastClickAt) > Debouncer {
		dstCond := *btn.w + 5
		*btn.lastClickAt = now

		btnXLimit := *btn.x + float64(*btn.w)
		btnYLimit := *btn.y + float64(*btn.h)

		xDeltaPow2 := math.Pow(*btn.x-float64(cursorX), 2)
		yDeltaPow2 := math.Pow(*btn.y-float64(cursorY), 2)
		distance := int(math.Sqrt(xDeltaPow2 + yDeltaPow2))

		cursorXInLimit := float64(cursorX) >= *btn.x
		cursorXInLimit = cursorXInLimit && float64(cursorX) <= btnXLimit
		cursorYInLimit := float64(cursorY) >= *btn.y
		cursorYInLimit = cursorYInLimit && float64(cursorY) <= btnYLimit

		// 5 is arbitrary padding.
		if distance <= dstCond && cursorXInLimit && cursorYInLimit {
			toExec()
		}
	}
}

func ChangeImageColor(img image.Image, rgb [3]uint8) image.Image {
	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	newImg := image.NewRGBA(rect)

	// Replacing pokemon image's pixels with black images while maintaining
	// their original alpha value.
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			_, _, _, originalA := img.At(x, y).RGBA()

			c := color.RGBA{
				R: rgb[0], G: rgb[1], B: rgb[2], A: uint8(originalA),
			}
			newImg.Set(x, y, c)
		}
	}

	return newImg
}
