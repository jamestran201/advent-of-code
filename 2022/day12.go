package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("day12_sample_input.txt")
	file, err := os.Open("day12_input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	grid := [][]uint8{}

	startX := -1
	startY := -1

	endX := -1
	endY := -1

	rowNum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		row := make([]uint8, len(line))
		for i := 0; i < len(line); i++ {
			row[i] = line[i]

			if string(row[i]) == "S" {
				startX = rowNum
				startY = i
			}

			if string(row[i]) == "E" {
				endX = rowNum
				endY = i
			}
		}

		rowNum++
		grid = append(grid, row)
	}

	day12Part1(grid, startX, startY, endX, endY)
	day12Part2(grid, endX, endY)
}

func day12Part1(grid [][]uint8, startX int, startY int, endX int, endY int) {
	result := aStar(grid, startX, startY, endX, endY)
	fmt.Println(result)
}

func day12Part2(grid [][]uint8, endX int, endY int) {
	result := uint64(0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			val := string(grid[i][j])
			if val == "S" || val == "a" {
				cur := aStar(grid, i, j, endX, endY)
				if result == 0 || (cur != 0 && cur < result) {
					result = cur
				}
			}
		}
	}

	fmt.Println(result)
}

func aStar(grid [][]uint8, startX int, startY int, endX int, endY int) uint64 {
	openSet := map[string]bool{}
	closedSet := map[string]bool{}
	distFromStart := map[string]uint64{}
	parents := map[string]string{}

	key := fmt.Sprintf("%d-%d", startX, startY)
	distFromStart[key] = 0
	parents[key] = key
	openSet[key] = true

	endKey := fmt.Sprintf("%d-%d", endX, endY)
	for len(openSet) > 0 {
		node := ""

		for k, _ := range openSet {
			if node == "" || (distFromStart[k]+heuristic(k, endX, endY) < distFromStart[node]+heuristic(node, endX, endY)) {
				node = k
			}
		}

		neighbors := getNeighbors(grid, node)
		if node != endKey && len(neighbors) > 0 {
			for _, k := range neighbors {
				if !openSet[k] && !closedSet[k] {
					openSet[k] = true
					parents[k] = node
					distFromStart[k] = distFromStart[node] + 1
				} else {
					if distFromStart[k] > distFromStart[node]+1 {
						distFromStart[k] = distFromStart[node] + 1
						parents[k] = node

						if closedSet[k] {
							delete(closedSet, k)
							openSet[k] = true
						}
					}
				}
			}
		}

		if node == "" {
			fmt.Println("SOMETHING IS WRONG!!")
			os.Exit(1)
		}

		if node == endKey {
			res := 0
			for parents[node] != node {
				node = parents[node]
				res++
			}

			return uint64(res)
		}

		delete(openSet, node)
		closedSet[node] = true
	}

	return 0
}

func heuristic(key string, x int, y int) uint64 {
	parts := strings.Split(key, "-")
	curX, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		panic(err)
	}

	curY, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		panic(err)
	}

	return uint64(math.Abs(curX-float64(x)) + math.Abs(curY-float64(y)))
}

func getNeighbors(grid [][]uint8, node string) []string {
	parts := strings.Split(node, "-")
	curX, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	curY, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	results := []string{}

	up := curX - 1
	down := curX + 1
	left := curY - 1
	right := curY + 1

	curHeight := grid[curX][curY]
	if string(curHeight) == "S" {
		curHeight = 97
	}

	if string(curHeight) == "E" {
		curHeight = 122
	}

	if up >= 0 && isValidMove(grid, up, curY, curHeight) {
		key := fmt.Sprintf("%d-%d", up, curY)
		results = append(results, key)
	}

	if down < len(grid) && isValidMove(grid, down, curY, curHeight) {
		key := fmt.Sprintf("%d-%d", down, curY)
		results = append(results, key)
	}

	if left >= 0 && isValidMove(grid, curX, left, curHeight) {
		key := fmt.Sprintf("%d-%d", curX, left)
		results = append(results, key)
	}

	if right < len(grid[0]) && isValidMove(grid, curX, right, curHeight) {
		key := fmt.Sprintf("%d-%d", curX, right)
		results = append(results, key)
	}

	return results
}

func isValidMove(grid [][]uint8, nX int, nY int, curHeight uint8) bool {
	newHeight := grid[nX][nY]
	if string(newHeight) == "S" {
		newHeight = 97
	}

	if string(newHeight) == "E" {
		newHeight = 122
	}

	return newHeight == curHeight+1 || newHeight <= curHeight
}
