package day3_rucksack_reorganization

import "fmt"

var priorityMap map[rune]int

func sumPrioritiesP1() {
	input := readInput()
	if input == nil {
		return
	}
	generatePriorityMap()
	sum := 0
	for _, str := range input {
		priority := calcPriority(str)
		sum += priority
	}
	fmt.Println(sum)
}

// calculate priority of common item in two compartments
func calcPriority(str string) int {
	charMap := make(map[rune]bool)
	for i := 0; i < len(str)/2; i++ {
		charMap[rune(str[i])] = true
	}
	for i := len(str) / 2; i < len(str); i++ {
		if charMap[rune(str[i])] {
			return priorityMap[rune(str[i])]
		}
	}
	return 0
}

// generate priority map
func generatePriorityMap() {
	priorityMap = make(map[rune]int)
	for i := 0; i < 26; i++ {
		priorityMap[rune('a'+i)] = i + 1
		priorityMap[rune('A'+i)] = i + 27
	}
}
