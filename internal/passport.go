package aoc

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var optionalFields = []string{"cid"}

type Passport struct {
	fields map[string]string
}

func NewPassportList(reader io.Reader) ([]Passport, error) {
	passports := make([]Passport, 0)
	fields := map[string]string{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		if len(line) == 0 {
			passports = append(passports, Passport{fields: fields})
			fields = map[string]string{}
		} else {
			tokens := strings.Split(line, " ")
			for _, token := range tokens {
				kv := strings.Split(token, ":")
				if len(kv) == 2 {
					fields[kv[0]] = kv[1]
				}
			}
		}
	}
	if len(fields) > 0 {
		passports = append(passports, Passport{fields: fields})
	}
	return passports, nil
}

func (p Passport) IsValid() bool {
	for _, field := range requiredFields {
		_, ok := p.fields[field]
		if !ok {
			return false
		}
	}
	return true
}

func (p Passport) isBirthYearValid() bool {
	if p.IsValid() {
		byr, err := strconv.Atoi(p.fields["byr"])
		return len(p.fields["byr"]) == 4 && err == nil && byr >= 1920 && byr <= 2002
	}
	return false
}

func (p Passport) isIssueYearValid() bool {
	if p.IsValid() {
		iyr, err := strconv.Atoi(p.fields["iyr"])
		return len(p.fields["iyr"]) == 4 && err == nil && iyr >= 2010 && iyr <= 2020
	}
	return false
}

func (p Passport) isExpiryYearValid() bool {
	if p.IsValid() {
		eyr, err := strconv.Atoi(p.fields["eyr"])
		return len(p.fields["eyr"]) == 4 && err == nil && eyr >= 2020 && eyr <= 2030
	}
	return false
}

func (p Passport) isHeightValid() bool {
	if !p.IsValid() {
		return false
	}
	if len(p.fields["hgt"]) < 3 {
		return false
	}
	suffix := p.fields["hgt"][len(p.fields["hgt"])-2:]
	if suffix != "cm" && suffix != "in" {
		return false
	}
	hgt, err := strconv.Atoi(p.fields["hgt"][:len(p.fields["hgt"])-2])
	if err != nil || (suffix == "cm" && (hgt < 150 || hgt > 193)) || (suffix == "in" && (hgt < 59 || hgt > 76)) {
		return false
	}
	return true
}

func (p Passport) isHairColourValid() bool {
	if !p.IsValid() {
		return false
	}
	cre := regexp.MustCompile(`^#[a-f0-9]{6}$`)
	return cre.Match([]byte(p.fields["hcl"]))
}

func (p Passport) isEyeColourValid() bool {
	if !p.IsValid() {
		return false
	}
	return p.fields["ecl"] == "amb" || p.fields["ecl"] == "blu" || p.fields["ecl"] == "brn" || p.fields["ecl"] == "gry" || p.fields["ecl"] == "grn" || p.fields["ecl"] == "hzl" || p.fields["ecl"] == "oth"
}

func (p Passport) isPassportIDValid() bool {
	if !p.IsValid() {
		return false
	}
	pre := regexp.MustCompile(`^[0-9]{9}$`)
	return pre.Match([]byte(p.fields["pid"]))
}

func (p Passport) IsStrictlyValid() bool {
	if p.IsValid() && p.isBirthYearValid() && p.isIssueYearValid() && p.isExpiryYearValid() && p.isHeightValid() && p.isHairColourValid() && p.isEyeColourValid() && p.isPassportIDValid() {
		return true
	}
	return false
}
