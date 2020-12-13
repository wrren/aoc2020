package aoc

import (
	"bufio"
	"fmt"
	"io"
)

type ConsoleInstruction struct {
	Instruction string
	Data        int
}

type ReplacementInstruction struct {
	Index    int
	Original ConsoleInstruction
}

type Console struct {
	Accumulator int
	Code        []ConsoleInstruction
	Replacement *ReplacementInstruction
}

func NewConsole(reader io.Reader) (Console, error) {
	scanner := bufio.NewScanner(reader)
	var instruction string
	var data int
	console := Console{Code: make([]ConsoleInstruction, 0)}

	for scanner.Scan() {
		line := scanner.Text()
		parsed, _ := fmt.Sscanf(line, "%s %d", &instruction, &data)
		if parsed == 2 {
			console.Code = append(console.Code, ConsoleInstruction{Instruction: instruction, Data: data})
		}
	}

	return console, nil
}

func (i ConsoleInstruction) Execute(c *Console, p int) int {
	switch i.Instruction {
	case "nop":
		break
	case "acc":
		c.Accumulator += i.Data
		break
	case "jmp":
		return p + i.Data
	}
	return p + 1
}

func (c *Console) Reset() {
	c.Accumulator = 0
	if c.Replacement != nil {
		c.Code[c.Replacement.Index] = c.Replacement.Original
	}
}

func (c *Console) Run() bool {
	p := 0
	h := map[int]bool{}
	for p < len(c.Code) {
		_, ok := h[p]
		if ok {
			return false
		}
		i := c.Code[p]
		h[p] = true
		p = i.Execute(c, p)
	}
	return true
}

func (c *Console) ReplaceNextInstruction() {
	index := 0
	if c.Replacement != nil {
		index = c.Replacement.Index + 1
	} else {
		index = 0
	}

	for index < len(c.Code) {
		i := c.Code[index]
		if i.Instruction == "nop" {
			c.Replacement = &ReplacementInstruction{
				Index:    index,
				Original: i,
			}
			c.Code[index].Instruction = "jmp"
			return
		} else if i.Instruction == "jmp" {
			c.Replacement = &ReplacementInstruction{
				Index:    index,
				Original: i,
			}
			c.Code[index].Instruction = "nop"
			return
		}
		index++
	}
}
