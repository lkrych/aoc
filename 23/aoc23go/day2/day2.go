package day2

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	input "github.com/lkrych/aoc23go/input"
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
		// INITIALIZE PER INPUT LINE VALUES
		var gameNumber int
		var gameSection string
		// separate the input into different components
		// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		// First
		// Split the line by ':' to get the part before the colon, which contains the game number.
		parts := strings.Split(line, ":")
		if len(parts) >= 1 {
			// extract gameNumber
			gameNumberWithText := strings.TrimSpace(parts[0]) // Remove leading/trailing whitespace
			gameNumberSplit := strings.Split(gameNumberWithText, " ")
			gameNumber, err = strconv.Atoi(gameNumberSplit[1])
			if err != nil {
				panic(err)
			}
			// next split apart each individual round within the game
			gameSection = parts[1]
			setOfRounds := strings.Split(gameSection, ";")

			// then keep track of the max possible values of each set of cubes per round of games

		}

		// add this number to calibrationSum
		possibleGamesSum += gameNumber
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	fmt.Println("Possible Games Sum: ", possibleGamesSum)
}
