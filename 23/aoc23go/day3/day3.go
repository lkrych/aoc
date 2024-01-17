package day3

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"

	"github.com/lkrych/aoc23go/input"
)

func checkForSymbols(b byte) bool {
	bs := string(b)
	symbols := []string{"#", "&", "*", "+", "$", "/", "%", "-", "@", "="}
	for _, s := range symbols {
		if bs == s {
			return true
		}
	}
	return false
}

func indexInRange(i int, len int) bool {
	if i > 0 && i < len {
		return true
	}
	return false
}

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

	// read into memory, this could be optimized by only reading the adjacent lines to the ones we are interested in
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// set up regexp to search for digits that we will use for each line
	pattern := `\d+`
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	// keep track of the matches we've already found, we don't want to double count
	foundEngineParts := map[int]int{}

	for i, line := range lines {
		fmt.Println(line)
		// Find all engine parts and their indices in the string
		matches := re.FindAllStringIndex(line, -1)

		// Print found numbers and their indices
		for _, match := range matches {
			// fmt.Println("match ", match)
			foundEnginePart := false
			startIdx, endIdx := match[0], match[1]
			// now we need to check positionally whether this int should be included in the foundEnginePartsMap
			// we do this by first testing the characters immediately to the right and left of this found int
			// then we will test all the spaces above and then all the spaces below the word

			// first start with the easy cases of left and right of the match
			if indexInRange(startIdx-1, len(line)) {
				if checkForSymbols(line[startIdx-1]) {
					foundEnginePart = true
				}
			}

			if indexInRange(endIdx, len(line)) {
				if checkForSymbols(line[endIdx]) {
					foundEnginePart = true
				}
			}

			// next check the characters above the current line
			if indexInRange(i-1, len(lines)) {
				prevLine := lines[i-1]
				// verify this line exists
				for i := startIdx - 1; i <= endIdx; i++ {
					if indexInRange(i, len(prevLine)) {
						if checkForSymbols(prevLine[i]) {
							foundEnginePart = true
						}
					}
				}
			}

			// next check the character below the current line
			if indexInRange(i+1, len(lines)) {
				nextLine := lines[i+1]
				// verify this line exists
				for i := startIdx - 1; i <= endIdx; i++ {
					if indexInRange(i, len(nextLine)) {
						if checkForSymbols(nextLine[i]) {
							foundEnginePart = true
						}
					}
				}
			}

			// if the enginePart is next to a symbol, save it to the map!
			numberStr := line[startIdx:endIdx]
			enginePart, err := strconv.Atoi(numberStr)
			if foundEnginePart {
				// fmt.Println("adding: ", enginePart)
				if err != nil {
					fmt.Printf("Error converting '%s' to integer: %v\n", numberStr, err)
					continue
				}
				value, exists := foundEngineParts[enginePart]
				if exists {
					foundEngineParts[enginePart] = value + 1
				} else {
					foundEngineParts[enginePart] = 1
				}
			} else {
				fmt.Println("didn't match ", enginePart)
			}

		}
	}

	enginePartSum := 0
	for enginePart, count := range foundEngineParts {
		enginePartSum += (enginePart * count)
	}

	// BEGIN CODING FOR DAY HERE
	// INITIALIZE GLOBAL VALUES

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	fmt.Println("Engine part sum: ", enginePartSum)
}
