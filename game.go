package proxx

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	board    *Board
	renderer *WindowBoardRenderer
}

func NewGame() *Game {
	g := &Game{}
	bc := BoardConfig{
		RowsNum:       10,
		ColsNum:       10,
		BlackHolesNum: 10,
	}
	g.board = NewBoard(bc)
	g.renderer = NewWindowBoardRenderer()
	return g
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		i := x / (tileSize + tileMargin)
		j := y / (tileSize + tileMargin)
		g.board.Click(i, j)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.Render(screen, g.board)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
