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
	fptr := flag.String("fpath", path.Join(os.Args[0], "../inputs/day7/input.txt"), "file path to read from")
	cptr := flag.String("colour", "shiny gold", "bag colour to check for")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	bags, err := aoc.NewBagList(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(aoc.CountCanContain(bags, *cptr))
	fmt.Println(aoc.CountMustContain(bags, *cptr))
}
