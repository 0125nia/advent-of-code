package day4_ceres_search

import (
	"bufio"
	bytes2 "bytes"
	"os"
)

func readInput() [][]rune {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var diagram [][]rune

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		bytes := scanner.Bytes()
		runes := bytes2.Runes(bytes)
		diagram = append(diagram, runes)
	}
	return diagram
}
