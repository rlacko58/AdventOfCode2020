package model

import (
	"log"
	"strconv"
	"strings"
)

type ContainsBag struct {
	Bag *Bag
	Number int
}

type Bag struct {
	Color string
	ParentBags []*Bag
	Contains []*ContainsBag
}

func GetBag(str *string, bags []*Bag) *Bag {
	for _, bag := range(bags) {
		if bag.Color == *str {
			return bag
		}
	}
	return nil
}

func CreateBags(str []string) []*Bag {
	// Create Bags
	bags := make([]*Bag, 0)
	for _, s := range(str) {
		parts := strings.Split(s, "bags contain")
		bag := new(Bag)
		bag.Color = strings.TrimSpace(parts[0])
		bags = append(bags, bag)
	}
	// Make connections
	for _, s := range(str) {
		parts := strings.Split(s, "bags contain")
		parts[0] = strings.TrimSpace(parts[0])
		bag := GetBag(&parts[0], bags)
		var contains []string

		if parts[1] != " no other bags." {
			// Trim bag strings
			trimmed := strings.ReplaceAll(parts[1], "bags", "")
			trimmed = strings.ReplaceAll(trimmed, "bag", "")
			contains = strings.Split(trimmed[:len(trimmed)-1], ",")
			// Trim spaces
			for i := range(contains) {
				contains[i] = strings.TrimSpace(contains[i])
			}
		}
		for _, v := range(contains) {
			parts := strings.SplitAfterN(v, " ", 2)
			number, err := strconv.Atoi(strings.TrimSpace(parts[0]))
			if err != nil {
				log.Fatal(err)
			}
			color := strings.TrimSpace(parts[1])
			containsBag := GetBag(&color, bags)
			thisContains := new(ContainsBag)
			thisContains.Number = number
			thisContains.Bag = containsBag
			bag.Contains = append(bag.Contains, thisContains)

			containsBag.ParentBags = append(containsBag.ParentBags, bag)
		}
	}

	return bags
}