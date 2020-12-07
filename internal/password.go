package aoc

import (
	"fmt"
)

type PasswordPolicy int

const (
	PasswordPolicyV1 PasswordPolicy = iota
	PasswordPolicyV2
)

type Password struct {
	min, max  int
	character byte
	password  string
}

func NewPassword(input string) (Password, error) {
	password := Password{}
	_, err := fmt.Sscanf(input, "%d-%d %c: %s", &password.min, &password.max, &password.character, &password.password)
	return password, err
}

func (p Password) IsValid(policy PasswordPolicy) bool {
	count := 0

	if policy == PasswordPolicyV1 {
		for i, _ := range p.password {
			if p.password[i] == p.character {
				count = count + 1
			}
		}
		return count >= p.min && count <= p.max
	}
	return (p.password[p.min-1] == p.character && p.password[p.max-1] != p.character) ||
		(p.password[p.min-1] != p.character && p.password[p.max-1] == p.character)
}
