package day04

import (
	"log"
	"regexp"
	"strconv"
)

const MinBirthYear = 1920
const MaxBirthYear = 2002

const MinIssueYear = 2010
const MaxIssueYear = 2020

const MinExpirationYear = 2020
const MaxExpirationYear = 2030

const MinHeightCm = 150
const MaxHeightCm = 193
const MinHeightIn = 59
const MaxHeightIn = 76

var EyeColors = []string{
	"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

type Passport struct {
	Byr int
	Iyr int
	Eyr int
	Hgt string
	Hcl string
	Ecl string
	Pid string
}

func (p Passport) GetHeight() (int, string) {
	numberString := p.Hgt[0:len(p.Hgt)-2]
	unitString := p.Hgt[len(p.Hgt)-2:]
	number, err := strconv.Atoi(numberString)
	if err != nil {
		log.Fatal(err)
	}
	return number, unitString
}

func (p Passport) hairColorIsValid() bool {
	re := regexp.MustCompile(`^#(?:\d|[a-f]){6}$`)
	return re.MatchString(p.Hcl)
}

func (p Passport) isHeightValid() bool {
	re := regexp.MustCompile(`^(?:\d{1,}(cm|in))$`)
	return re.MatchString(p.Hgt)
}

func (p Passport) isPassportIdValid() bool {
	re := regexp.MustCompile(`^(?:\d{9})$`)
	return re.MatchString(p.Pid)
}

func contains(sArr []string, s string) bool {
	for _, str := range sArr {
		if str == s {
			return true
		}
	}
	return false
}

func (p Passport) IsValid() bool {

	if !(p.Byr >= MinBirthYear && 
		 p.Byr <= MaxBirthYear) {
		return false
	}
	if !(p.Iyr >= MinIssueYear && 
		 p.Iyr <= MaxIssueYear) {
		return false
	}
	if !(p.Eyr >= MinExpirationYear && 
		 p.Eyr <= MaxExpirationYear) {
		return false
	}

	if(!p.isHeightValid()) {
		return false
	}

	height, unit := p.GetHeight()
	if unit == "cm" {
		if !(height >= MinHeightCm &&
			 height <= MaxHeightCm) {
			return false
		}
	} else if unit == "in" {
		if !(height >= MinHeightIn &&
			height <= MaxHeightIn) {
		   return false
	   }
	}

	if !p.hairColorIsValid() {
		return false
	}

	if !contains(EyeColors, p.Ecl) {
		return false
	}

	if !p.isPassportIdValid() {
		return false
	}

	return true
}