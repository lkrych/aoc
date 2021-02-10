package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data, err := ReadFile("../input/aoc2.txt")
	if err != nil {
		fmt.Printf("readFile failed: %v\n", err)
	}
	ans := validPasswords(data)
	fmt.Printf("There are %d valid passwords\n", ans)
}

func validPasswords(data []string) int {
	valid := 0
	// example line -> 1-3 a: abcde
	for _, line := range data {
		if len(line) < 1 {
			continue
		}
		split := strings.Split(line, " ")
		password := split[2]
		char := split[1][0]
		countRange := split[0]
		if isPasswordValid2(password, char, countRange) {
			valid++
		}
	}
	return valid
}

func isPasswordValid(pass string, targetChar byte, countRange string) bool {
	m := make(map[byte]int)
	bs := []byte(pass)
	// populate map count
	for _, chr := range bs {
		if val, ok := m[chr]; ok {
			m[chr] = val + 1
		} else {
			m[chr] = 1
		}
	}
	targetCount, _ := m[targetChar]
	splitRange := strings.Split(countRange, "-")
	low, err := strconv.Atoi(splitRange[0])
	if err != nil {
		fmt.Printf("There was an error converting the low range in %s \n", countRange)
	}
	high, err := strconv.Atoi(splitRange[1])
	if err != nil {
		fmt.Printf("There was an error converting the high range in %s \n", countRange)
	}

	if targetCount <= high && targetCount >= low {
		return true
	}
	return false

}

func isPasswordValid2(pass string, targetChar byte, countRange string) bool {
	splitRange := strings.Split(countRange, "-")
	low, err := strconv.Atoi(splitRange[0])
	if err != nil {
		fmt.Printf("There was an error converting the low range in %s \n", countRange)
	}
	high, err := strconv.Atoi(splitRange[1])
	if err != nil {
		fmt.Printf("There was an error converting the high range in %s \n", countRange)
	}
	// XOR
	if (pass[low-1] == targetChar) != (pass[high-1] == targetChar) {
		return true
	}
	return false

}
