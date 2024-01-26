package day6

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/lkrych/aoc23go/input"
)

func convertStringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}
	return i
}

func findWinningRanges(timeToRace int, maxDistance int) int {
	winningRanges := 0
	for i := 0; i < timeToRace; i++ {
		distance := findDistance(i, timeToRace)
		if distance > maxDistance {
			fmt.Println(distance, " is a winning range for pressing the button ", i, " for a total race time of ", timeToRace, "and a max dist of ", maxDistance)
			winningRanges += 1
		}
	}
	return winningRanges
}

func findDistance(timeToPressButton int, timeToRace int) int {
	totalSpeed := 1 * timeToPressButton
	totalTimeToRace := timeToRace - timeToPressButton

	return totalSpeed * totalTimeToRace
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

	// BEGIN CODING FOR DAY HERE
	// INITIALIZE GLOBAL VALUES

	// read from input
	times := []int{}
	distances := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		for _, el := range split[1:] {
			if el == "" {
				continue
			}
			intVal := convertStringToInt(el)
			if strings.Contains(line, "Time:") {
				times = append(times, intVal)
			} else if strings.Contains(line, "Distance:") {
				distances = append(distances, intVal)
			}
		}
	}

	fmt.Println(times)
	fmt.Println(distances)
	maxRanges := 1
	// iterate through races and determine the ran
	for i, t := range times {
		d := distances[i]
		maxRanges *= findWinningRanges(t, d)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println("Winning range multiple ", maxRanges)
}
