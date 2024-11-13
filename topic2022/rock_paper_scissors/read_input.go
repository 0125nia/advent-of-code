package rock_paper_scissors

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput() [][]rune {
	// 打开文件
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var input [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.ReplaceAll(line, " ", "")

		input = append(input, []rune(split))
	}

	return input
}
