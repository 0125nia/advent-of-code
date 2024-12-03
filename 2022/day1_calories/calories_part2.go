package day1_calories

import (
	"fmt"
)

// https://adventofcode.com/2022/day/1#part2

var topThree [3]int

func calculateCaloriesInTopThree() {
	input := readInput()
	if input == nil {
		return
	}
	initialTopThree(input[0], input[1], input[2])
	fmt.Println(topThree)
	for _, sum := range input {
		sliding(sum)
	}
	sum := 0
	for _, top := range topThree {
		sum += top
	}
	fmt.Println(sum)
}

func sliding(num int) {
	if num > topThree[0] {
		topThree[2] = topThree[1]
		topThree[1] = topThree[0]
		topThree[0] = num
	} else if num > topThree[1] {
		topThree[2] = topThree[1]
		topThree[1] = num
	} else if num > topThree[2] {
		topThree[2] = num
	}
}

// Since there are only three numbers, sorting with conditional judgments is more efficient, reducing function calls and slice wrapping.
func initialTopThree(num1, num2, num3 int) {
	// Compare and sort three numbers manually
	if num1 < num2 {
		num1, num2 = num2, num1
	}
	if num1 < num3 {
		num1, num3 = num3, num1
	}
	if num2 < num3 {
		num2, num3 = num3, num2
	}
	// Place the sorted number in topThree
	topThree = [3]int{num1, num2, num3}
}
