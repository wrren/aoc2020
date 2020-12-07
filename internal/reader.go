package aoc

import (
	"bufio"
	"io"
	"strconv"
)

type IntegerReader struct {
	scanner *bufio.Scanner
}

func NewIntegerReader(reader io.Reader) IntegerReader {
	return IntegerReader{scanner: bufio.NewScanner(reader)}
}

func (i *IntegerReader) Integer() (int, error) {
	return strconv.Atoi(i.scanner.Text())
}

func (i *IntegerReader) Scan() bool {
	return i.scanner.Scan()
}
