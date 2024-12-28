package thirteen

import (
	"aoc"
	"strings"
)

type valley struct {
	horizontal []string
	vertical   []string
}

func solve(file string, smudges int) int {
	valleys := readFile(file)
	total := 0
	for _, v := range valleys {
		total += countBeforeReflection(v.vertical, smudges)
		total += 100 * countBeforeReflection(v.horizontal, smudges)
	}
	return total
}

var aNonMatchingString = "aaaaaaaaaaaaaaaaaaaaa"

func countBeforeReflection(data []string, smudges int) int {
	prev := aNonMatchingString

	for i, current := range data {
		// Maybe a reflection check either side of it for equality
		diff := difference(current, prev)
		if diff <= smudges {
			a, b := i-2, i+1
			totalSmudge := diff
			for a >= 0 && b < len(data) {
				totalSmudge += difference(data[a], data[b])
				a--
				b++
			}
			if totalSmudge == smudges {
				return i
			}
		}
		prev = current
	}

	return 0
}

func difference(a string, b string) int {
	differences := 0
	for i, _ := range a {
		if a[i] != b[i] {
			differences++
		}
	}
	return differences
}

func readFile(file string) []*valley {
	scanner := aoc.OpenScanner(file)
	var valleys []*valley
	horizontal := make([]string, 0, 20)
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			valleys = append(valleys, &valley{horizontal: horizontal})
			horizontal = make([]string, 0, 20)
		} else {
			horizontal = append(horizontal, line)
		}
	}

	for _, v := range valleys {
		v.vertical = createColumns(v.horizontal)
	}

	return valleys
}

func createColumns(horizontal []string) []string {
	result := make([]string, len(horizontal[0]))

	for i := 0; i < len(horizontal[0]); i++ {
		column := ""
		for j, _ := range horizontal {
			column += string(horizontal[j][i])
		}
		result[i] = column
	}

	return result
}
