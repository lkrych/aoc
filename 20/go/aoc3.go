package main

// func main() {
// 	data, err := ReadFile("../input/aoc3.txt")
// 	if err != nil {
// 		fmt.Printf("readFile failed: %v\n", err)
// 	}
// 	ans := 1
// 	tobogganDirs := [][]int{{3, 1}, {1, 1}, {5, 1}, {7, 1}, {1, 2}}
// 	for _, tDir := range tobogganDirs {
// 		ans *= tobogganPath(data, tDir[0], tDir[1])
// 	}
// 	fmt.Printf("There trees encountered are %d \n", ans)
// }

func tobogganPath(data []string, xVel int, yVel int) int {
	xPos, yPos := 0, 0
	treeCount := 0

	for yPos < len(data) {
		row := data[yPos]
		if len(row) < 1 {
			break
		}
		if row[xPos] == '#' {
			treeCount++
		}
		xPos += xVel
		yPos += yVel
		// the map is stable so if we reach the end of the row, we just wrap back around
		xPos = xPos % len(row)
	}
	return treeCount
}
