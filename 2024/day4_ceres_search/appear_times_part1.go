package day4_ceres_search

import "fmt"

var input [][]rune
var directions [][]int
var l int
var XMASLen int
var XMAS = []rune("XMAS")

func init() {
	input = readInput()
	l = len(input)
	directions = [][]int{
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
	}
	XMASLen = len("XMAS")
}
func XMASAppearTimes1() {
	count := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			p := Point{i, j}
			if p.pair(XMAS[0]) {
				count += p.search()
			}
		}
	}
	fmt.Println(count)
}

type Point struct {
	x int
	y int
}

func (p *Point) search() int {
	count := 0
	for _, direction := range directions {
		if p.check(direction) {
			count++
		}
	}
	return count
}

func (p *Point) check(direction []int) bool {
	last := p.add(direction, XMASLen-1)

	if !last.isValid() || !last.pair(XMAS[XMASLen-1]) {
		return false
	}
	for i := XMASLen - 2; i > 0; i-- {
		point := p.add(direction, i)
		if !point.pair(XMAS[i]) {
			return false
		}
	}
	return true
}

func (p *Point) isValid() bool {
	if p.x < 0 || p.x >= l || p.y < 0 || p.y >= l {
		return false
	}
	return true
}

func (p *Point) add(direction []int, time int) Point {
	x := p.x
	y := p.y
	for i := 0; i < time; i++ {
		x += direction[0]
		y += direction[1]
	}
	return Point{x, y}
}

func (p *Point) pair(r rune) bool {
	return input[p.x][p.y] == r
}
