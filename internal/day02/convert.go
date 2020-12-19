package day02

import (
	"aoc-2020/pkg/convert"
	"regexp"
)

func ConvertToPolicy(str *string) Policy {
	re := regexp.MustCompile(`(\d+)-(\d+)\s([a-z]):\s(\w+)`)
	match := re.FindStringSubmatch(*str)
	return Policy{
		From: convert.ToInt(match[1]),
		To: convert.ToInt(match[2]),
		Letter: match[3],
		Password: match[4],
	}
}

func ConvertArrToPolicy(strArr *[]string) []Policy {
	policies := make([]Policy, 0)
	for _, val := range(*strArr) {
		policies = append(policies, ConvertToPolicy(&val))
	}
	return policies
}