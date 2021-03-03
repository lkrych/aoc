package main

import "fmt"

func main() {
	data, err := ReadFile("../input/aoc5.txt")
	if err != nil {
		fmt.Printf("readFile failed: %v\n", err)
	}

	ans := highestSeatId(data)

	fmt.Printf("The highest seat ID is %d \n", ans)
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

func seatId(seatCode string) int {
	// BBFFBBFRLL
	row := seatCode[:7]
	col := seatCode[7:]
	return seatBinary(row, "row")*8 + seatBinary(col, "col")
}

func seatBinary(seatCode string, seatType string) int {
	fmt.Println("seatCode: ", seatCode)
	lowerIdx := 0
	upperIdx := 7
	if seatType == "row" {
		upperIdx = 127
	}
	for i, c := range seatCode {
		diff := (upperIdx - lowerIdx + 1) / 2
		fmt.Println("upper: ", upperIdx)
		fmt.Println("lower: ", lowerIdx)
		switch c {
		case 'F', 'L':
			//take lower half
			if i < len(seatCode)-1 {
				upperIdx -= diff
			}
		case 'B', 'R':
			//take upper half
			if i < len(seatCode)-1 {
				lowerIdx += diff
			}
		}
	}
	fmt.Println("seatBinary: ", upperIdx+lowerIdx/2)
	return upperIdx + lowerIdx/2
}
