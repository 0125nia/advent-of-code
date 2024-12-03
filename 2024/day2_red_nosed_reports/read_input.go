package day2_red_nosed_reports

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() [][]int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var input [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		arr := make([]int, 0)
		for _, s := range split {
			num, _ := strconv.Atoi(s)
			arr = append(arr, num)
		}
		input = append(input, arr)
	}

	return input
}
