package object

import (
	objRes "whosthatpokemon/object/resources"

	"code.rocketnine.space/tslocum/messeji"
	"github.com/golang/freetype/truetype"
)

func NewInputField() *messeji.InputField {
	tt, _ := truetype.Parse(objRes.SenFont)
	fnt := truetype.NewFace(tt, &truetype.Options{Size: TxtStandardFontSize})

	input := messeji.NewInputField(fnt)
	input.TextField.SetHorizontal(messeji.AlignCenter)
	input.TextField.SetVertical(messeji.AlignCenter)

	input.SetHandleKeyboard(true)

	return input
}
