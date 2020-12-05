package model

const rowNum, columnNum int = 128, 8

type SeatPos struct {
	Row, Column int
}

func (p *SeatPos) GetSeatId() int {
	return p.Row * 8 + p.Column
}

func getLog(min, max int) int {
	return min + (max-min)/2
}

func MakeSeatPos(posString string) *SeatPos {
	pos := new(SeatPos)
	minRow, maxRow := 0, rowNum
	minColumn, maxColumn := 0, columnNum

	for _, s := range(posString) {
		switch(string(s)) {
		case "F":
			maxRow = getLog(minRow, maxRow)
		case "B":
			minRow = getLog(minRow, maxRow)
		case "L":
			maxColumn = getLog(minColumn, maxColumn)
		case "R":
			minColumn = getLog(minColumn, maxColumn)
		}
	}

	pos.Row = minRow
	pos.Column = minColumn

	return pos
}