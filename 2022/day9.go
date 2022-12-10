package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("day9_sample_input.txt")
	file, err := os.Open("day9_input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	visited := map[string]bool{"0-0": true}

	// for part 1
	// positions := [][]int64{
	// 	[]int64{0, 0},
	// 	[]int64{0, 0},
	// }

	// for part 2
	positions := [][]int64{
		[]int64{0, 0},
		[]int64{0, 0},
		[]int64{0, 0},
		[]int64{0, 0},
		[]int64{0, 0},
		[]int64{0, 0},
		[]int64{0, 0},
		[]int64{0, 0},
		[]int64{0, 0},
		[]int64{0, 0},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		direction := parts[0]
		steps, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		for i := steps; i > 0; i-- {
			// Move the head
			switch direction {
			case "U":
				positions[0][0]--
			case "D":
				positions[0][0]++
			case "L":
				positions[0][1]--
			case "R":
				positions[0][1]++
			}

			// See if the subsequent nodes need to be moved
			for j := 1; j < len(positions); j++ {
				before := positions[j-1]
				current := positions[j]

				if current[0] == before[0] && current[1] == before[1] {
					continue
				}

				if before[0] == current[0] { // if same row
					// If adjacent
					if before[1] == current[1]-1 || before[1] == current[1]+1 {
						continue
					}

					if before[1]-current[1] < 0 {
						current[1]--
					} else {
						current[1]++
					}

					if j == len(positions)-1 {
						key := fmt.Sprintf("%d-%d", current[0], current[1])
						visited[key] = true
					}
				} else if before[1] == current[1] { // if same col
					// If adjacent
					if before[0] == current[0]-1 || before[0] == current[0]+1 {
						continue
					}

					if before[0]-current[0] < 0 {
						current[0]--
					} else {
						current[0]++
					}

					if j == len(positions)-1 {
						key := fmt.Sprintf("%d-%d", current[0], current[1])
						visited[key] = true
					}
				} else { // diagonal
					// diagonally adjacent
					if before[0]+1 == current[0] && before[1]+1 == current[1] {
						continue
					} else if before[0]+1 == current[0] && before[1]-1 == current[1] {
						continue
					} else if before[0]-1 == current[0] && before[1]+1 == current[1] {
						continue
					} else if before[0]-1 == current[0] && before[1]-1 == current[1] {
						continue
					}

					// the nodes are not adjacent diagonally, so we need to move the current node in a diagonal
					// this part determines whether the node should be moved to the top left, top right, bottom left
					// or bottom right
					moveTop := (before[0] - current[0]) < 0
					moveLeft := (before[1] - current[1]) < 0
					if moveTop {
						current[0]--
					} else {
						current[0]++
					}

					if moveLeft {
						current[1]--
					} else {
						current[1]++
					}

					if j == len(positions)-1 {
						key := fmt.Sprintf("%d-%d", current[0], current[1])
						visited[key] = true
					}
				}
			}
		}
	}

	fmt.Println("Number of positions visited by tail:", len(visited))
}
