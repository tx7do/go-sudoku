package game

import "fmt"

type Row []int

func NewRow(numbersToFill int) Row {
	Row := Row{}

	for i := 0; i < numbersToFill; i++ {
		Row = append(Row, 0)
	}

	return Row
}

func (r Row) Print() {
	for _, col := range r {
		fmt.Println(col)
	}
}

func (r Row) HasSameNumber(number int) bool {
	for _, col := range r {
		if number == r[col] {
			return true
		}
	}
	return false
}

func (r Row) SetNumber(col, number int) {
	r[col] = number
}

func (r Row) GetNumber(col int) int {
	return r[col]
}
