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
			// fmt.Println(distance, " is a winning range for pressing the button ", i, " for a total race time of ", timeToRace, "and a max dist of ", maxDistance)
			winningRanges += 1
		}
	}
	return winningRanges
}

// attempt to speed up runtime by doing a binary search to narrow down the range
func findWinningRangesBinarySearch(timeToRace int, maxDistance int) int {
	binSearchUpperBound := 0
	binSearchLowerBound := 0

	//attempt to narrow down range of possible winning ranges
	// fmt.Println("Finding Ranges")
	t := 1
	for {
		d := findDistance(t, timeToRace)
		if d > maxDistance {
			if binSearchLowerBound == 0 && binSearchUpperBound == 0 {
				// fmt.Println("Setting lower bound ", t)
				binSearchLowerBound = t
			} else if binSearchUpperBound == 0 {
				// fmt.Println("Setting upper bound ", t)
				binSearchUpperBound = t
			}
		}
		t *= 2
		if binSearchLowerBound != 0 && binSearchUpperBound != 0 {
			break
		}
	}

	// fmt.Println("Found bounds")

	// iterate down from lower bound to find exact range
	possibleLowerBound := binSearchLowerBound - 1
	for {
		d := findDistance(possibleLowerBound, timeToRace)
		if d < maxDistance {
			break
		}
		possibleLowerBound -= 1
	}

	// fmt.Println("Found lowerBound")

	// iterate up from upper bound to find exact range
	possibleUpperBound := binSearchUpperBound + 1
	for {
		d := findDistance(possibleUpperBound, timeToRace)
		if d < maxDistance {
			break
		}
		possibleUpperBound += 1
	}

	// fmt.Println("Found upperBound")

	return possibleUpperBound - possibleLowerBound - 1
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

	// read from input
	times := []string{}
	distances := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		for _, el := range split[1:] {
			if el == "" {
				continue
			}
			if strings.Contains(line, "Time:") {
				times = append(times, el)
			} else if strings.Contains(line, "Distance:") {
				distances = append(distances, el)
			}
		}
	}

	totalTime := convertStringToInt(strings.Join(times, ""))
	totalDistance := convertStringToInt(strings.Join(distances, ""))

	fmt.Println("TotalTime: ", totalTime, " TotalDistance: ", totalDistance)

	numOfPossibleWins := 0
	// iterate through races and determine the ran
	numOfPossibleWins = findWinningRangesBinarySearch(totalTime, totalDistance)

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println("NumberOfPossibleWins ", numOfPossibleWins)
}
