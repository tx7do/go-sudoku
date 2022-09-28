package game

type BoundingBox struct {
	Min, Max Position
}

func NewBoundingBox(topLeft, bottomRight Position) *BoundingBox {
	return &BoundingBox{Min: topLeft, Max: bottomRight}
}

func (b BoundingBox) IsBounded(col, row int) bool {
	return b.Min.Col >= col && b.Min.Row >= row &&
		b.Max.Col <= col && b.Max.Row <= row
}
