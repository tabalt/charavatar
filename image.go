package charavatar

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

type Image struct {
	*image.RGBA
}

func NewImage(w, h int) *Image {
	return &Image{image.NewRGBA(image.Rect(0, 0, w, h))}
}

// draw image
func (img *Image) DrawImage(r image.Rectangle, src image.Image, sp image.Point, op draw.Op) {
	draw.Draw(img, r, src, sp, op)
}

// draw string
func (img *Image) DrawString(font *truetype.Font, c color.Color, str string, fontsize float64) {
	ctx := freetype.NewContext() // default 72dpi
	ctx.SetDst(img)
	ctx.SetClip(img.Bounds())
	ctx.SetSrc(image.NewUniform(c))
	ctx.SetFontSize(fontsize)
	ctx.SetFont(font)

	pt := freetype.Pt(0, int(-fontsize/6)+ctx.PointToFixed(fontsize).Ceil())
	ctx.DrawString(str, pt)
}

// fill background
func (img *Image) FillBackground(b image.Image) {
	img.DrawImage(img.Bounds(), b, image.ZP, draw.Over)
}
