package proxx

import (
	"math/rand"
)

const (
	maxRowsNum = 9
	maxColsNum = 9
)

type Board struct {
	config BoardConfig
	cells  map[int]map[int]ICell
}

func NewBoard(conf BoardConfig) *Board {
	b := &Board{
		config: conf,
		cells:  make(map[int]map[int]ICell),
	}
	b.fill()
	return b
}

func (b *Board) Click(i, j int) bool {
	c := b.getCell(i, j)
	return c.Click()
}

func (b *Board) Open() {
	b.iterateCells(func(i, j int, c ICell) {
		c.Open()
	})
}

func (b *Board) IsComplete() bool {
	count := 0
	b.iterateCells(func(i, j int, c ICell) {
		if !c.IsBlackHole() && c.IsOpened() {
			count++
		}
	})
	return count == b.countCells()-b.config.BlackHolesNum
}

func (b *Board) fill() {
	holes := generateHoles(b.config.BlackHolesNum, b.countCells())

	b.iterateCells(func(i, j int, c ICell) {
		if h, ok := holes[b.config.ColsNum*i+j]; ok {
			b.setCell(i, j, h)
		}
	})

	b.iterateCells(func(i, j int, c ICell) {
		if c == nil {
			count := b.countNeighboringHoles(i, j)
			b.setCell(i, j, NewEmptyCell(count))
		}
	})

	b.iterateCells(func(i, j int, c ICell) {
		c.SetNeighbors(b.findNeighbors(i, j))
	})
}

func (b *Board) countCells() int {
	return b.config.RowsNum * b.config.ColsNum
}

func (b *Board) getCell(rowIndex, colIndex int) ICell {
	row, ok := b.cells[rowIndex]
	if ok {
		return row[colIndex]
	}
	return nil
}

func (b *Board) setCell(rowIndex, colIndex int, c ICell) {
	row, ok := b.cells[rowIndex]
	if !ok {
		row = make(map[int]ICell)
		b.cells[rowIndex] = row
	}
	row[colIndex] = c
}

func (b *Board) countNeighboringHoles(rowIndex, colIndex int) int {
	count := 0
	ns := b.findNeighbors(rowIndex, colIndex)
	for _, c := range ns {
		switch c.(type) {
		case *BlackHole:
			count++
		}
	}
	return count
}

func (b *Board) findNeighbors(rowIndex, colIndex int) []ICell {
	var ns []ICell
	for i := -1; i <= 1; i++ {
		if row, ok := b.cells[rowIndex+i]; ok {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}
				if c, ok := row[colIndex+j]; ok {
					ns = append(ns, c)
				}
			}
		}
	}
	return ns
}

func (b *Board) iterateCells(callback func(i, j int, c ICell)) {
	for i := 0; i < b.config.RowsNum; i++ {
		for j := 0; j < b.config.ColsNum; j++ {
			c := b.getCell(i, j)
			callback(i, j, c)
		}
	}
}

func generateHoles(holesNum, maxIndex int) map[int]*BlackHole {
	holes := make(map[int]*BlackHole)
	for i := 0; i < holesNum; i++ {
		hi := generateHoleIndex(holes, maxIndex)
		holes[hi] = NewBlackHole()
	}
	return holes
}

func generateHoleIndex(holes map[int]*BlackHole, maxIndex int) int {
	i := rand.Intn(maxIndex)
	for hi := range holes {
		if hi == i {
			return generateHoleIndex(holes, maxIndex)
		}
	}
	return i
}
