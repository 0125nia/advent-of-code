package day1_historian_hysteria

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() [2][]int {
	// open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var input [2][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(split[0])
		num2, _ := strconv.Atoi(split[1])
		input[0] = append(input[0], num1)
		input[1] = append(input[1], num2)
	}
	return input
}
