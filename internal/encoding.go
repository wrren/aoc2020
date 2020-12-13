package aoc

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"
)

type XMAS struct {
	Codes []int
}

func NewXMAS(reader io.Reader) (XMAS, error) {
	scanner := bufio.NewScanner(reader)
	xmas := XMAS{Codes: make([]int, 0)}
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.Trim(line, " ")) == 0 {
			break
		}
		code, err := strconv.Atoi(line)
		if err != nil {
			return xmas, err
		}
		xmas.Codes = append(xmas.Codes, code)
	}
	return xmas, nil
}

func isValidCode(covered []int, code int) bool {
	for _, c := range covered {
		diff := code - c
		i := sort.SearchInts(covered, diff)
		if i != len(covered) && covered[i] == diff {
			return true
		}
	}
	return false
}

func (x XMAS) FindInvalidCode() (int, bool) {
	covered := make([]int, 0, len(x.Codes))
	for i := range x.Codes {
		if i < 25 {
			covered = append(covered, x.Codes[i])
			sort.Ints(covered)
			continue
		} else {
			if !isValidCode(covered, x.Codes[i]) {
				return x.Codes[i], true
			}
			covered = append(covered, x.Codes[i])
			sort.Ints(covered)
		}
	}
	return 0, false
}

func (x XMAS) FindWeakness(invalid int) (int, bool) {
	window := make([]int, 0)
	start := 0
	end := 0
	sum := 0

	for start < len(x.Codes) {
		if sum == invalid {
			sort.Ints(window)
			return window[0] + window[len(window)-1], true
		}
		if sum < invalid {
			sum += x.Codes[end]
			window = append(window, x.Codes[end])
			end++
		} else {
			sum -= window[0]
			window = window[1:]
			for sum > invalid {
				end--
				sum -= window[len(window)-1]
				window = window[0 : len(window)-1]
			}
		}
	}
	return 0, false
}
