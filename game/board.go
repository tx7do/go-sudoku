package game

import "fmt"

type Board []Row

func NewBoard(cellSize, numbersToFill int) *Board {
	board := &Board{}

	for i := 0; i < cellSize; i++ {
		*board = append(*board, NewRow(cellSize))
	}

	board.Generate(numbersToFill)

	return board
}

func (s Board) CellSize() int {
	return len(s)
}

func (s Board) Generate(numbersToFill int) {
	cellSize := s.CellSize()
	var randNum int
	for numbersToFill > 0 {
		colIndex := genRandomNumber(cellSize - 1)
		rowIndex := genRandomNumber(cellSize - 1)
		if s[rowIndex].GetNumber(colIndex) != 0 {
			continue
		}

		randNum = s.genUniqueRandomNumber(rowIndex, colIndex)
		s[rowIndex][colIndex] = randNum
		numbersToFill--
	}
}

func (s Board) Copy() Board {
	mySudoku := make(Board, 0)
	done := make(chan struct{})

	go func() {
		for _, _Row := range s {
			myRow := make(Row, 0)
			for _, _col := range _Row {
				myRow = append(myRow, _col)
			}
			mySudoku = append(mySudoku, myRow)
		}
		done <- struct{}{}
	}()
	<-done
	return mySudoku
}

func (s Board) getRow(rowID int) Row {
	return s[rowID]
}

func (s Board) getColumn(colID int) Row {
	var myColumn Row

	for _, Row := range s {
		for colIndex, col := range Row {
			if colID == colIndex {
				myColumn = append(myColumn, col)
			}
		}
	}
	return myColumn
}

func (s Board) Print() {
	for _, col := range s {
		fmt.Println(col)
	}
}

func (s Board) getBoundedBox(rowID int, colID int) Row {
	var myBB Row
	RowMin, RowMax, colMin, colMax := getBoundedBoxBoundaries(rowID, colID)

	for RowIndex, Row := range s {
		for colIndex, col := range Row {
			if (RowIndex >= RowMin && RowIndex <= RowMax) && (colIndex >= colMin && colIndex <= colMax) {
				myBB = append(myBB, col)
			}
		}
	}
	return myBB
}

func (s Board) Fill(rowID int, colID int, val int) {
	done := make(chan struct{})

	go func(s Board) {
		s[rowID][colID] = val
		done <- struct{}{}
	}(s)

	<-done
}

func (s Board) UnfilledCount() (unfilled int) {
	done := make(chan struct{})

	go func() {
		for _, Row := range s {
			for _, col := range Row {
				if col == 0 {
					unfilled++
				}
			}
		}
		done <- struct{}{}
	}()

	<-done
	return unfilled
}

func (s Board) genUniqueRandomNumber(rowID int, colID int) int {
	cellSize := s.CellSize()
	var randNum int

	iterationNum := 0
	for {
		iterationNum++
		randNum = genRandomNumber(cellSize)
		if !s.isPresent(randNum, rowID, colID) {
			break
		}
	}
	return randNum
}

func (s Board) isPresent(numToFill int, rowID int, colID int) bool {
	for rowInRex, Row := range s {
		for colIndex, colValue := range Row {

			if numToFill == colValue && rowInRex == rowID {
				return true
			}

			if numToFill == colValue && colIndex == colID {
				return true
			}

			switch {
			case colID <= 2:
				switch {
				case rowID <= 2:
					if isPresentBoundingBox(2, 2, rowInRex, colIndex, numToFill, colValue) {
						return true
					}
					break
				case rowID <= 5:
					if isPresentBoundingBox(5, 2, rowInRex, colIndex, numToFill, colValue) {
						return true
					}
					break
				case rowID <= 8:
					if isPresentBoundingBox(8, 2, rowInRex, colIndex, numToFill, colValue) {
						return true
					}
					break
				}
				break
			case colID <= 5:
				switch {
				case rowID <= 2:
					if isPresentBoundingBox(2, 5, rowInRex, colIndex, numToFill, colValue) {
						return true
					}
					break
				case rowID <= 5:
					if isPresentBoundingBox(5, 5, rowInRex, colIndex, numToFill, colValue) {
						return true
					}
					break
				case rowID <= 8:
					if isPresentBoundingBox(8, 5, rowInRex, colIndex, numToFill, colValue) {
						return true
					}
					break
				}
				break
			case colID <= 8:
				switch {
				case rowID <= 2:
					if isPresentBoundingBox(2, 8, rowInRex, colIndex, numToFill, colValue) {
						return true
					}
					break
				case rowID <= 5:
					if isPresentBoundingBox(5, 8, rowInRex, colIndex, numToFill, colValue) {
						return true
					}
					break
				case rowID <= 8:
					if isPresentBoundingBox(8, 8, rowInRex, colIndex, numToFill, colValue) {
						return true
					}
					break
				}
				break
			}
		}
	}
	return false
}
