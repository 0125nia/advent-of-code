package day6_tuning_trouble

import "fmt"

func getProcessingMarker() {
	input := getInput()
	var arr [4]rune
	for i := 0; i < 4; i++ {
		arr[i] = rune(input[i])
	}
	for i := 4; i < len(input); i++ {
		arr[i%4] = rune(input[i])
		if isUnique(arr) {
			fmt.Println(i + 1)
			return
		}
	}
}

func isUnique(arr [4]rune) bool {
	m := make(map[rune]bool)
	for _, v := range arr {
		m[v] = true
	}
	return len(m) == 4
}
