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
	fptr := flag.String("fpath", path.Join(os.Args[0], "../inputs/day4/input.txt"), "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	list, err := aoc.NewPassportList(f)
	if err != nil {
		log.Fatal(err)
	}

	valid := 0
	strictlyValid := 0
	for _, passport := range list {
		if passport.IsValid() {
			valid++
		}
		if passport.IsStrictlyValid() {
			strictlyValid++
		}
	}
	fmt.Println("Valid Passports:", valid, "/", len(list))
	fmt.Println("Strictly Valid Passports:", strictlyValid, "/", len(list))
}
