package input

import (
	"bufio"
	"fmt"
	"os"
)

// readInputFile takes in a filepath as a string and returns a bufio.Scanner
func ReadInputFile(filePath string) (*bufio.Scanner, error) {
	fmt.Println("Trying to read ", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	// Create a bufio.Scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	return scanner, nil
}
