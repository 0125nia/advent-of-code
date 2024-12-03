package day3_mull_it_over

import (
	"fmt"
	"regexp"
	"strconv"
)

func multiplicationSeek1() {
	input := readInput()
	res := 0
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	for _, line := range input {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			res += mul(match[1], match[2])
		}
	}
	fmt.Println(res)
}

func mul(a, b string) int {
	num1, _ := strconv.Atoi(a)
	num2, _ := strconv.Atoi(b)
	return num1 * num2
}
