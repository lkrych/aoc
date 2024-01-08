package day1

import (
	"fmt"
	"strconv"
	"unicode"

	input "github.com/lkrych/aoc23go/input" // relative import for input
)

func Part1() {
	// BOILERPLATE for getting file name from stdIn and reading line by line
	filename := input.GetFileName()
	filepath := fmt.Sprintf("../input/%s", filename)
	scanner, err := input.ReadInputFile(filepath)
	if err != nil {
		panic(err)
	}
	defer scanner.Scan() // Close the file when done reading

	// BEGIN CODING FOR DAY HERE
	// INITIALIZE GLOBAL VALUES
	calibrationSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		// INITIALIZE PER INPUT LINE VALUES
		var firstDigit, lastDigit string
		// iterate over every char
		for _, c := range line {
			// check if it is a digit
			if unicode.IsDigit(c) {
				// 0 is zero value in Golang
				// might be a bug here if first value is zero
				if firstDigit == "" {
					firstDigit = string(c)
				}
				// no matter what, set last digit
				lastDigit = string(c)
			}
		}
		// once the line has been parsed, create the two digit number
		concatenatedStr := firstDigit + lastDigit
		result, err := strconv.Atoi(concatenatedStr)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
		// add this number to calibrationSum
		calibrationSum += result
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	fmt.Println("Total Calibration Sum: ", calibrationSum)
}

func splitBasedOnNumbers(s string) []string {
	// FIND SOME WAY OF MATCHING OVERLAPPING REGEX

	fmt.Println("matches ", matches)

	wordToValue := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	// Iterate over the list and replace words if found in the map
	for i, m := range matches {
		if newValue, found := wordToValue[m]; found {
			matches[i] = newValue
		}
	}
	return matches
}

func Part2() {
	// BOILERPLATE for getting file name from stdIn and reading line by line
	filename := input.GetFileName()
	filepath := fmt.Sprintf("../input/%s", filename)
	scanner, err := input.ReadInputFile(filepath)
	if err != nil {
		panic(err)
	}
	defer scanner.Scan() // Close the file when done reading

	// BEGIN CODING FOR DAY HERE
	// INITIALIZE GLOBAL VALUES
	calibrationSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		// INITIALIZE PER INPUT LINE VALUES
		var firstDigit, lastDigit string

		digits := splitBasedOnNumbers(line)
		firstDigit = digits[0]
		lastDigit = digits[len(digits)-1]

		// once the line has been parsed, create the two digit number
		concatenatedStr := firstDigit + lastDigit
		result, err := strconv.Atoi(concatenatedStr)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
		// add this number to calibrationSum
		calibrationSum += result
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	fmt.Println("Total Calibration Sum: ", calibrationSum)
}
