package main

import "fmt"

func main() {
	data, err := ReadFile("../input/aoc3.txt")
	if err != nil {
		fmt.Printf("readFile failed: %v\n", err)
	}
	tobogganDirs := []int{3, 1}
	ans := tobogganPath(data, tobogganDirs[0], tobogganDirs[1])
	fmt.Printf("There trees encountered are %d \n", ans)
}

func tobogganPath(data []string, xVel int, yVel int) int {
	xPos := 0
	treeCount := 0
	for _, row := range data {
		fmt.Printf("xPos: %d\n", xPos)
		if len(row) < 1 {
			continue
		}
		if row[xPos] == '#' {
			treeCount++
		}
		xPos += xVel
		// the map is stable so if we reach the end of the row, we just wrap back around
		xPos = xPos % len(row)
	}
	return treeCount
}
