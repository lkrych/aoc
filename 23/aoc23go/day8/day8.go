package day8

import (
	"flag"
	"fmt"
	"strings"

	"github.com/lkrych/aoc23go/input"
)

type Node struct {
	left  string
	right string
}

func removeChars(s string, charsToRemove string) string {
	var sb strings.Builder
	for _, c := range s {
		if !strings.ContainsRune(charsToRemove, c) {
			sb.WriteRune(c)
		}
	}
	return sb.String()
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
	path := ""
	graph := map[string]Node{}
	moveOntoMap := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			moveOntoMap = true
			continue
		}
		if moveOntoMap {
			// AAA = (BBB, CCC)
			split := strings.Split(line, "=")
			src := strings.TrimSpace(split[0])
			dsts := removeChars(strings.TrimSpace(split[1]), "(,)")
			leftRight := strings.Split(dsts, " ")

			graph[src] = Node{
				left:  leftRight[0],
				right: leftRight[1],
			}
		} else {
			path += line
		}
	}

	fmt.Println("Path: ", path)
	fmt.Println("Graph: ", graph)

	currentStep := "AAA"
	pathEls := strings.Split(path, "")
	stepsTaken := 0
	// Now walk the path until you reach the current step
	for currentStep != "ZZZ" {
		for _, el := range pathEls {
			if currentStep == "ZZZ" {
				break
			}
			stepsTaken += 1
			currentStepNode := graph[currentStep]
			if el == "R" {
				currentStep = currentStepNode.right
			} else if el == "L" {
				currentStep = currentStepNode.left
			} else {
				panic("The path was neither right nor left!")
			}
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println("Steps taken: ", stepsTaken)
}
