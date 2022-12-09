package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("day8_sample_input.txt")
	file, err := os.Open("day8_input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	matrix := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		digits := strings.Split(line, "")

		row := make([]int, len(digits))
		for i, digit := range digits {
			num, err := strconv.Atoi(digit)
			if err != nil {
				panic(err)
			}

			row[i] = num
		}

		matrix = append(matrix, row)
	}

	// fmt.Println(matrix)
	// fmt.Println("Num rows:", len(matrix))
	// fmt.Println("Num cols:", len(matrix[0]))

	numRows := len(matrix)
	numCols := len(matrix[0])

	totalTreesSeen := (numRows+numCols)*2 - 4
	cellsAccountedFor := map[string]bool{}

	maxValuesForCol := make([]int, numCols)
	for i, v := range matrix[0] {
		maxValuesForCol[i] = v
	}

	for i := 1; i < numRows-1; i++ {
		maxValueInRow := matrix[i][0]

		for j := 1; j < numCols-1; j++ {
			val := matrix[i][j]
			key := fmt.Sprintf("%d-%d", i, j)

			if val > maxValueInRow && val > matrix[i][j-1] {
				maxValueInRow = val
				totalTreesSeen++
				cellsAccountedFor[key] = true

				// fmt.Println(key)
			}

			if val > maxValuesForCol[j] {
				if !cellsAccountedFor[key] && val > matrix[i-1][j] {
					totalTreesSeen++
					cellsAccountedFor[key] = true

					// fmt.Println(key)
				}

				maxValuesForCol[j] = val
			}
		}
	}

	// fmt.Println()

	maxValuesForCol = make([]int, numCols)
	for i, v := range matrix[numRows-1] {
		maxValuesForCol[i] = v
	}

	for i := numRows - 2; i >= 1; i-- {
		maxValueInRow := matrix[i][numCols-1]

		for j := numCols - 2; j >= 1; j-- {
			val := matrix[i][j]
			key := fmt.Sprintf("%d-%d", i, j)

			if cellsAccountedFor[key] {
				if val > maxValueInRow {
					maxValueInRow = val
				}

				if val > maxValuesForCol[j] {
					maxValuesForCol[j] = val
				}

				continue
			}

			if val > maxValueInRow {
				if !cellsAccountedFor[key] && val > matrix[i][j+1] {
					totalTreesSeen++
					cellsAccountedFor[key] = true

					// fmt.Println(key)
				}

				maxValueInRow = val
			}

			if val > maxValuesForCol[j] {
				if !cellsAccountedFor[key] && val > matrix[i+1][j] {
					totalTreesSeen++
					cellsAccountedFor[key] = true

					// fmt.Println(key)
				}

				maxValuesForCol[j] = val
			}
		}
	}

	fmt.Println("Total trees seens from the edges", totalTreesSeen)

	maxScenicScore := uint64(0)
	for key, _ := range cellsAccountedFor {
		parts := strings.Split(key, "-")

		i, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		j, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		score := calculateScenicScore(matrix, numRows, numCols, i, j)
		if score > maxScenicScore {
			maxScenicScore = score
		}
	}

	fmt.Println("Max scenic score:", maxScenicScore)
}

func calculateScenicScore(matrix [][]int, numRows int, numCols int, i int, j int) uint64 {
	distanceAbove := 0
	for k := i - 1; k >= 0; k-- {
		distanceAbove++

		if matrix[i][j] <= matrix[k][j] {
			break
		}
	}

	distanceBelow := 0
	for k := i + 1; k < numRows; k++ {
		distanceBelow++

		if matrix[i][j] <= matrix[k][j] {
			break
		}
	}

	distanceLeft := 0
	for k := j - 1; k >= 0; k-- {
		distanceLeft++

		if matrix[i][j] <= matrix[i][k] {
			break
		}
	}

	distanceRight := 0
	for k := j + 1; k < numCols; k++ {
		distanceRight++

		if matrix[i][j] <= matrix[i][k] {
			break
		}
	}

	return uint64(distanceAbove * distanceBelow * distanceLeft * distanceRight)
}
