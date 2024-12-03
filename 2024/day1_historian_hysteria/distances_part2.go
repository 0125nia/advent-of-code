package day1_historian_hysteria

import "fmt"

func distancePart2() {
	input := readInput()
	var frequencies = make(map[int]int)
	sum := 0

	for _, item := range input[1] {
		frequencies[item]++
	}
	for _, item := range input[0] {
		sum += item * frequencies[item]
	}
	fmt.Println(sum)
}
