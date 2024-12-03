package day3_mull_it_over

import (
	"fmt"
	"regexp"
)

func multiplicationSeek2() {
	input := readInput()
	str := "do()"
	for _, line := range input {
		str += line
	}
	str += "don't()"
	res := 0

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	unableMul := regexp.MustCompile(`do\(\)(.*?)don't\(\)`)

	submatches := unableMul.FindAllStringSubmatch(str, -1)
	for _, submatch := range submatches {
		line := submatch[1]
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			res += mul(match[1], match[2])
		}

	}

	fmt.Println(res)
}
