package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Choice struct {
	value               string
	choiceStrongAgainst string
	choiceWeakAgainst   string
	score               uint64
}

func (c *Choice) getScoreAgainst(anotherChoice *Choice) uint64 {
	if c.value == anotherChoice.value {
		return 3
	} else if c.value == anotherChoice.choiceWeakAgainst {
		return 6
	} else {
		return 0
	}
}

var choices = map[string]Choice{
	"A": {value: "A", choiceStrongAgainst: "C", choiceWeakAgainst: "B", score: 1},
	"B": {value: "B", choiceStrongAgainst: "A", choiceWeakAgainst: "C", score: 2},
	"C": {value: "C", choiceStrongAgainst: "B", choiceWeakAgainst: "A", score: 3},
}

var normalize_choices = map[string]string{
	"X": "A",
	"Y": "B",
	"Z": "C",
}

func calculateScoreForPart1(opponentChoiceStr string, yourChoiceStr string) uint64 {
	yourChoiceStrNormalized := normalize_choices[yourChoiceStr]
	yourChoice := choices[yourChoiceStrNormalized]
	opponentChoice := choices[opponentChoiceStr]
	return yourChoice.score + yourChoice.getScoreAgainst(&opponentChoice)
}

func calculateScoreForPart2(opponentChoiceStr string, roundResult string) uint64 {
	var yourChoice Choice
	var scoreForResult uint64
	opponentChoice := choices[opponentChoiceStr]

	switch roundResult {
	case "X": // you lose
		yourChoice = choices[opponentChoice.choiceStrongAgainst]
		scoreForResult = 0
	case "Y": // it's a draw
		yourChoice = opponentChoice
		scoreForResult = 3
	case "Z": // you win
		yourChoice = choices[opponentChoice.choiceWeakAgainst]
		scoreForResult = 6
	}

	return yourChoice.score + scoreForResult
}

func main() {
	file, err := os.Open("day2_input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	totalScore := uint64(0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputs := strings.Split(line, " ")
		// totalScore += calculateScoreForPart1(inputs[0], inputs[1])
		totalScore += calculateScoreForPart2(inputs[0], inputs[1])
	}

	fmt.Println("Your final score:", totalScore)
}
