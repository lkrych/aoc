package main

import (
	"fmt"
	"regexp"
	"strconv"
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
			if isPassportValid2(currentPassport) {
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
		if isPassportValid2(currentPassport) {
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

func isPassportValid2(passport map[string]string) bool {
	validFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	valid := true
	for _, field := range validFields {
		if val, ok := passport[field]; ok {
			switch field {
			case "byr":
				valid = checkIntVal(val, 1920, 2002)
			case "iyr":
				valid = checkIntVal(val, 2010, 2020)
			case "eyr":
				valid = checkIntVal(val, 2020, 2030)
			case "hgt":
				valid = checkHgt(val)
			case "hcl":
				valid = checkHcl(val)
			case "ecl":
				valid = checkEcl(val)
			case "pid":
				valid = checkPid(val)
			default:
			}
		} else {
			// key isn't in the passport
			valid = false
		}
		if !valid {
			// if we found an invalid value, just break the loop
			break
		}
	}
	if valid {
		fmt.Println("valid: ", passport)
	}
	return valid
}

func checkIntVal(val string, low int, high int) bool {
	valid := true
	intval, err := strconv.Atoi(val)
	if err != nil {
		valid = false
	}
	if intval < low || intval > high {
		valid = false
	}
	return valid
}

func checkHgt(val string) bool {
	valid := true
	if strings.Contains(val, "in") {
		idx := strings.Index(val, "in")
		hgt := val[:idx]
		valid = checkIntVal(hgt, 59, 76)
	} else if strings.Contains(val, "cm") {
		idx := strings.Index(val, "cm")
		hgt := val[:idx]
		valid = checkIntVal(hgt, 150, 193)
	} else {
		valid = false
	}
	return valid
}

func checkHcl(val string) bool {
	match, _ := regexp.MatchString("#([a-f0-9])", val)
	return match
}

func checkEcl(val string) bool {
	valid := false
	validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, col := range validColors {
		if val == col {
			valid = true
		}
	}
	return valid
}

func checkPid(val string) bool {
	valid := true
	if len(val) != 9 {
		valid = false
	}
	_, err := strconv.Atoi(val)
	if err != nil {
		valid = false
	}
	return valid
}
