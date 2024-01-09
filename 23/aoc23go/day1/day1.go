package day1

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"

	input "github.com/lkrych/aoc23go/input" // relative import for input
)

func Part1() {
	// BOILERPLATE for getting file name from stdIn and reading line by line
	filename := flag.String("f", "", "input file")
	// Parse the command-line arguments to read the flag value
	flag.Parse()
	filepath := fmt.Sprintf("../input/%s", *filename)
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

	delimiters := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	// Create a map to store the indices of the found words
	wordIndices := make(map[string][]int)

	// Iterate through matches and record the indices
	for _, d := range delimiters {
		idx := strings.Index(s, d)
		if 
	}

	// Sort the words by their index in the string
	var result []string
	for word := range wordIndices {
		result = append(result, word)
	}
	sort.Slice(result, func(i, j int) bool {
		return wordIndices[result[i]] < wordIndices[result[j]]
	})

	fmt.Println("Results ", result)

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
	for i, m := range result {
		if newValue, found := wordToValue[m]; found {
			result[i] = newValue
		}
	}
	return result
}

func Part2() {
	// BOILERPLATE for getting file name from stdIn and reading line by line
	filename := flag.String("f", "", "input file")
	// Parse the command-line arguments to read the flag value
	flag.Parse()
	filepath := fmt.Sprintf("../input/%s", *filename)
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
