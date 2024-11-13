package day5_supply_stacks

import (
	"fmt"
)

var stack map[int]*Stack

func getTopBox() {
	stacks, moves := readInput()
	stack = generateStack(stacks)

	for _, moveArr := range moves {
		move(moveArr)
	}
	fmt.Println(getTopItems())
}

func move(arr []int) {
	if arr[0] == 0 {
		return
	}
	count := arr[0]
	from := arr[1]
	to := arr[2]
	if count > stack[from].Size() {
		count = stack[from].Size()
	}
	for i := 0; i < count; i++ {
		item, _ := stack[from].Pop()
		stack[to].Push(item)
	}
}

func getTopItems() string {
	runes := make([]rune, len(stack))
	for i := 1; i <= len(stack); i++ {
		peek, b := stack[i].Peek()
		if b {
			runes = append(runes, peek)
		}
	}
	return string(runes)
}

// generateStack create a map of stack
func generateStack(stacks []string) map[int]*Stack {
	lastLine := stacks[len(stacks)-1]
	var stackIndex []int
	stacksMap := make(map[int]*Stack)
	for index, char := range lastLine {
		if char >= '1' && char <= '9' {
			stackIndex = append(stackIndex, index)
			stacksMap[int(char-'0')] = &Stack{}
		}
	}
	for i := len(stacks) - 2; i >= 0; i-- {
		stacksLine := stacks[i]
		for i, stack := range stacksMap {
			index := stackIndex[i-1]
			if len(stacksLine) < index+1 || stacksLine[index] == ' ' {
				continue
			}
			stack.Push(rune(stacksLine[index]))
		}
	}
	return stacksMap
}

type Stack []rune

// Push add a new item to the stack.
func (s *Stack) Push(v rune) {
	*s = append(*s, v)
}

// Pop remove the top item of the stack.
func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	index := len(*s) - 1
	value := (*s)[index]
	*s = (*s)[:index]
	return value, true
}

// IsEmpty check if stack is empty or not.
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Size return stack size.
func (s *Stack) Size() int {
	return len(*s)
}

// Peek return top item of the stack.
func (s *Stack) Peek() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	index := len(*s) - 1
	return (*s)[index], true
}
