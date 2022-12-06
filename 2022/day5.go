package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Stack struct {
	values []string
}

func (s *Stack) push(value string) {
	s.values = append(s.values, value)
}

func (s *Stack) pop() string {
	lastIndex := len(s.values) - 1
	value := s.values[lastIndex]
	s.values = s.values[:lastIndex]

	return value
}

func (s *Stack) isEmpty() bool {
	return len(s.values) == 0
}

func makeStacks(n int) []*Stack {
	result := make([]*Stack, n)
	for i := 0; i < n; i++ {
		result[i] = &Stack{values: make([]string, 0)}
	}

	return result
}

func main() {
	// file, err := os.Open("day5_sample_input.txt")
	file, err := os.Open("day5_input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	stackCount := 0
	var stacks []*Stack

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineLength := len(line)

		if lineLength == 0 {
			for i, stack := range stacks {
				newStack := Stack{values: make([]string, 0)}
				for !stack.isEmpty() {
					value := stack.pop()
					newStack.push(value)
				}

				stacks[i] = &newStack
			}

			continue
		}

		if stackCount == 0 {
			stackCount = (lineLength + 1) / 4
			stacks = makeStacks(stackCount)
		}

		if strings.Contains(line, "move") {
			reg, _ := regexp.Compile("[[:digit:]]+")
			matches := reg.FindAllString(line, -1)

			nMoves, _ := strconv.Atoi((matches[0]))
			srcIndex, _ := strconv.Atoi((matches[1]))
			srcIndex--
			dstIndex, _ := strconv.Atoi((matches[2]))
			dstIndex--

			srcStack := stacks[srcIndex]
			dstStack := stacks[dstIndex]

			// Logic for part 1
			// for i := 0; i < nMoves; i++ {
			// 	dstStack.push(srcStack.pop())
			// }

			// Logic for part 2
			tmpStack := Stack{values: make([]string, 0)}
			for i := 0; i < nMoves; i++ {
				tmpStack.push(srcStack.pop())
			}
			for i := 0; i < nMoves; i++ {
				dstStack.push(tmpStack.pop())
			}
		} else {
			stackIndex := 0
			for i := 1; i < lineLength; i += 4 {
				value := string(line[i])
				if value >= "1" && value <= "9" {
					break
				} else if value != " " {
					stacks[stackIndex].push(value)
				}

				stackIndex += 1
			}
		}
	}

	for _, stack := range stacks {
		fmt.Print(stack.pop())
	}
}
