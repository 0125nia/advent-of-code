package day5_supply_stacks

import "fmt"

var crateMover9001StackMap map[int]*CrateMover9001Stack

func getCrateMover9001StackTopBox() {
	stacks, moves := read()
	crateMover9001StackMap = generateCrateMover9001(stacks)
	for _, moveArr := range moves {
		moveP2(moveArr)
	}
	fmt.Println(getTopItemsP2())
}

func moveP2(moveArr []int) {
	count := moveArr[0]
	from := moveArr[1]
	to := moveArr[2]
	pop, _ := crateMover9001StackMap[from].Pop(count)
	crateMover9001StackMap[to].Push(pop...)
}

func getTopItemsP2() string {
	runes := make([]rune, len(crateMover9001StackMap))
	for i := 1; i <= len(crateMover9001StackMap); i++ {
		peek, b := crateMover9001StackMap[i].Peek()
		if b {
			runes = append(runes, peek)
		}
	}
	return string(runes)
}

func generateCrateMover9001(stacks []string) map[int]*CrateMover9001Stack {
	lastLine := stacks[len(stacks)-1]
	var crateMoverIndexs []int
	crateMoverStackMap := make(map[int]*CrateMover9001Stack)
	for index, char := range lastLine {
		if char >= '1' && char <= '9' {
			crateMoverIndexs = append(crateMoverIndexs, index)
			crateMoverStackMap[int(char-'0')] = &CrateMover9001Stack{}
		}
	}
	for i := len(stacks) - 2; i >= 0; i-- {
		stacksLine := stacks[i]
		for i, craneMover9001Stack := range crateMoverStackMap {
			index := crateMoverIndexs[i-1]
			if len(stacksLine) < index+1 || stacksLine[index] == ' ' {
				continue
			}
			craneMover9001Stack.Push(rune(stacksLine[index]))
		}
	}
	return crateMoverStackMap
}

type CrateMover9001Stack []rune

func (s *CrateMover9001Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *CrateMover9001Stack) Peek() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return (*s)[len(*s)-1], true
}

func (s *CrateMover9001Stack) Pop(count int) ([]rune, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	if count > len(*s) {
		count = len(*s)
	}

	index := len(*s) - count
	element := (*s)[index:]
	*s = (*s)[:index]
	return element, true
}

// Push multiple values onto the stack 4 5 6
func (s *CrateMover9001Stack) Push(v ...rune) {
	for _, r := range v {
		*s = append(*s, r)
	}
}
