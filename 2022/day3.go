package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func getPriority(runeValue rune) uint64 {
	// Map characters A - Z to 27 - 52
	// Map characters a - z to 1 - 26
	asciiDecimal := int(runeValue)
	if asciiDecimal <= 90 {
		asciiDecimal -= 38
	} else {
		asciiDecimal -= 96
	}

	return uint64(asciiDecimal)
}

func makeInitialItemTypes() map[rune]bool {
	commonItemTypes := map[rune]bool{}
	for r := 'a'; r <= 'z'; r++ {
		R := unicode.ToUpper(r)
		commonItemTypes[r] = true
		commonItemTypes[R] = true
	}

	return commonItemTypes
}

func solvePart1() {
	// file, err := os.Open("day3_sample_input.txt")
	file, err := os.Open("day3_input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	prioritySum := uint64(0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		midPoint := len(line) / 2
		firstHalfItems := map[rune]bool{}
		for _, runeValue := range line[:midPoint] {
			firstHalfItems[runeValue] = true
		}

		for _, runeValue := range line[midPoint:] {
			if !firstHalfItems[runeValue] {
				continue
			}

			prioritySum += getPriority(runeValue)
			break
		}
	}

	fmt.Println("The total priority is: ", prioritySum)
}

func solvePart2() {
	// file, err := os.Open("day3_sample_input.txt")
	file, err := os.Open("day3_input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	prioritySum := uint64(0)

	commonItemTypes := makeInitialItemTypes()
	positionInGroup := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		newCommonItemTypes := map[rune]bool{}
		for _, runeValue := range line {
			if commonItemTypes[runeValue] {
				newCommonItemTypes[runeValue] = true
			}
		}
		commonItemTypes = newCommonItemTypes

		positionInGroup += 1

		if positionInGroup == 3 {
			for runeValue := range commonItemTypes {
				prioritySum += getPriority(runeValue)
				break
			}

			commonItemTypes = makeInitialItemTypes()
			positionInGroup = 0
		}
	}

	fmt.Println("The total priority is: ", prioritySum)
}

func main() {
	// solvePart1()
	solvePart2()
}
