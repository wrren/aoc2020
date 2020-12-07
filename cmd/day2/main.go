package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	aoc "github.com/wrren/aoc2020/internal"
)

func main() {
	fptr := flag.String("fpath", path.Join(os.Args[0], "../inputs/day2/input.txt"), "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	validV1 := 0
	validV2 := 0

	scanner := bufio.NewScanner(f)
	passwords := make([]aoc.Password, 0)
	for scanner.Scan() {
		password, err := aoc.NewPassword(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		passwords = append(passwords, password)
		if password.IsValid(aoc.PasswordPolicyV1) {
			validV1 = validV1 + 1
		}
		if password.IsValid(aoc.PasswordPolicyV2) {
			validV2 = validV2 + 1
		}
	}

	fmt.Println("Policy V1:", validV1, "Valid Passwords of", len(passwords), "Passwords")
	fmt.Println("Policy V2:", validV2, "Valid Passwords of", len(passwords), "Passwords")
}
