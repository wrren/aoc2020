package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	aoc "github.com/wrren/aoc2020/internal"
)

func main() {
	fptr := flag.String("fpath", path.Join(os.Args[0], "../inputs/day10/input.txt"), "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	adapters, err := aoc.NewAdapterList(f)
	if err != nil {
		log.Fatal(err)
	}
	differences := aoc.CountAdapterJoltageDifferences(adapters)
	fmt.Println("Differences:", differences[1]*differences[3])
	fmt.Println("Valid Combinations:", aoc.CountValidAdapterCombinations(adapters))
}
