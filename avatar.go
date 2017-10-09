package charavatar

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/golang/freetype/truetype"
)

type Avatar struct {
	font            *truetype.Font
	frontColor      color.Color
	backgroundColor color.Color
	width           int
	height          int
	char            rune
}

func NewAvatar(f *truetype.Font, fc, bc color.Color, w, h int, char rune) *Avatar {
	return &Avatar{f, fc, bc, w, h, char}
}

// build avatar image
func (a *Avatar) BuildImage() *Image {
	dst := NewImage(a.width, a.height)
	dst.FillBackground(image.NewUniform(a.backgroundColor))

	fontSize := int(float64(a.height) * 0.6)

	face := truetype.NewFace(a.font, &truetype.Options{Size: float64(fontSize)})
	defer face.Close()
	awidth, _ := face.GlyphAdvance(a.char)

	charWidth := int(float64(awidth) / 64)
	charHeight := fontSize

	charImg := NewImage(charWidth, charHeight)
	charImg.DrawString(a.font, a.frontColor, string(a.char), float64(fontSize))

	bs := charImg.Bounds().Size()
	left := (a.width - bs.X) / 2
	top := (a.height - bs.Y) / 2

	dst.DrawImage(image.Rect(left, top, left+bs.X, top+bs.Y), charImg, image.ZP, draw.Over)

	return dst
}
