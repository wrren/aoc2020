package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"sort"

	aoc "github.com/wrren/aoc2020/internal"
)

func main() {
	fptr := flag.String("fpath", path.Join(os.Args[0], "../inputs/day1/input.txt"), "file path to read from")
	tptr := flag.Int("target", 2020, "Target sum")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := aoc.NewIntegerReader(f)
	entries := make([]int, 0)

	for reader.Scan() {
		i, err := reader.Integer()
		if err != nil {
			log.Fatal(err)
		}
		entries = append(entries, i)
	}

	sort.Ints(entries)

	m1, m2 := aoc.FindEntriesWithSum(entries, *tptr)

	if m1 != len(entries) && m2 != len(entries) {
		fmt.Println()
		fmt.Println(m1, m2, m1*m2)
	} else {
		log.Fatalln("Failed to find entries summing to", *tptr)
	}

	for _, entry := range entries {
		diff := *tptr - entry
		m1, m2 = aoc.FindEntriesWithSum(entries, diff)
		if m1 != len(entries) && m2 != len(entries) {
			fmt.Println(entry, m1, m2, entry*m1*m2)
			return
		}
	}

	log.Fatalln("Failed to find 3 entries summing to", *tptr)
}
