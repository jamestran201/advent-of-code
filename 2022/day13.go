package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// file, err := os.Open("day13_sample_input.txt")
	file, err := os.Open("day13_input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	resultPart1 := 0
	packets := [2]string{}
	pairIndex := 1
	i := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		packets[i] = line
		i++

		if i != 2 {
			continue
		}

		result := compareLists(packets[0], packets[1])

		fmt.Println(result)
		fmt.Println()

		if result == "ordered" {
			resultPart1 += pairIndex
		}

		i = 0
		pairIndex++
	}

	fmt.Println("Result", resultPart1)
}

func compareLists(lp string, rp string) string {
	fmt.Println(lp)
	fmt.Println(rp)

	lp = lp[1 : len(lp)-1]
	rp = rp[1 : len(rp)-1]

	lSize := len(lp)
	rSize := len(rp)

	li := 0
	ri := 0

	result := "tie"

	for li < lSize && ri < rSize {
		lv, tli := getValue(lp, li)
		rv, tri := getValue(rp, ri)

		if lv == "[" {
			lv, tli = parseList(lp, tli)
		}

		if rv == "[" {
			rv, tri = parseList(rp, tri)
		}

		if isDigit(lv) && isDigit(rv) {
			// debug
			fmt.Println(lv)
			fmt.Println(rv)
			fmt.Println()

			lNum, err := strconv.Atoi(lv)
			if err != nil {
				panic(err)
			}

			rNum, err := strconv.Atoi(rv)
			if err != nil {
				panic(err)
			}

			if lNum < rNum {
				result = "ordered"
				break
			} else if lNum > rNum {
				result = "unordered"
				break
			}
		} else if string(lv[0]) == "[" && string(rv[0]) == "[" {
			fmt.Println(lv)
			fmt.Println(rv)
			fmt.Println()

			result = compareLists(lv, rv)

			if result != "tie" {
				break
			}
		} else if isDigit(rv) && string(lv[0]) == "[" {
			fmt.Println(lv)
			fmt.Println(rv)
			fmt.Println()

			result = compareLists(
				lv,
				fmt.Sprintf("[%s]", rv),
			)

			if result != "tie" {
				break
			}
		} else if isDigit(lv) && string(rv[0]) == "[" {
			fmt.Println(lv)
			fmt.Println(rv)
			fmt.Println()

			result = compareLists(
				fmt.Sprintf("[%s]", lv),
				rv,
			)

			if result != "tie" {
				break
			}
		}

		li = tli
		ri = tri
	}

	if result == "tie" {
		fmt.Println(li)
		fmt.Println(ri)
		if li == len(lp) && ri < len(rp) {
			result = "ordered"
		} else if ri == len(rp) && li < len(lp) {
			result = "unordered"
		}
	}

	fmt.Println(result)
	return result
}

func parseList(p string, start int) (string, int) {
	depth := 1
	end := start
	for depth > 0 && end < len(p) {
		if string(p[end]) == "[" {
			depth++
		} else if string(p[end]) == "]" {
			depth--
		}

		end++
	}

	return p[start-1 : end], end
}

func getValue(p string, i int) (string, int) {
	v := string(p[i])
	if (i+1 == len(p)) || v == "[" || v == "]" || v == "," {
		// return v, i
		return v, i + 1
	}

	end := i + 1
	for end < len(p) {
		v = string(p[end])
		if v >= "0" && v <= "9" {
			end++
		} else {
			break
		}
	}

	return p[i:end], end
}

func isDigit(s string) bool {
	firstChar := string(s[0])
	return firstChar != "[" && firstChar != "]" && firstChar != ","
}
