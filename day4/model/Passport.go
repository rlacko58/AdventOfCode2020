package model

import (
	"day4/consts"
	"log"
	"regexp"
	"strconv"
)

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

	if !(p.Byr >= consts.MinBirthYear && 
		 p.Byr <= consts.MaxBirthYear) {
		return false
	}
	if !(p.Iyr >= consts.MinIssueYear && 
		 p.Iyr <= consts.MaxIssueYear) {
		return false
	}
	if !(p.Eyr >= consts.MinExpirationYear && 
		 p.Eyr <= consts.MaxExpirationYear) {
		return false
	}

	if(!p.isHeightValid()) {
		return false
	}

	height, unit := p.GetHeight()
	if unit == "cm" {
		if !(height >= consts.MinHeightCm &&
			 height <= consts.MaxHeightCm) {
			return false
		}
	} else if unit == "in" {
		if !(height >= consts.MinHeightIn &&
			height <= consts.MaxHeightIn) {
		   return false
	   }
	}

	if !p.hairColorIsValid() {
		return false
	}

	if !contains(consts.EyeColors, p.Ecl) {
		return false
	}

	if !p.isPassportIdValid() {
		return false
	}

	return true
}