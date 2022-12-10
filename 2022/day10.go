package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cpu struct {
	registerValue       int64
	cycle               uint64
	cycleToCollectData  uint64
	totalSignalStrength int64
	tube                *CarthodeTube
}

func (c *Cpu) process(instruction string) {
	parts := strings.Split(instruction, " ")
	command := parts[0]

	var arg int64
	if len(parts) == 2 {
		arg, _ = strconv.ParseInt(parts[1], 10, 64)
	}

	switch command {
	case "addx":
		c.evaluateCycle()
		c.evaluateCycle()

		c.registerValue += arg
	case "noop":
		c.evaluateCycle()
	}
}

func (c *Cpu) evaluateCycle() {
	c.collectSignalStrength()
	c.tube.draw(c.registerValue)
	c.cycle++
}

func (c *Cpu) collectSignalStrength() {
	if c.cycle == c.cycleToCollectData && c.cycle <= 220 {
		c.totalSignalStrength += int64(c.cycle) * c.registerValue

		c.cycleToCollectData += 40
	}
}

type CarthodeTube struct {
	screen [6][40]string
	row    int64
	col    int64
}

func (c *CarthodeTube) draw(spriteMidIndex int64) {
	if c.col == spriteMidIndex-1 || c.col == spriteMidIndex || c.col == spriteMidIndex+1 {
		c.screen[c.row][c.col] = "#"
	} else {
		c.screen[c.row][c.col] = "."
	}

	c.col++
	if c.col == int64(len(c.screen[0])) {
		c.row++
		c.col = 0
	}

}

func main() {
	// file, err := os.Open("day10_sample_input.txt")
	file, err := os.Open("day10_input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	cpu := Cpu{
		registerValue:      1,
		cycle:              1,
		cycleToCollectData: 20,
		tube:               &CarthodeTube{},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cpu.process(scanner.Text())
	}

	fmt.Println("Part 1:", cpu.totalSignalStrength)

	fmt.Println("Part 2:")
	for _, row := range cpu.tube.screen {
		fmt.Println(row)
	}
}
