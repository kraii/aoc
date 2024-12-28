package three

import (
	"strconv"
	"unicode"
)

func solveB(file string) int {
	lines := allLines(file)

	total := 0
	for i, line := range lines {
		for j, c := range line {
			if c == '*' {
				total += calculateRatio(lines, i, j)
			}
		}
	}
	return total
}

func calculateRatio(lines []string, x int, y int) int {
	found := findAdjacentNumbers(lines, x, y)
	if len(found) == 2 {
		return found[0] * found[1]
	} else {
		return 0
	}
}

func findAdjacentNumbers(lines []string, x int, y int) []int {
	var found []int

	maxI := len(lines) - 1
	maxJ := len(lines[0]) - 1

	fromI, toI := clamp(x-1, maxI), clamp(x+1, maxI)
	fromJ, toJ := clamp(y-1, maxJ), clamp(y+1, maxJ)

	for i := fromI; i <= toI; i++ {
		for j := fromJ; j <= toJ; j++ {
			char := lines[i][j]
			if unicode.IsDigit(rune(char)) {
				n, newJ := getNumber(lines[i], j)
				found = append(found, n)
				// skip included indices to not include a number twice
				j = newJ
			}
		}
	}

	return found
}

func getNumber(line string, startIndex int) (int, int) {
	acc := ""
	// Search backwards
	i := startIndex - 1
	for i >= 0 && unicode.IsDigit(rune(line[i])) {
		acc = string(line[i]) + acc
		i--
	}
	// search forwards
	i = startIndex
	for i < len(line) && unicode.IsDigit(rune(line[i])) {
		acc += string(line[i])
		i++
	}
	result, err := strconv.Atoi(acc)
	if err != nil {
		panic(err)
	}
	return result, i
}
