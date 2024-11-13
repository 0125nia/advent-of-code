package day3_rucksack_reorganization

import "fmt"

func sumPrioritiesP2() {
	input := readInput()
	if input == nil {
		return
	}
	generatePriorityMap()
	sum := 0
	for i := 0; i < len(input); i += 3 {
		priority := calcPriorityP2(input[i], input[i+1], input[i+2])
		sum += priority
	}
	fmt.Println(sum)
}

func calcPriorityP2(first, second, third string) int {
	charMap := make(map[rune]int)
	for _, char := range first {
		charMap[char] = 1
	}
	for _, char := range second {
		if charMap[char] == 1 {
			charMap[char] = 2
		}
	}
	for _, char := range third {
		if charMap[char] == 2 {
			charMap[char] = 3
		}
	}
	for key, value := range charMap {
		if value == 3 {
			return priorityMap[key]
		}
	}
	return 0
}
