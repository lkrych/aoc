package day4

import (
	"flag"
	"fmt"
	"strings"

	"github.com/lkrych/aoc23go/input"
)

func removeWhitespace(s string) string {
	fields := strings.Fields(s) // Fields splits the string s around each instance of one or more consecutive white space characters
	return strings.Join(fields, "")
}

func marshalListIntoMap(l []string) (nMap map[string]bool) {
	nMap = make(map[string]bool)
	for _, item := range l {
		item = removeWhitespace(item)
		// skip empty string
		if item == "" {
			continue
		}
		nMap[item] = true
	}
	return
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
	totalPoints := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		// INITIALIZE PER INPUT LINE VALUES
		pointsPerLine := 0
		splitOnColon := strings.Split(line, ":")
		splitOnBar := strings.Split(splitOnColon[1], "|")
		winningNums := strings.TrimSpace(splitOnBar[0])
		winningNumsList := strings.Split(winningNums, " ")
		winningMap := marshalListIntoMap(winningNumsList)
		candidateNums := strings.TrimSpace(splitOnBar[1])
		candidateNumsList := strings.Split(candidateNums, " ")
		candidateMap := marshalListIntoMap(candidateNumsList)

		// fmt.Println("winning map: %v, \n", winningMap)
		// fmt.Println("candidate map: %v, \n", candidateMap)

		for win := range winningMap {
			if _, ok := candidateMap[win]; ok {
				// if the winning number was found in candidates string
				if pointsPerLine == 0 {
					pointsPerLine += 1
				} else {
					pointsPerLine *= 2
				}
			}
		}
		fmt.Println(pointsPerLine)
		totalPoints += pointsPerLine
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	fmt.Println("Total Points: ", totalPoints)
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

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// BEGIN CODING FOR DAY HERE
	// INITIALIZE GLOBAL VALUES

	copiesCollection := map[int]int{}
	for i := 0; i < len(lines); i++ {
		// init copiesCollection map
		copiesCollection[i] = 1
	}

	for i := 0; i < len(lines); i++ {
		lineCount := copiesCollection[i]
		line := lines[i]
		fmt.Println(line)
		// fmt.Printf("%v\n", copiesCollection)
		// INITIALIZE PER INPUT LINE VALUES
		splitOnColon := strings.Split(line, ":")
		splitOnBar := strings.Split(splitOnColon[1], "|")
		winningNums := strings.TrimSpace(splitOnBar[0])
		winningNumsList := strings.Split(winningNums, " ")
		winningMap := marshalListIntoMap(winningNumsList)
		candidateNums := strings.TrimSpace(splitOnBar[1])
		candidateNumsList := strings.Split(candidateNums, " ")
		candidateMap := marshalListIntoMap(candidateNumsList)

		// iterate over the number of copies
		for j := 0; j < lineCount; j++ {
			// keep track of the number of matches
			matches := 0
			for win := range winningMap {
				if _, ok := candidateMap[win]; ok {
					matches += 1
				}
			}
			// fmt.Printf("found %d matches \n", matches)

			// now add the copies back to the map
			for k := 1; k < matches+1; k++ {
				val := copiesCollection[i+k]
				copiesCollection[i+k] = val + 1
			}
		}
	}
	totalPoints := 0
	for _, lineCount := range copiesCollection {
		totalPoints += lineCount
	}

	fmt.Println("Total Points: ", totalPoints)
}
