package twelve

import (
	"aoc"
	"strconv"
	"strings"
)

type cacheKey struct {
	row           string
	damagedGroups string
}

var cache = make(map[cacheKey]int)

func solve(file string, copies int) int {
	total := 0
	scanner := aoc.OpenScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		unfoldedRecord := unfold(parts[0], "?", copies)
		unfoldedCounts := unfold(parts[1], ",", copies)
		total += countPossible(unfoldedRecord, parse(unfoldedCounts))
	}
	return total
}

func unfold(value string, sep string, times int) string {
	joins := make([]string, times)
	for i := 0; i < times; i++ {
		joins[i] = value
	}
	return strings.Join(joins, sep)
}

func parse(s string) []int {
	var result []int
	for _, g := range strings.Split(s, ",") {
		r, err := strconv.Atoi(g)
		if err != nil {
			panic("Trying to convert to int " + g)
		}
		result = append(result, r)
	}
	return result
}

func countPossible(row string, damagedGroups []int) int {
	if len(row) == 0 {
		if len(damagedGroups) == 0 {
			return 1
		}
		return 0
	}
	if len(damagedGroups) == 0 {
		// If there's no damaged left then success
		if !strings.ContainsRune(row, '#') {
			return 1
		} else {
			return 0
		}
	}
	key := toCacheKey(row, damagedGroups)
	if v, ok := cache[key]; ok {
		return v
	}

	group := damagedGroups[0]
	if len(row) < group {
		return 0
	}

	result := 0
	rest := row[1:]
	switch row[0] {
	case '.':
		result = countPossible(rest, damagedGroups)
	case '?':
		result = countPossible("#"+rest, damagedGroups) + countPossible("."+rest, damagedGroups)
	case '#':
		if noWorkingSprings(row, group) && chainIsNotTooLong(row, group) {
			if len(row) == group {
				result = countPossible("", damagedGroups[1:])
			} else {
				// we need to strip off an extra character as there needs to be a break between the contiguous groups
				result = countPossible(row[group+1:], damagedGroups[1:])
			}

		}
	}
	cache[key] = result
	return result
}

func toCacheKey(row string, groups []int) cacheKey {
	groupsStr := make([]string, len(groups))
	for i, group := range groups {
		groupsStr[i] = strconv.Itoa(group)
	}
	return cacheKey{row, strings.Join(groupsStr, "-")}
}

func noWorkingSprings(row string, group int) bool {
	return !strings.ContainsRune(row[:group], '.')
}

func chainIsNotTooLong(row string, group int) bool {
	return len(row) == group || row[group] != '#'
}
