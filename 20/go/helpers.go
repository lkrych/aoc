package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// ReadFile takes a filename as argument and returns a slice of strings for each line
func ReadFile(filename string) ([]string, error) {
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
		trimmed := strings.TrimSuffix(line, "\n")

		if err != nil && err != io.EOF {
			break
		}

		data = append(data, trimmed)

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
