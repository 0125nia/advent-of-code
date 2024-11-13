package day2_rock_paper_scissors

import "fmt"

//https://adventofcode.com/2022/day/2

// A    B     C
// X    Y     Z
// Rock paper scissor
// 1    2     3
// win 6 draw 3 lose 0
func calcTotalScoreInRPS() {
	input := readInput()
	if input == nil {
		return
	}
	sum := 0
	for _, game := range input {
		score := calcScorePerGame(game[0], game[1])
		sum += score
	}
	fmt.Println(sum)
}

var scoreMap = map[rune]map[rune]int{
	'X': {'A': 4, 'B': 1, 'C': 7},
	'Y': {'A': 8, 'B': 5, 'C': 2},
	'Z': {'A': 3, 'B': 9, 'C': 6},
}

func calcScorePerGame(rival, me rune) int {
	return scoreMap[me][rival]
}
