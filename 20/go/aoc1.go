package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

// func main() {
// 	data, err := ReadFile("../input/aoc1.txt")
// 	if err != nil {
// 		fmt.Printf("readFile failed: %v\n", err)
// 	}
// 	target := 2020
// 	ans := threeSum(data, target)
// 	fmt.Printf("Three sum of the data is %d\n", ans)
// }

func twoSum(data []string, target int) int {
	m := make(map[int]bool)

	for _, line := range data {
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		m[n] = true
		diff := target - n
		if _, ok := m[diff]; ok {
			return n * diff
		}
	}
	return 0
}

func threeSum(data []string, target int) int {
	intData := make([]int, 0)
	for _, line := range data {

		if len(line) < 1 {
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		intData = append(intData, n)
	}
	// first sort the input
	sort.Ints(intData)
	lower := 0
	middle := 1
	higher := len(intData) - 1
	currentSum := 0
	// create a sliding window that checks all the vals based on their sums
	for higher > lower {
		currentSum = intData[lower] + intData[middle] + intData[higher]
		if currentSum == target {
			return intData[lower] * intData[middle] * intData[higher]
		}
		if currentSum < target {
			if higher-middle > 1 {
				middle++
			} else {
				lower++
				middle = lower + 1
			}
		} else {
			//reset because we've gone too high, this time with the higher value decremented
			if middle-lower > 1 {
				lower++
			} else {
				higher--
				middle = lower + 1
			}
		}
	}
	return 0
}
