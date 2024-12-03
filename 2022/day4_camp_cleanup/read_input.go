package day4_camp_cleanup

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() [][]int {
	// open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var input [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// replace all commas and dashes with spaces
		line = strings.ReplaceAll(line, ",", " ")
		line = strings.ReplaceAll(line, "-", " ")
		fields := strings.Split(line, " ")
		var nums []int

		// convert string to int
		for _, field := range fields {
			num, _ := strconv.Atoi(field)
			nums = append(nums, num)
		}
		input = append(input, nums)
	}
	return input
}
