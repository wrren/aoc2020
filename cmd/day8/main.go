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
	fptr := flag.String("fpath", path.Join(os.Args[0], "../inputs/day8/input.txt"), "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	console, err := aoc.NewConsole(f)
	if err != nil {
		log.Fatal(err)
	}
	if !console.Run() {
		fmt.Println("Code looped indefinitely, ACC=", console.Accumulator)
	}
	for {
		console.Reset()
		console.ReplaceNextInstruction()
		if console.Run() {
			fmt.Println("Code did not loop indefinitely, ACC=", console.Accumulator)
			return
		}
	}
}
