package day9

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/lkrych/aoc23go/input"
)

func convertStringListToInt(l []string) []int {
	li := make([]int, len(l))
	for i, el := range l {
		eli, err := strconv.Atoi(el)
		if err != nil {
			// skip over elements that can't be converted into ints
			continue
		}
		li[i] = eli
	}
	return li
}

func findDiffs(l []int) (diffs []int) {
	// fmt.Println("Finding diff of : ", l)
	for i := 0; i < len(l)-1; i++ {
		d := l[i+1] - l[i]
		diffs = append(diffs, d)
	}
	// fmt.Println("diffs: ", diffs)
	// fmt.Println()
	return
}

func allZeroes(l []int) bool {
	for _, el := range l {
		if el != 0 {
			return false
		}
	}
	return true
}

func findDiffSum(l []int) (sum int) {
	diffs := [][]int{
		l,
	}
	for {
		l = findDiffs(l)
		diffs = append(diffs, l)
		if allZeroes(l) {
			break
		}
	}
	// walk backwards through the diffs
	lastEl := 0
	for i := len(diffs) - 1; i >= 0; i-- {
		diff := diffs[i]
		lastEl = diff[len(diff)-1] + lastEl
		// fmt.Println("Last el of :", diff, " is ", lastEl)
	}

	return lastEl
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

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(strings.TrimSpace(line), " ")
		ints := convertStringListToInt(splitLine)
		sum += findDiffSum(ints)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println("The sum of the patterns is: ", sum)
}

func findDiffSumBackward(l []int) (sum int) {
	diffs := [][]int{
		l,
	}
	for {
		l = findDiffs(l)
		diffs = append(diffs, l)
		if allZeroes(l) {
			break
		}
	}
	// walk backwards through the diffs
	firstEl := 0
	for i := len(diffs) - 1; i >= 0; i-- {
		diff := diffs[i]
		firstEl = diff[0] - firstEl
		// fmt.Println("Last el of :", diff, " is ", lastEl)
	}

	return firstEl
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

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(strings.TrimSpace(line), " ")
		ints := convertStringListToInt(splitLine)
		sum += findDiffSumBackward(ints)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println("The sum of the patterns is: ", sum)
}
