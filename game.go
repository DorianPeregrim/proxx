package proxx

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	rowsNum       = 10
	colsNum       = 10
	blackHolesNum = 10
)

type Game struct {
	board    *Board
	renderer *WindowBoardRenderer
}

func NewGame() *Game {
	g := &Game{}
	bc := BoardConfig{
		RowsNum:       rowsNum,
		ColsNum:       colsNum,
		BlackHolesNum: blackHolesNum,
	}
	g.board = NewBoard(bc)
	g.renderer = NewWindowBoardRenderer()
	return g
}

func (g *Game) Update() error {
	if g.board.IsComplete() {
		return nil
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		i := x / (tileSize + tileMargin)
		j := y / (tileSize + tileMargin)
		if !g.board.Click(i, j) {
			g.board.Open()
			g.board.isGameOver = true
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.Render(screen, g.board)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	w := tileSize*rowsNum + tileMargin*rowsNum + tileMargin
	h := tileSize*colsNum + tileMargin*colsNum + tileMargin
	return w, h
}
