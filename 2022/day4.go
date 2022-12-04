package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Section struct {
	start uint64
	end   uint64
}

func (s *Section) Contains(otherSection Section) bool {
	return s.start <= otherSection.start && otherSection.end <= s.end
}

func (s *Section) Overlaps(otherSection Section) bool {
	return (s.start <= otherSection.start && otherSection.start <= s.end) ||
		(otherSection.start <= s.start && s.start <= otherSection.end)
}

func NewSection(input string) Section {
	bounds := strings.Split(input, "-")

	start, err := strconv.ParseUint(bounds[0], 10, 64)
	if err != nil {
		panic(err)
	}

	end, err := strconv.ParseUint(bounds[1], 10, 64)
	if err != nil {
		panic(err)
	}

	return Section{start: start, end: end}
}

func main() {
	// file, err := os.Open("day4_sample_input.txt")
	file, err := os.Open("day4_input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	fullyOverlappedPairsCount := 0
	overlapCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.Split(line, ",")

		section1 := NewSection(pair[0])
		section2 := NewSection(pair[1])
		if section1.Contains(section2) || section2.Contains(section1) {
			fullyOverlappedPairsCount += 1
			overlapCount += 1
		} else if section1.Overlaps(section2) {
			overlapCount += 1
		}
	}

	fmt.Println("Fully overlapped pairs count:", fullyOverlappedPairsCount)
	fmt.Println("Overlapped pairs count:", overlapCount)
}
