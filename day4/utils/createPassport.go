package utils

import (
	"day4/model"
	"log"
	"strconv"
	"strings"
)

func convertToInt(text string) int {
	val, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal(err)
	}
	return val
}


func CreatePassport(str string) (model.Passport, bool) {
	parts := strings.Split(str, " ")
	m := make(map[string]string)
	for _, part := range parts {
		val := strings.Split(part, ":")
		m[val[0]] = val[1]
	}

	if !(isFourDigitNumber(m["byr"]) &&
		 isFourDigitNumber(m["iyr"]) &&
		 isFourDigitNumber(m["eyr"])) {
		return model.Passport{}, false
	}
	passport := model.Passport{
		Byr: convertToInt(m["byr"]),
		Iyr: convertToInt(m["iyr"]),
		Eyr: convertToInt(m["eyr"]),
		Hgt: m["hgt"],
		Hcl: m["hcl"],
		Ecl: m["ecl"],
		Pid: m["pid"],
	}

	if(passport.IsValid()) {
		return passport, true
	}

	return model.Passport{}, false
}