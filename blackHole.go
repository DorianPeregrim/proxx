package proxx

type BlackHole struct {
	Cell
}

func NewBlackHole() *BlackHole {
	return &BlackHole{
		Cell{
			isOpened: false,
		},
	}
}

func (bh *BlackHole) Click() bool {
	bh.Open()
	return false
}

func (bh *BlackHole) String() string {
	return "H"
}

func (bh *BlackHole) IsBlackHole() bool {
	return true
}
