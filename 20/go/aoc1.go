package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := readFile("../input/aoc1.txt")
	if err != nil {
		fmt.Printf("readFile failed: %v\n", err)
	}
	target := 2020
	ans := twoSum(data, target)
	fmt.Printf("Two sum of the data is %d\n", ans)
}

func readFile(filename string) ([]string, error) {
	data := make([]string, 0)
	file, err := os.Open(filename)
	if err != nil {
		return data, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		data = append(data, line)

		if err != nil {
			break
		}
	}
	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
		return data, err
	}
	return data, nil
}

func twoSum(data []string, target int) int {
	m := make(map[int]bool)

	for _, line := range data {
		trimmed := strings.TrimSuffix(line, "\n")
		n, err := strconv.Atoi(trimmed)
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
