package aoc

import (
	"bufio"
	"fmt"
	"io"
)

type BoardingPass struct {
	seatBSP         string
	Row, Column, ID int
}

func NewBoardingPass(seatBSP string, rowCount, colCount int) (BoardingPass, error) {
	rowStart := 0
	rowEnd := rowCount
	colStart := 0
	colEnd := colCount

	for _, e := range seatBSP {
		switch e {
		case 'F':
			rowEnd = ((rowEnd - rowStart) / 2) + rowStart
		case 'B':
			rowStart = ((rowEnd - rowStart) / 2) + rowStart
		case 'L':
			colEnd = ((colEnd - colStart) / 2) + colStart
		case 'R':
			colStart = ((colEnd - colStart) / 2) + colStart
		}
	}
	if rowEnd-rowStart > 1 {
		return BoardingPass{}, fmt.Errorf("Not enough BSP indices to find row")
	}
	if colEnd-colStart > 1 {
		return BoardingPass{}, fmt.Errorf("Not enough BSP indices to find column")
	}
	return BoardingPass{seatBSP: seatBSP, Row: rowStart, Column: colStart, ID: (rowStart * 8) + colStart}, nil
}

func NewBoardingPassList(reader io.Reader, rowCount, colCount int) ([]BoardingPass, error) {
	passes := make([]BoardingPass, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		pass, err := NewBoardingPass(scanner.Text(), rowCount, colCount)
		if err != nil {
			return passes, err
		}
		passes = append(passes, pass)
	}
	return passes, nil
}
