package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("day11_sample_input.txt")
	file, err := os.Open("day11_input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	monkeys := []*Monkey{}

	monkeyID := 0
	worryLevels := []string{}
	leftOperator := ""
	rightOperator := ""
	operand := ""
	valueForTest := uint64(1)
	monkeyTrue := 0
	monkeyFalse := 0

	monkeyAttrLine := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		switch monkeyAttrLine {
		case 1:
			parts := strings.Split(line, " ")
			id := strings.TrimRight(parts[1], ":")
			monkeyID, err = strconv.Atoi(id)
			if err != nil {
				panic(err)
			}
		case 2:
			parts := strings.Split(line, ":")
			worryLevels = strings.Split(
				strings.TrimLeft(parts[1], " "),
				", ",
			)
		case 3:
			parts := strings.Split(line, "=")
			operators := strings.Split(
				strings.TrimLeft(parts[1], " "),
				" ",
			)
			leftOperator = operators[0]
			operand = operators[1]
			rightOperator = operators[2]
		case 4:
			parts := strings.Split(line, " ")
			valueForTest, err = strconv.ParseUint(parts[len(parts)-1], 10, 64)
			if err != nil {
				panic(err)
			}
		case 5:
			parts := strings.Split(line, " ")
			monkeyTrue, err = strconv.Atoi(parts[len(parts)-1])
			if err != nil {
				panic(err)
			}
		case 6:
			parts := strings.Split(line, " ")
			monkeyFalse, err = strconv.Atoi(parts[len(parts)-1])
			if err != nil {
				panic(err)
			}

			worryLevelsInt := make([]uint64, len(worryLevels))
			for i, level := range worryLevels {
				v, err := strconv.ParseUint(level, 10, 64)
				if err != nil {
					panic(err)
				}

				worryLevelsInt[i] = v
			}

			monkey := &Monkey{
				id:            monkeyID,
				worryLevels:   worryLevelsInt,
				leftOperator:  leftOperator,
				rightOperator: rightOperator,
				worryOperand:  operand,
				valueForTest:  valueForTest,
				monkeyTrue:    monkeyTrue,
				monkeyFalse:   monkeyFalse,
			}
			monkeys = append(monkeys, monkey)
		}

		if monkeyAttrLine <= 6 {
			monkeyAttrLine++
		} else {
			monkeyAttrLine = 1
		}
	}

	// For part 1
	// rounds := 20

	// For part 2
	rounds := 10000

	for i := 1; i <= rounds; i++ {
		for _, m := range monkeys {
			results := m.do(i)
			for k, v := range results {
				monkeys[k].receive(v)
			}
		}

		for _, m := range monkeys {
			fmt.Println(m.worryLevels)
		}
		fmt.Println()
	}

	counts := make([]int, len(monkeys))
	for i, m := range monkeys {
		fmt.Println(m.inspectCount)
		counts[i] = m.inspectCount
	}

	sort.Ints(counts)

	fmt.Println("result:", counts[len(counts)-1]*counts[len(counts)-2])
}

type Monkey struct {
	id            int
	worryLevels   []uint64
	leftOperator  string
	rightOperator string
	worryOperand  string
	valueForTest  uint64
	monkeyTrue    int
	monkeyFalse   int
	inspectCount  int
}

func (m *Monkey) do(round int) map[int][]uint64 {
	results := map[int][]uint64{
		m.monkeyTrue:  []uint64{},
		m.monkeyFalse: []uint64{},
	}

	for _, worry := range m.worryLevels {
		newLevel := m.calculateWorryLevel(worry, round)
		if newLevel%m.valueForTest == 0 {
			results[m.monkeyTrue] = append(results[m.monkeyTrue], newLevel)
		} else {
			results[m.monkeyFalse] = append(results[m.monkeyFalse], newLevel)
		}

		m.inspectCount++
	}

	m.worryLevels = []uint64{}

	return results
}

func (m *Monkey) calculateWorryLevel(worry uint64, round int) uint64 {
	var left, right uint64

	if m.leftOperator == "old" {
		left = worry
	} else {
		l, err := strconv.ParseUint(m.leftOperator, 10, 64)
		if err != nil {
			panic(err)
		}

		left = l
	}

	if m.rightOperator == "old" {
		right = worry
	} else {
		r, err := strconv.ParseUint(m.rightOperator, 10, 64)
		if err != nil {
			panic(err)
		}

		right = r
	}

	var intermediate uint64
	switch m.worryOperand {
	case "*":
		intermediate = left * right
	case "+":
		intermediate = left + right
	}

	// Use this for part 1
	// result := math.Floor(float64(intermediate) / 3.0)

	// Use this for part 2
	result := intermediate % 9699690

	return uint64(result)
}

func (m *Monkey) receive(items []uint64) {
	newLevels := make([]uint64, len(m.worryLevels)+len(items))
	for i, l := range m.worryLevels {
		newLevels[i] = l
	}

	for i, l := range items {
		newLevels[i+len(m.worryLevels)] = l
	}

	m.worryLevels = newLevels
}
