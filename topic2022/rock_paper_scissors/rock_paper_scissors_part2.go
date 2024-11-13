package rock_paper_scissors

import "fmt"

//https://adventofcode.com/2022/day/2#part2

func calcTotalScoreInRPSP2() {
	input := readInput()
	if input == nil {
		return
	}
	sum := 0
	for _, game := range input {
		score := calcScorePerGameP2(game[0], game[1])
		sum += score
	}
	fmt.Println(sum)
}

// A    B     C
// Rock paper scissor

//	X    Y    Z
//
// lose draw win
// 0    3    6
var scoreMap2 = map[rune]map[rune]int{
	'X': {'A': 3, 'B': 1, 'C': 2}, //lose s r p 3 1 2
	'Y': {'A': 4, 'B': 5, 'C': 6}, //draw r p s 1 2 3
	'Z': {'A': 8, 'B': 9, 'C': 7}, //win  p s r 2 3 1
}

// if
func calcScorePerGameP2(rival, me rune) int {
	return scoreMap2[me][rival]
}
