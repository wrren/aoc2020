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
	fptr := flag.String("fpath", path.Join(os.Args[0], "../inputs/day6/input.txt"), "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	groups, err := aoc.NewCustomsFormGroupList(f)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, g := range groups {
		sum += g.CountUniqueAnswers()
	}
	fmt.Println("Unique Sum:", sum)

	sum = 0
	for _, g := range groups {
		sum += g.CountUnanimousAnswers()
	}
	fmt.Println("Unanimous Sum:", sum)
}
