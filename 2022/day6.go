package main

import (
	"bufio"
	"fmt"
	"os"
)

type MarkerDetector struct {
	window       map[rune]int
	windowSize   int
	valueByIndex map[int]rune
}

func (m *MarkerDetector) insert(value rune, index int) {
	m.window[value]++
	m.valueByIndex[index] = value

	if index >= m.windowSize {
		m.removeOutOfBoundsValue(index)
	}
}

func (m *MarkerDetector) removeOutOfBoundsValue(newIndex int) {
	index := newIndex - m.windowSize
	valueToRemove := m.valueByIndex[index]

	delete(m.valueByIndex, index)

	m.window[valueToRemove]--
	if m.window[valueToRemove] == 0 {
		delete(m.window, valueToRemove)
	}
}

func (m *MarkerDetector) isMarkerDetected() bool {
	return len(m.window) == m.windowSize
}

func NewMarkerDetector(windowSize int) *MarkerDetector {
	return &MarkerDetector{window: map[rune]int{}, windowSize: windowSize, valueByIndex: map[int]rune{}}
}

func main() {
	// file, err := os.Open("day6_sample_input.txt")
	file, err := os.Open("day6_input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	packetMarkerDetector := NewMarkerDetector(4)
	messageMarkerDetector := NewMarkerDetector(14)

	lengthToPacketMarker := 0
	lengthToMessageMarker := 0

	for i := 0; i < len(line); i++ {
		if !packetMarkerDetector.isMarkerDetected() {
			packetMarkerDetector.insert(rune(line[i]), i)
			lengthToPacketMarker++
		}

		messageMarkerDetector.insert(rune(line[i]), i)
		lengthToMessageMarker++

		if messageMarkerDetector.isMarkerDetected() {
			break
		}
	}

	fmt.Println("Result for part 1:", lengthToPacketMarker)
	fmt.Println("Result for part 2:", lengthToMessageMarker)
}
