package aoc

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type ContainmentRule struct {
	BagColour string
	Quantity  int
}

type Bag struct {
	Colour string
	Rules  []ContainmentRule
}

func NewBag(colour string) Bag {
	return Bag{Colour: colour, Rules: make([]ContainmentRule, 0)}
}

func (b *Bag) AddRule(rule ContainmentRule) {
	b.Rules = append(b.Rules, rule)
}

func (b Bag) CanContain(colour string) bool {
	for _, rule := range b.Rules {
		if rule.BagColour == colour {
			return true
		}
	}
	return false
}

func CountCanContain(bags []Bag, colour string) int {
	matches := map[string]bool{}
	toCheck := []string{colour}

	for {
		if len(toCheck) == 0 {
			break
		}
		colour = toCheck[0]

		for _, bag := range bags {
			for _, rule := range bag.Rules {
				if rule.BagColour == colour {
					_, ok := matches[bag.Colour]
					if !ok {
						toCheck = append(toCheck, bag.Colour)
						matches[bag.Colour] = true
					}
				}
			}
		}

		toCheck = toCheck[1:]
	}

	return len(matches)
}

func CountMustContain(bags []Bag, colour string) int {
	return countMustContainRecursive(bags, colour, 1) - 1
}

func countMustContainRecursive(bags []Bag, colour string, multiplier int) int {
	sum := multiplier
	for _, bag := range bags {
		if bag.Colour == colour {
			for _, rule := range bag.Rules {
				sum += countMustContainRecursive(bags, rule.BagColour, multiplier*rule.Quantity)
			}
			break
		}
	}
	return sum
}

func NewBagList(reader io.Reader) ([]Bag, error) {
	scanner := bufio.NewScanner(reader)
	bags := make([]Bag, 0)
	bagRegex := regexp.MustCompile(`([a-z]+ [a-z]+) bags contain (.+)`)
	ruleRegex := regexp.MustCompile(`([0-9]+) ([a-z]+ [a-z]+) bag`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := bagRegex.FindStringSubmatch(line)
		if len(matches) != 3 {
			return bags, fmt.Errorf("Line '%s' does not match regex", line)
		}
		colour := matches[1]
		bag := NewBag(colour)
		rules := strings.Split(matches[2], ",")
		for _, ruleString := range rules {
			if strings.Trim(ruleString, " ") == "no other bags." {
				break
			}
			matches = ruleRegex.FindStringSubmatch(strings.Trim(ruleString, " "))
			if len(matches) != 3 {
				return bags, fmt.Errorf("Rule '%s' does not match regex", ruleString)
			}
			quantity, err := strconv.Atoi(matches[1])
			if err != nil {
				return bags, fmt.Errorf("Rule '%s' quantity is not a number", ruleString)
			}
			bag.AddRule(ContainmentRule{BagColour: matches[2], Quantity: quantity})
		}
		bags = append(bags, bag)
	}

	return bags, nil
}
