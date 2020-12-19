package day04

import (
	"aoc-2020/pkg/convert"
	"strconv"
	"strings"
)

var requiredFields = []string{
	"byr", "iyr", "eyr",
	"hgt", "hcl", "ecl", 
	"pid",
}

type validatePass func(*string) bool

func isFourDigitNumber(text string) bool {
	if len(text) != 4 {
		return false
	}
	if _, err := strconv.Atoi(text); err != nil {
		return false
	}
	return true
}

func ContainsWords(text *string) bool {
	for _, word := range(requiredFields) {
		if ! strings.Contains(*text, word) {
			return false
		} 
	}
	return true
}

func CreatePassport(str *string) (Passport, bool) {
	parts := strings.Split(*str, " ")
	m := make(map[string]string)
	for _, part := range parts {
		val := strings.Split(part, ":")
		m[val[0]] = val[1]
	}

	if !(isFourDigitNumber(m["byr"]) &&
		 isFourDigitNumber(m["iyr"]) &&
		 isFourDigitNumber(m["eyr"])) {
		return Passport{}, false
	}
	passport := Passport{
		Byr: convert.ToInt(m["byr"]),
		Iyr: convert.ToInt(m["iyr"]),
		Eyr: convert.ToInt(m["eyr"]),
		Hgt: m["hgt"],
		Hcl: m["hcl"],
		Ecl: m["ecl"],
		Pid: m["pid"],
	}

	if(passport.IsValid()) {
		return passport, true
	}

	return Passport{}, false
}

func IsValidPass(passStr *string) bool {
	_, isvalid := CreatePassport(passStr)
	return isvalid
}

func CountValidPass(checkPassFunc validatePass, strArr *[]string) int {
	validPass := 0
	for _, val := range(*strArr) {
		if checkPassFunc(&val) {
			validPass ++
		}
	}
	return validPass
}