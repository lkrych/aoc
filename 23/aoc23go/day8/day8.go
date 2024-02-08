package day8

import (
	"flag"
	"fmt"
	"strings"

	"github.com/lkrych/aoc23go/input"
)

type Node struct {
	left      string
	right     string
	foundLoop bool
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

// euclidian algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Function to find the Least Common Multiple (LCM) of two numbers
func lcm(a, b int) int {
	return a / gcd(a, b) * b // Ensure the multiplication does not overflow
}

// Function to find the LCM of an array of numbers
func lcmOfArray(arr []int) int {
	fmt.Println("LCM of: ", arr)
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result = lcm(result, arr[i])
	}
	return result
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

	path := ""
	graph := map[string]Node{}
	startingNodes := []string{}
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
			if src[2] == 'A' {
				startingNodes = append(startingNodes, src)
			}
		} else {
			path += line
		}
	}
	fmt.Println("Starting nodes: ", startingNodes)
	// Allocate a new slice with the same length as the original
	pathEls := strings.Split(path, "")
	loopCounts := make([]int, len(startingNodes))
	foundLoops := make([]bool, len(startingNodes))

OuterLoop:
	for {
		for _, el := range pathEls {
			for i, node := range startingNodes {
				foundLoop := foundLoops[i]
				// if we've found the loop already, no need to increment the counter
				if foundLoop {
					continue
				}
				currentStep := node
				currentStepNode := graph[currentStep]
				if el == "R" {
					startingNodes[i] = currentStepNode.right
				} else if el == "L" {
					startingNodes[i] = currentStepNode.left
				} else {
					panic("The path was neither right nor left!")
				}

				// if we find the end, set that we've found a loop and don't increment the counter
				if node[2] == 'Z' {
					foundLoops[i] = true
					continue
				}
				// increment the loop counter
				loopCounts[i] += 1
			}

			// if we have counted all the loops then we are done!
			foundAllLoops := true
			for _, fl := range foundLoops {
				foundAllLoops = fl
			}

			if foundAllLoops {
				break OuterLoop
			}
		}
	}

	// find the LCM of all the loops

	lcm := lcmOfArray(loopCounts)

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println("Steps taken: ", lcm)
}
