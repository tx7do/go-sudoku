package game

type Position struct {
	Col, Row int
}

func NewPosition(col, row int) *Position {
	return &Position{Col: col, Row: row}
}

func (p Position) IsPosition(col, row int) bool {
	return p.Col == col && p.Row == row
}
