package nine

import (
	"aoc"
	"strconv"
	"strings"
)

func solve(file string, f func(a []int) int) int {
	scanner := aoc.OpenScanner(file)
	total := 0
	for scanner.Scan() {
		seq := parse(scanner.Text())
		total += f(seq)
	}
	return total
}

func parse(line string) []int {
	fields := strings.Fields(line)
	result := make([]int, len(fields))
	for i, f := range fields {
		n, err := strconv.Atoi(f)
		if err != nil {
			panic("Failed on " + line)
		}
		result[i] = n
	}
	return result
}

func next(sequence []int) int {
	deltas := findDeltas(sequence)
	last := sequence[len(sequence)-1]
	if allEqual(deltas) {
		return last + deltas[0]
	} else {
		return last + next(deltas)
	}
}

func prev(sequence []int) int {
	deltas := findDeltas(sequence)
	first := sequence[0]
	if allEqual(deltas) {
		return first - deltas[0]
	} else {
		return first - prev(deltas)
	}
}

func allEqual(values []int) bool {
	first := values[0]
	for _, value := range values[1:] {
		if first != value {
			return false
		}
	}
	return true
}

func findDeltas(sequence []int) []int {
	result := make([]int, len(sequence)-1)
	a, b := 0, 1

	for b < len(sequence) {
		d := sequence[b] - sequence[a]
		result[a] = d
		a++
		b++
	}
	return result
}
