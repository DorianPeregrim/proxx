package proxx

import (
	"fmt"

	"github.com/fatih/color"
)

type BoardRenderer struct {
}

func NewBoardRenderer() *BoardRenderer {
	return &BoardRenderer{}
}

func (br *BoardRenderer) Render(b *Board) string {
	var out string
	b.iterateCells(func(i, j int, c ICell) {
		if i == 0 && j == 0 {
			for k := 0; k <= b.config.ColsNum; k++ {
				out += color.YellowString(fmt.Sprintf(" %d ", k))
			}
			out += "\n"
		}

		if j == 0 {
			out += color.YellowString(fmt.Sprintf(" %d ", i+1))
		}

		var char string
		if c.IsOpened() {
			if c.IsBlackHole() {
				char = color.RedString(b.cells[i][j].String())
			} else {
				char = color.GreenString(b.cells[i][j].String())
			}
		} else {
			char = "X"
		}

		out += fmt.Sprintf(" %s ", char)

		if j+1 == b.config.ColsNum {
			out += "\n"
		}
	})
	return out
}
