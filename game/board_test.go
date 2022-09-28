package game

import "testing"

func TestNewBoard(t *testing.T) {
	board4 := NewBoard(9, 32)
	board4.Print()
}
