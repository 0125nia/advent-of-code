package day1_calories

import (
	"fmt"
)

// https://adventofcode.com/2022/day/1
func calcCalories() {
	input := readInput()
	if input == nil {
		return
	}
	var maxCalories int
	for _, calories := range input {
		if calories > maxCalories {
			maxCalories = calories
		}
	}
	fmt.Println(maxCalories)
}
