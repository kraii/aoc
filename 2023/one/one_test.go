package one

import (
	"aoc"
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

func TestOne(t *testing.T) {
	f, _ := os.Open(aoc.GetFilePath("one/one.txt"))
	scanner := bufio.NewScanner(bufio.NewReader(f))
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		first := line[strings.IndexFunc(line, unicode.IsNumber)]
		last := line[strings.LastIndexFunc(line, unicode.IsNumber)]

		calibrationValue, _ := strconv.Atoi(string(first) + string(last))
		sum += calibrationValue
	}
	println(sum)
}

func TestDebugOne(t *testing.T) {
	line := "338"
	println(strings.IndexFunc(line, unicode.IsNumber))
}

func TestOnePart2(t *testing.T) {
	f, _ := os.Open(aoc.GetFilePath("one/one.txt"))
	scanner := bufio.NewScanner(bufio.NewReader(f))
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += valueOf(line)
	}
	println(sum)
}

func TestOnePart2Examples(t *testing.T) {
	examples := map[string]int{
		"two1nine":                          29,
		"eightwothree":                      83,
		"abcone2threexyz":                   13,
		"xtwone3four":                       24,
		"4nineeightseven2":                  42,
		"zoneight234":                       14,
		"7pqrstsixteen":                     76,
		"123":                               13,
		"sixdddkcqjdnzzrgfourxjtwosevenhg9": 69,
		"pcf2":                              22,
		"eightpkhgpcnc8eightfive1hdtcjjdcsevennpz": 87,
	}
	for k, v := range examples {
		assert.Equal(t, v, valueOf(k), fmt.Sprintf("%s should give %d", k, v))
	}
}

var digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func valueOf(s string) int {
	var results []string
	for i, r := range s {
		if unicode.IsDigit(r) {
			results = append(results, string(r))
		} else {
			for d, digit := range digits {
				if strings.HasPrefix(s[i:], digit) {
					results = append(results, strconv.Itoa(d+1))
				}
			}
		}
	}

	result, err := strconv.Atoi(results[0] + results[len(results)-1])

	if err != nil {
		panic(err)
	}
	return result
}
