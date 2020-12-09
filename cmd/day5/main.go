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
	fptr := flag.String("fpath", path.Join(os.Args[0], "../inputs/day5/input.txt"), "file path to read from")
	rptr := flag.Int("rows", 128, "number of seat rows on the aircraft")
	cptr := flag.Int("columns", 8, "number of seat columns on the aircraft")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	passes, err := aoc.NewBoardingPassList(f, *rptr, *cptr)
	if err != nil {
		log.Fatal(err)
	}
	sort.Slice(passes, func(i, j int) bool {
		return passes[i].ID < passes[j].ID
	})
	fmt.Println("Highest Seat ID:", passes[len(passes)-1].ID)

	for i, pass := range passes {
		if (i+1) < len(passes) && passes[i+1].ID == (pass.ID+2) {
			fmt.Println("Missing Seat ID:", (pass.Row*8)+pass.Column+1)
			return
		}
	}
}
