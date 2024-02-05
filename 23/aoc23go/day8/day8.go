package day8

import (
	"flag"
	"fmt"

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

	// BEGIN CODING FOR DAY HERE
	// INITIALIZE GLOBAL VALUES
	possibleGamesSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	fmt.Println("Possible Games Sum: ", possibleGamesSum)
}
