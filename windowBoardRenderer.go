package proxx

import (
	"image/color"
	"io"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	tileSize   = 25
	tileMargin = 2
)

type WindowBoardRenderer struct {
	cellFont        font.Face
	alertFont       font.Face
	shadowAlertFont font.Face
}

var parsedFont *opentype.Font

func NewWindowBoardRenderer() *WindowBoardRenderer {
	return &WindowBoardRenderer{
		cellFont:        createFont(12),
		alertFont:       createFont(18),
		shadowAlertFont: createFont(18.2),
	}
}

func createFont(size float64) font.Face {
	if parsedFont == nil {
		fontFile, _ := os.Open("../assets/Roboto-Regular.ttf")
		fontData, _ := io.ReadAll(fontFile)
		parsedFont, _ = opentype.Parse(fontData)
	}

	op := &opentype.FaceOptions{
		Size:    size,
		DPI:     96,
		Hinting: font.HintingFull,
	}
	f, _ := opentype.NewFace(parsedFont, op)
	return f
}

func (r *WindowBoardRenderer) Render(screen *ebiten.Image, b *Board) {
	b.iterateCells(func(i, j int, c ICell) {
		x := i*tileSize + (i+1)*tileMargin
		y := j*tileSize + (j+1)*tileMargin
		r.renderCell(screen, c, x, y)
	})

	if b.isGameOver {
		r.renderText(screen, "GAME OVER", color.RGBA{R: 255, A: 255}, r.alertFont, true)
	} else if b.IsComplete() {
		r.renderText(screen, "YOU WIN!!!", color.RGBA{G: 255, A: 255}, r.alertFont, true)
	}
}

func (r *WindowBoardRenderer) renderCell(screen *ebiten.Image, c ICell, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))

	t := ebiten.NewImage(tileSize, tileSize)
	t.Fill(color.White)

	if c.IsOpened() {
		str := c.String()
		r.renderText(t, str, color.RGBA{A: 255}, r.cellFont, false)
	}

	screen.DrawImage(t, op)
}

func (r *WindowBoardRenderer) renderText(screen *ebiten.Image, txt string, c color.Color, f font.Face, bg bool) {
	bound, _ := font.BoundString(f, txt)
	w := (bound.Max.X - bound.Min.X).Ceil()
	h := (bound.Max.Y - bound.Min.Y).Ceil()

	if bg {
		tb, op := renderTextBackground(screen, w+10, h+10)
		sw, sh := tb.Size()
		x := (sw - w) / 2
		y := (sh-h)/2 + h
		text.Draw(tb, txt, f, x, y, c)
		screen.DrawImage(tb, op)
	} else {
		sw, sh := screen.Size()
		x := (sw - w) / 2
		y := (sh-h)/2 + h
		text.Draw(screen, txt, f, x, y, c)
	}
}

func renderTextBackground(screen *ebiten.Image, w, h int) (*ebiten.Image, *ebiten.DrawImageOptions) {
	sw, sh := screen.Size()
	x := (sw - w) / 2
	y := (sh - h) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))

	i := ebiten.NewImage(w, h)
	i.Fill(color.White)

	return i, op
}
