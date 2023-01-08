package proxx

type Cell struct {
	neighbors []ICell
	isOpened  bool
}

func (c *Cell) IsOpened() bool {
	return c.isOpened
}

func (c *Cell) String() string {
	return "0"
}

func (c *Cell) Click() bool {
	if !c.IsOpened() {
		c.Open()
		for _, nc := range c.neighbors {
			nc.Click()
		}
	}

	return true
}

func (c *Cell) SetNeighbors(ns []ICell) {
	c.neighbors = ns
}

func (c *Cell) Open() {
	c.isOpened = true
}

func (c *Cell) IsBlackHole() bool {
	return false
}
