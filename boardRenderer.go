package proxx

type boardRenderer interface {
	Render(b *Board) string
}
