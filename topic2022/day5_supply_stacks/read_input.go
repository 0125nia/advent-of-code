package day5_supply_stacks

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() ([]string, [][]int) {
	// open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	flag := true

	var stackStr []string
	var moves [][]int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			flag = false
			continue
		}
		if flag {
			stackStr = append(stackStr, line)
		} else {
			nums := handlerMoveLine(line)
			moves = append(moves, nums)
		}
	}
	return stackStr, moves
}

func handlerMoveLine(line string) []int {
	line = strings.ReplaceAll(line, "move ", "")
	line = strings.ReplaceAll(line, "from ", "")
	line = strings.ReplaceAll(line, "to ", "")
	fields := strings.Split(line, " ")
	var nums []int
	for _, field := range fields {
		num, _ := strconv.Atoi(field)
		nums = append(nums, num)
	}
	return nums
}
