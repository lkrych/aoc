package day10

import (
	"flag"
	"fmt"
	"strings"

	"github.com/lkrych/aoc23go/input"
)

type coordinate struct {
	x int
	y int
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

	// strategy:
	// find S and start exploring until you find S again.

	// to do this we will first load the entire map into memory
	// note: this could be memory-intensive

	sPos := &coordinate{}
	x := 0
	y := 0
	pipeMap := [][]string{}
	for scanner.Scan() {
		x = 0
		line := scanner.Text()
		splitLine := strings.Split(strings.TrimSpace(line), "")
		for _, el := range splitLine {
			if el == "S" {
				sPos.x = x
				sPos.y = y
			}
			x++
		}
		pipeMap = append(pipeMap, splitLine)
		y++
	}

	for _, el := range pipeMap {
		fmt.Println(el)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	dist := 0
	fmt.Println("The longest distance from the start point is: ", dist)
}
