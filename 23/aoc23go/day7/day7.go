package day7

import (
	"flag"
	"fmt"
	"log"
	"sort"
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

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
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

type CardRank int

// A slice of CardRanks
type ByCardRank []CardRank

// Implement the sort.Interface for ByCardRank
func (a ByCardRank) Len() int           { return len(a) }
func (a ByCardRank) Less(i, j int) bool { return a[i] < a[j] }
func (a ByCardRank) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type CamelCard struct {
	cards     string
	handType  CamelCardType
	bid       int
	cardRanks []CardRank
}

// A slice of CamelCard
type ByCamelCard []CamelCard

// Implement the sort.Interface for ByCamelCard
func (cc ByCamelCard) Len() int { return len(cc) }
func (cc ByCamelCard) Less(i, j int) bool {
	// First, sort by handType
	if cc[i].handType != cc[j].handType {
		return cc[i].handType < cc[j].handType
	}
	// Start by comparing the first card in each hand. If these cards are different, the hand with the stronger
	// first card is considered stronger. If the first card in each hand have the same label, however,
	// then move on to considering the second card in each hand. If they differ, the hand with the higher
	// second card wins; otherwise, continue with the third card in each hand, then the fourth, then the fifth.
	for k := 0; k < len(cc[i].cardRanks); k++ {
		if cc[i].cardRanks[k] != cc[j].cardRanks[k] {
			return cc[i].cardRanks[k] < cc[j].cardRanks[k]
		}
	}
	// default should never be called
	return cc[i].cardRanks[0] != cc[j].cardRanks[0]
}
func (cc ByCamelCard) Swap(i, j int) { cc[i], cc[j] = cc[j], cc[i] }

// split apart the source string and convert it into a list of CardRanks
func (c *CamelCard) findCardRanks() {
	CardRankMap := map[string]int{
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}
	cardRanks := []CardRank{}
	splitRank := strings.Split(c.cards, "")
	for _, rank := range splitRank {
		if isInt(rank) {
			cardRanks = append(cardRanks, CardRank(convertStringToInt(rank)))
		} else {
			crIntVal := CardRankMap[rank]
			cardRanks = append(cardRanks, CardRank(crIntVal))
		}
	}
	c.cardRanks = cardRanks
}

func (c *CamelCard) findHandType() {
	// put cards into map
	m := map[CardRank]int{}
	for _, cr := range c.cardRanks {
		count, found := m[cr]
		if !found {
			m[cr] = 1
		} else {
			m[cr] = count + 1
		}
	}
	// fmt.Println(c.cardRanks, m)
	three := false
	pair := 0
	for _, v := range m {
		switch v {
		case 5:
			c.handType = FiveOfAKind
		case 4:
			c.handType = FourOfAKind
		case 3:
			three = true
		case 2:
			pair += 1
		}
	}
	if three && pair == 1 {
		c.handType = FullHouse
	} else if three {
		c.handType = ThreeOfAKind
	} else if pair == 2 {
		c.handType = TwoPair
	} else if pair == 1 {
		c.handType = OnePair
	} else if c.handType == 0 {
		c.handType = HighCard
	}
}

func (c *CamelCard) FindHandTypeAndHighCard() {
	c.findCardRanks()
	// next we need to find the Hand Type
	c.findHandType()
}

func createCard(line string) CamelCard {
	split := strings.Split(line, " ")

	c := CamelCard{
		cards: split[0],
		bid:   convertStringToInt(split[1]),
	}
	// find the handType and high card
	c.FindHandTypeAndHighCard()
	return c
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
		c := createCard(line)
		// fmt.Println(c.cards, c.handType)
		// fmt.Println()
		cards = append(cards, c)
	}

	sort.Sort(ByCamelCard(cards))

	// find the multiplier
	for i, c := range cards {
		fmt.Println(c.cards, c.handType)
		totalWinnings += (i + 1) * c.bid
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println("TotalWinnings", totalWinnings)
}
