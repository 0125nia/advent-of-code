package day4_camp_cleanup

import "fmt"

func calcOverlaps() {
	input := readInput()
	if input == nil {
		return
	}
	count := 0
	for _, rangePair := range input {
		if judgePairOverlap(rangePair) {
			count++
		}
	}
	fmt.Println(count)
}

// index 0 < 2 && 1 > 2 || 2 < 0 && 3 > 0
func judgePairOverlap(rangePair []int) bool {
	return rangePair[0] <= rangePair[2] &&
		rangePair[1] >= rangePair[2] ||
		rangePair[2] <= rangePair[0] &&
			rangePair[3] >= rangePair[0]
}
