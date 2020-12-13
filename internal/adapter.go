package aoc

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"
)

type Adapter int

func NewAdapterList(reader io.Reader) ([]Adapter, error) {
	scanner := bufio.NewScanner(reader)
	adapters := make([]Adapter, 0)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		if len(line) == 0 {
			break
		}
		joltage, err := strconv.Atoi(line)
		if err != nil {
			return adapters, err
		}
		adapters = append(adapters, Adapter(joltage))
	}
	return adapters, nil
}

func CountAdapterJoltageDifferences(adapters []Adapter) map[int]int {
	differences := map[int]int{}
	sort.Slice(adapters, func(i, j int) bool {
		return adapters[i] < adapters[j]
	})

	last := 0
	for _, a := range adapters {
		differences[int(a)-last]++
		last = int(a)
	}
	differences[3]++

	return differences
}

func CountValidAdapterCombinations(adapters []Adapter) uint64 {
	adapters = append(adapters, 0)
	sort.Slice(adapters, func(i, j int) bool {
		return adapters[i] < adapters[j]
	})
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	paths := map[int][]int{}
	for i, ai := range adapters {
		if i == len(adapters)-1 {
			break
		}
		for j, aj := range adapters[i+1:] {
			index := i + 1 + j
			if aj-ai <= 3 {
				_, ok := paths[index]
				if !ok {
					paths[index] = make([]int, 0, 1)
				}
				paths[index] = append(paths[index], i)
			} else {
				break
			}
		}
	}

	var combinations uint64 = 0
	memo := map[int]uint64{}
	recursivelyCountPaths(paths, len(adapters)-1, &combinations, &memo)
	return combinations
}

func recursivelyCountPaths(paths map[int][]int, index int, combinations *uint64, memo *map[int]uint64) {
	c, ok := (*memo)[index]
	if ok {
		*combinations += c
		return
	}
	if index == 0 {
		*combinations++
		return
	}
	c = *combinations
	for _, p := range paths[index] {
		recursivelyCountPaths(paths, p, combinations, memo)
	}
	(*memo)[index] = *combinations - c
}
