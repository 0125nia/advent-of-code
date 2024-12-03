package day1_historian_hysteria

import (
	"fmt"
	"sort"
)

func distancePart1() {
	input := readInput()

	sort.Ints(input[0])
	sort.Ints(input[1])

	var sum int

	for i := 0; i < len(input[0]); i++ {
		sum += abs(input[1][i], input[0][i])
	}

	fmt.Println(sum)
}

func abs(a, b int) int {
	absDiff := a - b
	if absDiff < 0 {
		absDiff = -absDiff
	}
	return absDiff
}
