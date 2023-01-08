package proxx

import "strconv"

type EmptyCell struct {
	Cell
	CountNeighboringHoles int
}

func NewEmptyCell(countHoles int) *EmptyCell {
	return &EmptyCell{
		Cell: Cell{
			isOpened: false,
		},
		CountNeighboringHoles: countHoles,
	}
}

func (ec *EmptyCell) Click() bool {
	if ec.CountNeighboringHoles > 0 {
		ec.Open()
		return true
	}
	return ec.Cell.Click()
}

func (ec *EmptyCell) String() string {
	return strconv.Itoa(ec.CountNeighboringHoles)
}
