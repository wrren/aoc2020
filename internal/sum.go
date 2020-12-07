package aoc

import (
	"sort"
)

func FindEntriesWithSum(entries []int, sum int) (a, b int) {
	for _, entry := range entries {
		diff := sum - entry
		if diff < entries[0] || diff > entries[len(entries)-1] {
			continue
		}
		i := sort.SearchInts(entries, diff)
		if i != len(entries) && entries[i] == diff {
			return entry, entries[i]
		}
	}
	return len(entries), len(entries)
}
