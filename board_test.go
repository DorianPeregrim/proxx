package proxx

import "testing"

func TestNewBoard(t *testing.T) {
	b := createBoard()
	b.iterateCells(func(i, j int, c ICell) {
		if c.IsOpened() {
			t.Errorf("TestNewBoard: has opened cell on i=%d j=%d", i, j)
		}
	})
}

func TestBoard_Click(t *testing.T) {
	b := createBoard()
	i := 1
	j := 1
	c := b.getCell(i, j)

	if c.IsOpened() {
		t.Errorf("TestBoard_Click: cell i=%d j=%d is opened", i, j)
	}

	b.Click(i, j)
	if !c.IsOpened() {
		t.Errorf("TestBoard_Click: cell i=%d j=%d is closed", i, j)
	}
}

func TestBoard_Open(t *testing.T) {
	b := createBoard()
	b.iterateCells(func(i, j int, c ICell) {
		if c.IsOpened() {
			t.Errorf("TestBoard_Open: has opened cell on i=%d j=%d", i, j)
		}
	})

	b.Open()
	b.iterateCells(func(i, j int, c ICell) {
		if !c.IsOpened() {
			t.Errorf("TestBoard_Open: has closed cell on i=%d j=%d", i, j)
		}
	})
}

func TestBoard_IsComplete(t *testing.T) {
	b := createBoard()
	if b.IsComplete() {
		t.Errorf("TestBoard_IsComplete: board is complete")
	}

	b.iterateCells(func(i, j int, c ICell) {
		if !c.IsBlackHole() {
			c.Open()
		}
	})

	if !b.IsComplete() {
		t.Errorf("TestBoard_IsComplete: board is not complete")
	}
}

func createBoard() *Board {
	bc := BoardConfig{
		RowsNum:       3,
		ColsNum:       3,
		BlackHolesNum: 3,
	}

	return NewBoard(bc)
}
