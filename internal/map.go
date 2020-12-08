package aoc

import (
	"bufio"
	"io"
)

type TreeMap struct {
	lines []string
}

func NewTreeMap(reader io.Reader) (TreeMap, error) {
	scanner := bufio.NewScanner(reader)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return TreeMap{lines: lines}, nil
}

func (t TreeMap) CountTrees(column, row, right, down int) int {
	trees := 0
	for row < len(t.lines) {
		if t.lines[row][column] == '#' {
			trees++
		}
		column = (column + right) % len(t.lines[row])
		row = row + down
	}

	return trees
}
