package three

import (
	"aoc"
	"strconv"
	"unicode"
)

func solveA(file string) int {
	lines := allLines(file)

	total := 0
	var currentJs []int
	for i, line := range lines {
		lineLength := len(line)
		for j := 0; j < lineLength; j++ {
			if unicode.IsDigit(rune(line[j])) {
				currentJs = append(currentJs, j)
			} else {
				if len(currentJs) > 0 {
					if isPartNumber(lines, i, currentJs, lineLength) {
						total += value(line, currentJs)
					}
				}
				currentJs = []int{}
			}
		}
		if len(currentJs) > 0 {
			if isPartNumber(lines, i, currentJs, lineLength) {
				total += value(line, currentJs)
			}
		}
	}

	return total
}

func value(line string, js []int) int {
	v := toString(line, js)
	result, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return result
}

func toString(line string, js []int) string {
	v := ""
	for _, j := range js {
		v += string(line[j])
	}
	return v
}

func isPartNumber(lines []string, x int, ys []int, lineLength int) bool {
	maxI := len(lines) - 1
	maxJ := lineLength - 1

	fromI, toI := clamp(x-1, maxI), clamp(x+1, maxI)
	fromJ, toJ := clamp(ys[0]-1, maxJ), clamp(ys[len(ys)-1]+1, maxJ)

	for i := fromI; i <= toI; i++ {
		for j := fromJ; j <= toJ; j++ {
			char := lines[i][j]
			if isSymbol(char) {
				return true
			}
		}
	}

	return false
}

var symbols = map[uint8]bool{
	'*': true,
	'$': true,
	'+': true,
	'#': true,
	'=': true,
	'/': true,
	'%': true,
	'@': true,
	'-': true,
	'&': true,
}

func isSymbol(char uint8) bool {
	_, present := symbols[char]
	return present
}

func clamp(x int, upper int) int {
	if x < 0 {
		return 0
	}
	if x > upper {
		return upper
	}
	return x
}

func allLines(file string) []string {
	scanner := aoc.OpenScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
