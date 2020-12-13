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
	fptr := flag.String("fpath", path.Join(os.Args[0], "../inputs/day9/input.txt"), "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	xmas, err := aoc.NewXMAS(f)
	if err != nil {
		log.Fatal(err)
	}

	invalid, found := xmas.FindInvalidCode()
	if !found {
		fmt.Println("Failed to find any invalid codes.")
		return
	} else {
		fmt.Println("Found invalid code: ", invalid)
	}
	weakness, found := xmas.FindWeakness(invalid)
	if !found {
		fmt.Println("Failed to find weakness.")
		return
	} else {
		fmt.Println("Found weakness: ", weakness)
	}
}
