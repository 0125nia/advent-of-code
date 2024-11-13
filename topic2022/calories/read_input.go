package calories

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ReadInput https://adventofcode.com/20XX/day/X/input
func readInput() []int {
	// 打开文件
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var blockSum []int
	var currentBlock []int // 当前块

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// 如果是空行，表示一个块的结束
		if line == "" {
			if len(currentBlock) > 0 {
				sum := 0
				for _, num := range currentBlock {
					sum += num
				}
				blockSum = append(blockSum, sum) // 添加当前块到blocks
				currentBlock = []int{}           // 重置当前块
			}
			continue
		}
		// 将当前行转换为整数并加入当前块
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error converting line to integer:", err)
			return nil
		}
		currentBlock = append(currentBlock, num)
	}

	// 将最后一个块加入blocks（如果最后一行不是空行）
	if len(currentBlock) > 0 {
		sum := 0
		for _, num := range currentBlock {
			sum += num
		}
		blockSum = append(blockSum, sum)
	}
	return blockSum
}
