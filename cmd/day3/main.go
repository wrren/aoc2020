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
	fptr := flag.String("fpath", path.Join(os.Args[0], "../inputs/day3/input.txt"), "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	treeMap, err := aoc.NewTreeMap(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", treeMap.CountTrees(0, 0, 3, 1))
	fmt.Println("Part 2:", treeMap.CountTrees(0, 0, 1, 1)*treeMap.CountTrees(0, 0, 3, 1)*treeMap.CountTrees(0, 0, 5, 1)*treeMap.CountTrees(0, 0, 7, 1)*treeMap.CountTrees(0, 0, 1, 2))
}
