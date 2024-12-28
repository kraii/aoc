package five

import (
	"aoc"
	"fmt"
	"math"
	"runtime"
	"strconv"
	"strings"
	"unicode"
)

type gardenRange struct {
	sourceRangeStart int
	sourceRangeEnd   int
	destRangeStart   int
}

type mapping = []*gardenRange

type almanac struct {
	seeds    []int
	mappings []mapping
}

func solveA(file string) int {
	a := readAlmanac(file)
	closestLoc := math.MaxInt
	for _, seed := range a.seeds {
		loc := findLocation(seed, a)
		if loc < closestLoc {
			closestLoc = loc
		}
	}
	return closestLoc
}

// Brute force awayyyyyyy
func solveBWithBruteForce(file string) int {
	a := readAlmanac(file)

	jobs := make(chan []int, runtime.NumCPU())
	results := make(chan int, runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		go seedRangeWorker(a, jobs, results)
	}

	for i := 0; i < len(a.seeds); i += 2 {
		jobs <- a.seeds[i : i+2]
	}
	close(jobs)
	closestLoc := math.MaxInt
	for r := 0; r < len(a.seeds)/2; r++ {
		rangeClosest := <-results
		if rangeClosest < closestLoc {
			closestLoc = rangeClosest
		}
	}

	return closestLoc
}

func seedRangeWorker(a *almanac, input <-chan []int, results chan<- int) {
	for job := range input {

		fmt.Printf("Starting job %v\n", job)
		closestLoc := math.MaxInt
		rangeStart := job[0]
		rangeEnd := rangeStart + job[1]

		for seed := rangeStart; seed <= rangeEnd; seed++ {
			loc := findLocation(seed, a)
			if loc < closestLoc {
				closestLoc = loc
			}
		}

		results <- closestLoc
	}

}

func findLocation(seed int, a *almanac) int {
	current := seed
	for _, mapping := range a.mappings {
		for _, r := range mapping {
			if r.contains(current) {
				current = r.convert(current)
				break
			}
		}
	}
	return current
}

func (r *gardenRange) contains(input int) bool {
	return r.sourceRangeStart <= input && input <= r.sourceRangeEnd
}

func (r *gardenRange) convert(input int) int {
	return input - r.sourceRangeStart + r.destRangeStart
}

func readAlmanac(file string) *almanac {
	var seeds []int
	var mappings [][]*gardenRange

	scanner := aoc.OpenScanner(file)
	var current []*gardenRange
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "seeds:") {
			seedsStr := strings.Split(line, ":")[1]
			seeds = parse(seedsStr)
		} else if strings.Contains(line, "map:") {
			current = []*gardenRange{}
		} else if len(current) > 0 && len(strings.TrimSpace(line)) == 0 {
			mappings = append(mappings, current)
		} else if strings.ContainsFunc(line, unicode.IsDigit) {
			current = append(current, parseMapping(line))
		}
	}

	return &almanac{
		seeds:    seeds,
		mappings: mappings,
	}
}

func parseMapping(line string) *gardenRange {
	values := parse(line)
	rangeLength := values[2]
	return &gardenRange{
		sourceRangeStart: values[1],
		destRangeStart:   values[0],
		sourceRangeEnd:   values[1] + rangeLength,
	}
}

func parse(line string) []int {
	fields := strings.Fields(line)
	values := make([]int, len(fields))
	for i, v := range fields {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		values[i] = n
	}
	return values
}

func solveBWithBinarySearch(file string) int {
	a := readAlmanac(file)

	closestLoc := math.MaxInt
	for i := 0; i < len(a.seeds); i += 2 {
		start := a.seeds[i]
		length := a.seeds[i+1]
		closestLoc = min(closestLoc, lowestLocForSeedRange(a, start, length))
	}

	return closestLoc
}

// binary search ish
func lowestLocForSeedRange(a *almanac, start int, length int) int {
	if length < 1 {
		panic("I dun goofed")
	}
	if length == 1 {
		return min(findLocation(start, a), findLocation(start+1, a))
	}

	halfRange := length / 2
	pivot := start + halfRange
	end := start + length

	locStart := findLocation(start, a)
	locMiddle := findLocation(pivot, a)
	locEnd := findLocation(end, a)

	// if range is uniform we don't need to continue searching in one half of the range
	// we just take the mapped value from min seed of the range
	closestLoc := locStart

	if locMiddle != locStart+halfRange {
		closestLoc = min(closestLoc, lowestLocForSeedRange(a, start, halfRange))
	}
	if locEnd != locMiddle+halfRange {
		closestLoc = min(closestLoc, lowestLocForSeedRange(a, pivot+1, halfRange))
	}
	return closestLoc
}
