package day1_historian_hysteria

import (
	"fmt"
	"math"
	"sort"
)

func distancePart1() {
	input := readInput()

	sort.Ints(input[0])
	sort.Ints(input[1])

	var sum float64

	for i := 0; i < len(input[0]); i++ {
		sum += math.Abs(float64(input[1][i]) - float64(input[0][i]))
	}

	fmt.Println(int64(sum))
}
