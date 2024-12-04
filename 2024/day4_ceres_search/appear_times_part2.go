package day4_ceres_search

import "fmt"

var directions2 [][]int

func init() {
	directions2 = [][]int{
		{1, -1},
		{-1, -1},
		{-1, 1},
		{1, 1},
	}
}

func MASAppearTimes2() {
	count := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			p := Point{i, j}
			if p.pair('A') {
				if p.checkMAS() {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

func (p *Point) checkMAS() bool {
	for i := 0; i < 2; i++ {
		p1 := p.add(directions2[i], 1)
		p2 := p.add(directions2[i+2], 1)
		if !(p1.isValid() && p2.isValid() &&
			((p1.pair('M') && p2.pair('S')) || (p1.pair('S') && p2.pair('M')))) {
			return false
		}
	}
	return true
}
