package input

import (
	"bufio"
	"fmt"
	"os"
)

// getFileName asks stdin for a file name and returns answer as string
func GetFileName() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter test file: ")

	// Read input from stdin line by line
	var input string
	for scanner.Scan() {
		input = scanner.Text()
	}

	return input
}

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
