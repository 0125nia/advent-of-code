package day3_mull_it_over

import (
	"bufio"
	"fmt"
	"os"
)

func readInput() []string {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}
