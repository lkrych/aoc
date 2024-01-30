package day7

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

type CamelCardType int

const (
	HighCard CamelCardType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type CamelCard struct {
	cards    string
	handType CamelCardType
	bid      int
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
	totalWinnings := 0

	cards := []CamelCard{}
	// first parse through all the hands
	for scanner.Scan() {
		// example input: 32T3K 765
		line := scanner.Text()
		split := strings.Split(line, " ")

		c := CamelCard{}
		c.cards = split[0]
		c.bid = convertStringToInt(split[1])
		c.FindCardType()
		cards = append(cards, c)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println("TotalWinnings", totalWinnings)
}
