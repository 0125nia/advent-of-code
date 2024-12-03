package day2_red_nosed_reports

import "fmt"

func monotonicityPart2() {
	input := readInput()
	count := 0

	for _, arr := range input {
		for i := 0; i < len(arr); i++ {
			// 排除第 i 个元素，检查剩余部分
			newRow := append(append([]int(nil), arr[:i]...), arr[i+1:]...)
			if isMonotonicWithRemove(newRow) {
				count++
				break
			}
		}
	}
	fmt.Println(count)
}

func isMonotonicWithRemove(arr []int) bool {

	diff := make([]int, len(arr)-1)

	for i := 0; i < len(arr)-1; i++ {
		diff[i] = arr[i+1] - arr[i]
	}

	if isMonotonicByDiff(diff) {
		return true
	}
	return false
}

// Determine whether the array of differences is monotonic
// and the adjacent differences are between 1 and 3
func isMonotonicByDiff(diff []int) bool {
	increasing, decreasing := true, true
	for i := 0; i < len(diff); i++ {
		d := abs(diff[i])
		if d < 1 || d > 3 {
			return false // If the difference is not between 1 and 3, return false
		}
		if diff[i] > 0 {
			decreasing = false
		} else if diff[i] < 0 {
			increasing = false
		}
	}
	return increasing || decreasing
}
