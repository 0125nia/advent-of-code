package day6_tuning_trouble

import "fmt"

func getProcessingMarkerP2() {
	input := getInput()
	var arr [14]rune
	for i := 0; i < 14; i++ {
		arr[i] = rune(input[i])
	}
	for i := 14; i < len(input); i++ {
		arr[i%14] = rune(input[i])
		if isUniqueP2(arr) {
			fmt.Println(i + 1)
			return
		}
	}
}

func isUniqueP2(arr [14]rune) bool {
	m := make(map[rune]bool)
	for _, v := range arr {
		m[v] = true
	}
	return len(m) == 14
}
