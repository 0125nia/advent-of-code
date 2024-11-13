package day4_camp_cleanup

import "fmt"

func calcFullyContainNum() {
	input := readInput()
	if input == nil {
		return
	}
	count := 0
	for _, rangePair := range input {
		if judgePairFullyContain(rangePair) {
			fmt.Println(rangePair[0], rangePair[1], rangePair[2], rangePair[3])
			count++
		}
	}
	fmt.Println(count)
}

// 4 5 6 7 8 4-8
//
//	5 6 7   5-7
//
// 4 <= 5 <= 8
// 4 <= 7 <= 8
// only 4 <= 5 and 7 <= 8
func judgePairFullyContain(rangePair []int) bool {
	return rangePair[0] <= rangePair[2] &&
		rangePair[1] >= rangePair[3] ||
		rangePair[2] <= rangePair[0] &&
			rangePair[1] >= rangePair[3]
}
