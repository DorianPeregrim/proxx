package proxx

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	tileSize   = 20
	tileMargin = 2
	dpi        = 72
)

var backgroundColor = color.RGBA{R: 0xbb, G: 0xad, B: 0xa0, A: 0xff}

type WindowBoardRenderer struct {
}

func NewWindowBoardRenderer() *WindowBoardRenderer {
	return &WindowBoardRenderer{}
}

func (r *WindowBoardRenderer) Render(screen *ebiten.Image, b *Board) string {
	//screen.Fill(backgroundColor)

	b.iterateCells(func(i, j int, c ICell) {
		x := i*tileSize + (i+1)*tileMargin
		y := j*tileSize + (j+1)*tileMargin
		renderCell(screen, c, x, y)
	})
	return ""
}

func renderCell(screen *ebiten.Image, c ICell, x, y int) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(float64(x), float64(y))
	t := ebiten.NewImage(tileSize, tileSize)
	t.Fill(color.White)

	tt, _ := opentype.Parse(fonts.MPlus1pRegular_ttf)
	f, _ := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    16,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if c.IsOpened() {
		str := c.String()
		bound, _ := font.BoundString(f, str)
		w := (bound.Max.X - bound.Min.X).Ceil()
		h := (bound.Max.Y - bound.Min.Y).Ceil()
		x = (tileSize - w) / 2
		y = (tileSize-h)/2 + h
		text.Draw(t, str, f, x, y, color.Gray16{0x4DFF})
	}

	screen.DrawImage(t, op)
}
