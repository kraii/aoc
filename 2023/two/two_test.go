package two

import (
	"aoc"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

var colours = map[string]int{"red": 12, "blue": 14, "green": 13}

func TestTwoA(t *testing.T) {
	scanner := aoc.OpenScanner("two/two-input.txt")
	total := 0
	gameId := 0
	for scanner.Scan() {
		gameId++
		line := scanner.Text()
		split := strings.Split(line, ":")
		revealsStr := split[1]
		reveals := strings.Split(revealsStr, ";")

		valid := true
		for _, reveal := range reveals {
			result := parseReveals(reveal)
			if !isValid(result) {
				valid = false
				break
			}
		}
		if valid {
			total += gameId
		}
	}
	println(total)

}

func TestTwoB(t *testing.T) {
	scanner := aoc.OpenScanner("two/two-input.txt")
	total := 0
	gameId := 0
	for scanner.Scan() {
		gameId++
		line := scanner.Text()
		split := strings.Split(line, ":")
		revealsStr := split[1]
		reveals := strings.Split(revealsStr, ";")

		minViable := make(map[string]int, 4)
		for _, reveal := range reveals {
			result := parseReveals(reveal)
			for colour, count := range result {
				if count > minViable[colour] {
					minViable[colour] = count
				}
			}
		}
		total += product(minViable)
	}
	println(total)

}

func isValid(result map[string]int) bool {
	for colour, available := range colours {
		if result[colour] > available {
			return false
		}
	}
	return true
}

func product(counts map[string]int) int {
	result := 0
	for colour, _ := range colours {
		if 0 == result {
			result += counts[colour]
		} else {
			result = result * counts[colour]
		}
	}
	return result
}

func parseReveals(reveals string) map[string]int {
	split := strings.Split(reveals, ",")
	result := make(map[string]int, 4)

	for _, s := range split {
		count, colour := parseReveal(s)
		result[colour] = count
	}
	return result
}

func parseReveal(reveal string) (int, string) {
	split := strings.Fields(reveal)
	count, err := strconv.Atoi(trim(split[0]))
	if err != nil {
		fmt.Printf("%v\n", split)
		fmt.Printf("%s | %s\n", reveal, trim(split[0]))
		panic(err)
	}
	return count, trim(split[1])
}

func trim(s string) string {
	return strings.Trim(s, " ,:")
}
