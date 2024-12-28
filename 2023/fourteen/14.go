package fourteen

import (
	"aoc"
	"strings"
)

type platform [][]rune

func (p platform) print() {
	println(p.toString())
	println()
}

func (p platform) toString() string {
	pStrings := make([]string, len(p))
	for i, runes := range p {
		pStrings[i] = string(runes)
	}
	return strings.Join(pStrings, "\n")
}

func solvePart1(file string) int {
	plat := parse(file)
	tiltNegative(plat, north)

	return countLoad(plat)
}

func solvePart2(file string, maxCycles int) int {
	plat := parse(file)
	spin(plat, maxCycles)

	return countLoad(plat)
}

func spin(plat platform, maxCycles int) {
	println("spin", maxCycles)
	seen := make(aoc.Set[string])
	cycleEnd := ""
	cycleLength := 0
	for i := 0; i < maxCycles; i++ {
		tiltNegative(plat, north)
		tiltNegative(plat, west)
		tiltPositive(plat, south)
		tiltPositive(plat, east)
		afterSpin := plat.toString()

		// Cycle detected
		if seen.Contains(afterSpin) {
			if cycleEnd == afterSpin {
				// I am not sure why the -1. But it worked?
				spin(plat, ((maxCycles-i)%cycleLength)-1)
				return
			}
			if cycleLength == 0 {
				cycleEnd = afterSpin
			}
			cycleLength++
		}
		seen.Add(afterSpin)
	}

}

func countLoad(plat platform) int {
	total := 0
	for y, row := range plat {
		for _, r := range row {
			if r == 'O' {
				total += len(plat) - y
			}
		}
	}
	return total
}

func north(x int, y int) (int, int) {
	return x, y - 1
}

func west(x int, y int) (int, int) {
	return x - 1, y
}

func south(x int, y int) (int, int) {
	return x, y + 1
}

func east(x int, y int) (int, int) {
	return x + 1, y
}

type roller func(x int, y int) (int, int)

func tiltNegative(p platform, next roller) {
	for y, row := range p {
		for x, row := range row {
			if row == 'O' {
				rollRock(p, x, y, next)
			}
		}
	}
}

func tiltPositive(p platform, next roller) {
	for y := len(p) - 1; y >= 0; y-- {
		for x := len(p[y]) - 1; x >= 0; x-- {
			if p[y][x] == 'O' {
				rollRock(p, x, y, next)
			}
		}
	}
}

func rollRock(p platform, startX int, startY int, next roller) {
	prevX, prevY := startX, startY
	x, y := next(startX, startY)
	for x >= 0 && y >= 0 && x < len(p[0]) && y < len(p) {
		titleBelow := p[y][x]
		// can the rock keep rolling in this direction?
		if titleBelow != '.' {
			// Write . first so that if we haven't moved anywhere there's no problem
			p[startY][startX] = '.'
			p[prevY][prevX] = 'O'
			return
		}
		prevX, prevY = x, y
		x, y = next(x, y)
	}
	// we got to the bottom
	p[startY][startX] = '.'
	p[prevY][prevX] = 'O'
}

func parse(file string) platform {
	scanner := aoc.OpenScanner(file)
	var result platform

	for scanner.Scan() {
		result = append(result, []rune(scanner.Text()))
	}
	return result
}
