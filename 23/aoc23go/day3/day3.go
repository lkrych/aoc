package day3

import (
	"flag"
	"fmt"
	"log"
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
	if i >= 0 && i < len {
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

// findAllIndices finds all indices of a character in a string.
func findAllIndices(s string, char rune) []int {
	var indices []int
	for i, c := range s {
		if c == char {
			indices = append(indices, i)
		}
	}
	return indices
}

func getEnginePartIndicesPerLine(l string) (matches [][]int, err error) {
	pattern := `\d+`
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}
	matches = re.FindAllStringIndex(l, -1)
	return
}

func getEnginePartFromMatch(m []int, l string) (enginePart int) {
	// if the enginePart is next to a symbol, save it to the map!
	numberStr := l[m[0]:m[1]]
	enginePart, _ = strconv.Atoi(numberStr)
	return
}

func checkForEngineParts(gIdx int, lIdx int, lines []string) (engineParts []int) {
	// so the first thing we will do is check to see if there are neighboring engine parts on the same line
	sameLineMatches, err := getEnginePartIndicesPerLine((lines[lIdx]))
	if err != nil {
		log.Fatal("Regex failure")
	}
	// we are looking for numbers directly to the left or right of a gear
	for _, match := range sameLineMatches {
		// check to the left of gear
		// 765*
		if gIdx == match[1] {
			engineParts = append(engineParts, getEnginePartFromMatch(match, lines[lIdx]))
		}
		// check to the right of gear
		//*765
		if gIdx == match[0]-1 {
			engineParts = append(engineParts, getEnginePartFromMatch(match, lines[lIdx]))
		}
	}

	// next we need to scan the line above
	if indexInRange(lIdx-1, len(lines)) {
		prevLine := lines[lIdx-1]
		prevLineMatches, err := getEnginePartIndicesPerLine((prevLine))
		if err != nil {
			log.Fatal("Regex failure")
		}
		for _, match := range prevLineMatches {
			enginePart := getEnginePartFromMatch(match, prevLine)
			// fmt.Println("prevLineMatches: ")
			// fmt.Println("   Checking ", enginePart, "with match ", match)
			// check to the left of gear
			// 765.
			// ...*
			if gIdx == match[1] {
				engineParts = append(engineParts, enginePart)
			}
			// check to the right of gear
			// .765
			// *...
			if gIdx == match[0]-1 {
				engineParts = append(engineParts, enginePart)
			}
			// check if the number is directly above
			// 765
			// .*..
			for overlapIdx := match[0]; overlapIdx < match[1]; overlapIdx++ {
				if gIdx == overlapIdx {
					engineParts = append(engineParts, enginePart)
				}
			}
		}
	}

	// then we need to scan the line below
	if indexInRange(lIdx+1, len(lines)) {
		nextLine := lines[lIdx+1]
		nextLineMatches, err := getEnginePartIndicesPerLine((nextLine))
		if err != nil {
			log.Fatal("Regex failure")
		}
		for _, match := range nextLineMatches {
			enginePart := getEnginePartFromMatch(match, nextLine)
			// fmt.Println("nextLineMatches: ")
			// fmt.Println("   Checking ", enginePart, "with match ", match)
			// check to the left of gear
			// ...*
			// 765.
			if gIdx == match[1] {
				engineParts = append(engineParts, enginePart)
			}
			// check to the right of gear
			// *...
			// .765
			if gIdx == match[0]-1 {
				engineParts = append(engineParts, enginePart)
			}
			// check if the number is directly below
			// .*..
			// 765
			for overlapIdx := match[0]; overlapIdx < match[1]; overlapIdx++ {
				if gIdx == overlapIdx {
					engineParts = append(engineParts, enginePart)
				}
			}
		}
	}

	return engineParts
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

	// read into memory, this could be optimized by only reading the adjacent lines to the ones we are interested in
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	gearRatio := 0

	for i, line := range lines {
		fmt.Println(line)
		// Find all engine parts and their indices in the string
		gearIndices := findAllIndices(line, '*')

		for _, gIdx := range gearIndices {
			// fmt.Println("Checking gearIdx ", gIdx)
			engineParts := checkForEngineParts(gIdx, i, lines)
			// fmt.Println("Found engineParts: ", engineParts)
			if len(engineParts) == 2 {
				fmt.Println("found a gear: ", gIdx, " and subsequent engine parts ", engineParts)
				gearRatio += engineParts[0] * engineParts[1]
			}
		}
	}

	// BEGIN CODING FOR DAY HERE
	// INITIALIZE GLOBAL VALUES

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	fmt.Println("Gear ratio: ", gearRatio)
}
