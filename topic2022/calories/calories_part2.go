package calories

import (
	"fmt"
)

// https://adventofcode.com/2022/day/1#part2

var topThree [3]int

func calculateCaloriesInFirstThree() {
	input := readInput()
	if input == nil {
		return
	}
	initialTopThree(input[0], input[1], input[2])
	fmt.Println(topThree)
	for _, sum := range input {
		sliding(sum)
	}
	sum := 0
	for _, top := range topThree {
		sum += top
	}
	fmt.Println(sum)
}

func sliding(num int) {
	if num > topThree[0] {
		topThree[2] = topThree[1]
		topThree[1] = topThree[0]
		topThree[0] = num
	} else if num > topThree[1] {
		topThree[2] = topThree[1]
		topThree[1] = num
	} else if num > topThree[2] {
		topThree[2] = num
	}
}

// 由于只有三个数，用条件判断进行排序更高效，减少了函数调用和切片包装。
func initialTopThree(num1, num2, num3 int) {
	// 手动比较和排序三个数
	if num1 < num2 {
		num1, num2 = num2, num1
	}
	if num1 < num3 {
		num1, num3 = num3, num1
	}
	if num2 < num3 {
		num2, num3 = num3, num2
	}
	// 将排序后的数放入 topThree
	topThree = [3]int{num1, num2, num3}
}
