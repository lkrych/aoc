package main

import (
	"fmt"
	"strings"
)

func main() {
	data, err := ReadFile("../input/aoc4.txt")
	if err != nil {
		fmt.Printf("readFile failed: %v\n", err)
	}

	ans := validPassports(data)

	fmt.Printf("The number of valid passports is %d \n", ans)
}

func validPassports(data []string) int {
	valid := 0
	currentPassport := make(map[string]string)
	for _, line := range data {
		if len(line) <= 1 {
			// we are at an empty line and we need to clear the current passport
			if isPassportValid(currentPassport) {
				valid++
			}
			currentPassport = make(map[string]string)
		} else {
			splitBySpace := strings.Split(line, " ")
			for _, kv := range splitBySpace {
				splitByColon := strings.Split(kv, ":")
				currentPassport[splitByColon[0]] = splitByColon[1]
			}
		}
	}
	if len(currentPassport) > 1 {
		if isPassportValid(currentPassport) {
			valid++
		}
	}
	return valid
}

func isPassportValid(passport map[string]string) bool {
	validFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	valid := true
	for _, field := range validFields {
		if _, ok := passport[field]; !ok {
			valid = false
		}
	}
	return valid
}
