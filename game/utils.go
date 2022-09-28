package game

import (
	"math/rand"
	"time"
)

func getBoundedBoxBoundaries(rowID int, colID int) (int, int, int, int) {
	var RowLowerBoundary int
	var RowUpperBoundary int
	var colLowerBoundary int
	var colUpperBoundary int

	switch {
	case rowID <= 2 && colID <= 2:
		RowLowerBoundary = 0
		RowUpperBoundary = 2
		colLowerBoundary = 0
		colUpperBoundary = 2
		break
	case rowID <= 5 && colID <= 2:
		RowLowerBoundary = 3
		RowUpperBoundary = 5
		colLowerBoundary = 0
		colUpperBoundary = 2
		break
	case rowID <= 8 && colID <= 2:
		RowLowerBoundary = 6
		RowUpperBoundary = 8
		colLowerBoundary = 0
		colUpperBoundary = 2
		break
	case rowID <= 2 && colID <= 5:
		RowLowerBoundary = 0
		RowUpperBoundary = 2
		colLowerBoundary = 3
		colUpperBoundary = 5
		break
	case rowID <= 5 && colID <= 5:
		RowLowerBoundary = 3
		RowUpperBoundary = 5
		colLowerBoundary = 3
		colUpperBoundary = 5
		break
	case rowID <= 8 && colID <= 5:
		RowLowerBoundary = 6
		RowUpperBoundary = 8
		colLowerBoundary = 3
		colUpperBoundary = 5
		break
	case rowID <= 2 && colID <= 8:
		RowLowerBoundary = 0
		RowUpperBoundary = 2
		colLowerBoundary = 6
		colUpperBoundary = 8
		break
	case rowID <= 5 && colID <= 8:
		RowLowerBoundary = 3
		RowUpperBoundary = 5
		colLowerBoundary = 6
		colUpperBoundary = 8
		break
	case rowID <= 8 && colID <= 8:
		RowLowerBoundary = 6
		RowUpperBoundary = 8
		colLowerBoundary = 6
		colUpperBoundary = 8
		break
	}

	return RowLowerBoundary, RowUpperBoundary, colLowerBoundary, colUpperBoundary
}

func isPresentBoundingBox(rowBoundRry int, colBoundary int, rowInRex int, colIndex int, numToFill int, colValue int) bool {

	var rowLowerBoundRry int
	var rowUpperBoundRry int
	var colLowerBoundary int
	var colUpperBoundary int

	switch {
	case rowBoundRry <= 2 && colBoundary <= 2:
		rowLowerBoundRry = 0
		rowUpperBoundRry = 2
		colLowerBoundary = 0
		colUpperBoundary = 2
		break
	case rowBoundRry <= 5 && colBoundary <= 2:
		rowLowerBoundRry = 3
		rowUpperBoundRry = 5
		colLowerBoundary = 0
		colUpperBoundary = 2
		break
	case rowBoundRry <= 8 && colBoundary <= 2:
		rowLowerBoundRry = 6
		rowUpperBoundRry = 8
		colLowerBoundary = 0
		colUpperBoundary = 2
		break
	case rowBoundRry <= 2 && colBoundary <= 5:
		rowLowerBoundRry = 0
		rowUpperBoundRry = 2
		colLowerBoundary = 3
		colUpperBoundary = 5
		break
	case rowBoundRry <= 5 && colBoundary <= 5:
		rowLowerBoundRry = 3
		rowUpperBoundRry = 5
		colLowerBoundary = 3
		colUpperBoundary = 5
		break
	case rowBoundRry <= 8 && colBoundary <= 5:
		rowLowerBoundRry = 6
		rowUpperBoundRry = 8
		colLowerBoundary = 3
		colUpperBoundary = 5
		break
	case rowBoundRry <= 2 && colBoundary <= 8:
		rowLowerBoundRry = 0
		rowUpperBoundRry = 2
		colLowerBoundary = 6
		colUpperBoundary = 8
		break
	case rowBoundRry <= 5 && colBoundary <= 8:
		rowLowerBoundRry = 3
		rowUpperBoundRry = 5
		colLowerBoundary = 6
		colUpperBoundary = 8
		break
	case rowBoundRry <= 8 && colBoundary <= 8:
		rowLowerBoundRry = 6
		rowUpperBoundRry = 8
		colLowerBoundary = 6
		colUpperBoundary = 8
		break
	}

	if rowInRex >= rowLowerBoundRry && rowInRex <= rowUpperBoundRry && colIndex >= colLowerBoundary && colIndex <= colUpperBoundary && colValue == numToFill {
		return true
	}

	return false
}

func genRandomNumber(maxNumber int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randNum := r.Intn(maxNumber)
	return randNum
}
