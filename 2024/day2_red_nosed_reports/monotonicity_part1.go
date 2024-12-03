package day2_red_nosed_reports

import (
	"fmt"
)

func monotonicityPart1() {

	input := readInput()

	count := 0

	for _, arr := range input {
		fmt.Print(arr)
		if isMonotonic(arr) {
			fmt.Print(" true")
			count++
		}
		fmt.Println()
	}

	fmt.Println(count)
}

func isMonotonic(arr []int) bool {
	increasing, decreasing := true, true
	for i := 0; i < len(arr)-1; i++ {
		abs := abs(arr[i], arr[i+1])
		if abs > 3 || abs < 1 {
			return false
		}
		if arr[i]-arr[i+1] > 0 {
			decreasing = false
		} else if arr[i]-arr[i+1] < 0 {
			increasing = false
		}
	}
	return increasing || decreasing
}

func abs(a, b int) int {
	absDiff := a - b
	if absDiff < 0 {
		absDiff = -absDiff
	}
	return absDiff
}
