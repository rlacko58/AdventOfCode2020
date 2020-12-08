package main

import (
	"day4/utils"
	"fmt"
)

var requiredFields = []string{
	"byr", "iyr", "eyr",
	"hgt", "hcl", "ecl", 
	"pid",
}

func task1() {
	passports := utils.GetStrings()
	validPasswords := 0
	for _, pass := range passports {
		if utils.ContainsWords(pass, requiredFields) {
			validPasswords ++
		}
	}
	fmt.Println("Valid passwords:", validPasswords)
}

func task2() {
	passports := utils.GetStrings()
	validPasswords := 0
	for _, pass := range passports {
		_, isvalid := utils.CreatePassport(pass)
		if(isvalid) {
			validPasswords ++
		}
	}
	fmt.Println("Valid passwords:", validPasswords)
}

func main() {
	// task1()
	task2()
}