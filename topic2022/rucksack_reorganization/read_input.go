package rucksack_reorganization

import (
	"bufio"
	"fmt"
	"os"
)

func readInput() []string {
	// open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var input []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	return input
}
