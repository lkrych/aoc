package day3

import (
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/lkrych/aoc23go/input"
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

	//
	symbols := []rune{"#", "&", "*", "+", "$", "/", "%", "-", "@", "="}

	// keep track of the matches we've already found, we don't want to double count
	var foundEngineParts map[int]bool

	for i, line := range lines {
		// Find all engine parts and their indices in the string
		matches := re.FindAllStringIndex(text, -1)

		// Print found numbers and their indices
		for _, match := range matches {
			startIdx, endIdx := match[0], match[1]
			// now we need to check positionally whether this int should be included in the foundEnginePartsMap
			// we do this by first testing the characters immediately to the right and left of this found int
			// then we will test all the spaces above and then all the spaces below the word

		}
			
		}
	}

	// BEGIN CODING FOR DAY HERE
	// INITIALIZE GLOBAL VALUES
	

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	fmt.Println("Possible Games Sum: ", possibleGamesSum)
}
