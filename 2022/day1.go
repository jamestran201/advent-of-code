package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type caloriesCounterInterface interface {
	add(uint64)
	postProcess()
	result() uint64
}

type MaxCaloriesCounter struct {
	currentCalories uint64
	maxCalories     uint64
}

func (mc *MaxCaloriesCounter) add(calories uint64) {
	mc.currentCalories += calories
}

func (mc *MaxCaloriesCounter) postProcess() {
	if mc.currentCalories > mc.maxCalories {
		mc.maxCalories = mc.currentCalories
	}

	mc.currentCalories = 0
}

func (mc *MaxCaloriesCounter) result() uint64 {
	return mc.maxCalories
}

type Top3CaloriesCounter struct {
	currentCalories uint64
	top3Calories    [3]uint64
}

func (tc *Top3CaloriesCounter) add(calories uint64) {
	tc.currentCalories += calories
}

func (tc *Top3CaloriesCounter) postProcess() {
	index, minCalories := tc.findMinCalories()
	if tc.currentCalories > minCalories {
		tc.top3Calories[index] = tc.currentCalories
	}

	tc.currentCalories = 0
}

func (tc *Top3CaloriesCounter) result() uint64 {
	return tc.top3Calories[0] + tc.top3Calories[1] + tc.top3Calories[2]
}

func (tc *Top3CaloriesCounter) findMinCalories() (int, uint64) {
	index := 0
	minCalories := tc.top3Calories[0]
	for i, calories := range tc.top3Calories {
		if calories < minCalories {
			index = i
			minCalories = calories
		}
	}

	return index, minCalories
}

type CaloriesCounter struct {
	specializedCounter caloriesCounterInterface
}

func (c *CaloriesCounter) Process(input string) {
	if c.isDoneCountingForAnElf(input) {
		c.specializedCounter.postProcess()
	} else {
		calories, err := strconv.ParseUint(input, 10, 64)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		c.specializedCounter.add(calories)
	}
}

func (c *CaloriesCounter) Result() uint64 {
	return c.specializedCounter.result()
}

func (c *CaloriesCounter) isDoneCountingForAnElf(s string) bool {
	return s == ""
}

func part1() {
	file, err := os.Open("day1_input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	caloriesCounter := &CaloriesCounter{specializedCounter: &MaxCaloriesCounter{}}
	for scanner.Scan() {
		caloriesCounter.Process(scanner.Text())
	}

	fmt.Println("Max calories:", caloriesCounter.Result())
}

func part2() {
	file, err := os.Open("day1_input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	caloriesCounter := &CaloriesCounter{specializedCounter: &Top3CaloriesCounter{}}
	for scanner.Scan() {
		caloriesCounter.Process(scanner.Text())
	}

	fmt.Println("Max calories:", caloriesCounter.Result())
}

func main() {
	puzzle_part := 1 // either 1 or 2

	file, err := os.Open("day1_input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	var caloriesCounter *CaloriesCounter
	if puzzle_part == 1 {
		caloriesCounter = &CaloriesCounter{specializedCounter: &MaxCaloriesCounter{}}
	} else {
		caloriesCounter = &CaloriesCounter{specializedCounter: &Top3CaloriesCounter{}}
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		caloriesCounter.Process(scanner.Text())
	}

	fmt.Println("Max calories:", caloriesCounter.Result())
}
