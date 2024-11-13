package calories

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ReadInput https://adventofcode.com/20XX/day/X/input
func readInput() []int {
	// open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var blockSum []int
	var currentBlock []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// If it is empty, it indicates the end of a block
		if line == "" {
			if len(currentBlock) > 0 {
				sum := 0
				for _, num := range currentBlock {
					sum += num
				}
				blockSum = append(blockSum, sum) // Adds the current block to blocks
				currentBlock = []int{}           // Reset current block
			}
			continue
		}
		// Converts the current row to an integer and adds the current block
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error converting line to integer:", err)
			return nil
		}
		currentBlock = append(currentBlock, num)
	}

	// Add the last block to blocks (if the last row is not empty)
	if len(currentBlock) > 0 {
		sum := 0
		for _, num := range currentBlock {
			sum += num
		}
		blockSum = append(blockSum, sum)
	}
	return blockSum
}
