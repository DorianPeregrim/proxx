package proxx

type ICell interface {
	SetNeighbors([]ICell)
	IsBlackHole() bool
	IsOpened() bool
	Open()
	Click() bool
	String() string
}
