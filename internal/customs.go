package aoc

import (
	"bufio"
	"io"
	"strings"
)

type CustomsForm struct {
	answers []rune
}

type CustomsFormGroup struct {
	forms []CustomsForm
}

func NewCustomsForm(answers []rune) (CustomsForm, error) {
	return CustomsForm{answers: answers}, nil
}

func NewCustomsFormGroup(lines []string) (CustomsFormGroup, error) {
	forms := make([]CustomsForm, 0)

	for _, line := range lines {
		form, err := NewCustomsForm([]rune(line))
		if err != nil {
			return CustomsFormGroup{}, err
		}
		forms = append(forms, form)
	}

	return CustomsFormGroup{forms: forms}, nil
}

func NewCustomsFormGroupList(reader io.Reader) ([]CustomsFormGroup, error) {
	scanner := bufio.NewScanner(reader)
	lines := make([]string, 0)
	groups := make([]CustomsFormGroup, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.Trim(line, " ")) == 0 {
			group, err := NewCustomsFormGroup(lines)
			if err != nil {
				return groups, err
			}
			lines = make([]string, 0)
			groups = append(groups, group)
		} else {
			lines = append(lines, line)
		}
	}
	if len(lines) > 0 {
		group, err := NewCustomsFormGroup(lines)
		if err != nil {
			return groups, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func (g CustomsFormGroup) CountUniqueAnswers() int {
	answers := map[rune]bool{}

	for _, f := range g.forms {
		for _, r := range f.answers {
			answers[r] = true
		}
	}

	return len(answers)
}

func (g CustomsFormGroup) CountUnanimousAnswers() int {
	answers := map[rune]int{}

	for _, f := range g.forms {
		for _, r := range f.answers {
			answers[r]++
		}
	}
	unanimous := 0
	for _, v := range answers {
		if v == len(g.forms) {
			unanimous++
		}
	}
	return unanimous
}
