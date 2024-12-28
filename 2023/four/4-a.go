package four

import (
	"aoc"
	"slices"
	"strings"
)

func solveA(file string) int {
	scanner := aoc.OpenScanner(file)

	total := 0

	for scanner.Scan() {
		total += solveLine(scanner.Text())
	}

	return total
}

func solveLine(line string) int {
	score := 0

	withoutTitle := strings.Split(line, ":")[1]
	split := strings.Split(withoutTitle, "|")

	winningNumbers := strings.Fields(split[0])
	cardNumbers := strings.Fields(split[1])

	for _, n := range cardNumbers {
		if slices.Contains(winningNumbers, n) {
			if score == 0 {
				score++
			} else {
				score = score * 2
			}
		}
	}

	return score
}
