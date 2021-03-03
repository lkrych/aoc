package main

import (
	"fmt"
	"sort"
)

func main() {
	data, err := ReadFile("../input/aoc5.txt")
	if err != nil {
		fmt.Printf("readFile failed: %v\n", err)
	}

	// ans := highestSeatId(data)

	// fmt.Printf("The highest seat ID is %d \n", ans)
	ans := mySeatId(data)

	fmt.Printf("My seat ID is %d \n", ans)
}

func highestSeatId(data []string) int {
	highestSeatId := 0
	for _, seat := range data {
		if len(seat) < 1 {
			continue
		}
		seatId := seatId(seat)
		if seatId > highestSeatId {
			highestSeatId = seatId
		}
	}
	return highestSeatId
}

func mySeatId(data []string) int {
	seatIds := make([]int, 0)
	for _, seat := range data {
		if len(seat) < 1 {
			continue
		}
		seatId := seatId(seat)
		seatIds = append(seatIds, seatId)
	}
	sort.Ints(seatIds)

	for i, seatId := range seatIds {
		if i == 0 || i == len(seatIds)-1 {
			//skip first and last indices
			continue
		}
		diff := seatIds[i] - seatIds[i-1]
		if diff > 1 {
			return seatId - 1
		}
	}
	return 0
}

func seatId(seatCode string) int {
	// BBFFBBFRLL
	row := seatCode[:7]
	col := seatCode[7:]
	return seatBinary(row, "row")*8 + seatBinary(col, "col")
}

func seatBinary(seatCode string, seatType string) int {
	lowerIdx := 0
	upperIdx := 7
	if seatType == "row" {
		upperIdx = 127
	}
	for i, c := range seatCode {
		diff := (upperIdx - lowerIdx + 1) / 2
		if i == len(seatCode)-1 {
			break
		}
		// cut down options
		switch c {
		case 'F', 'L':
			upperIdx -= diff
		case 'B', 'R':
			lowerIdx += diff
		}
	}
	lastChar := seatCode[len(seatCode)-1]
	seatBinary := 0
	switch lastChar {
	case 'F', 'L':
		seatBinary = lowerIdx
	case 'B', 'R':
		seatBinary = upperIdx
	}
	return seatBinary

}
