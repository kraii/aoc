package eleven

import (
	"aoc"
)

type point struct {
	x, y int
}

type galaxyPair [2]point

func solve(file string, emptyGrowthFactor int) int {
	galaxies := findGalaxies(file, emptyGrowthFactor)
	pairs := findPairs(galaxies)

	total := 0
	for _, pair := range pairs {
		a, b := pair[0], pair[1]
		distance := distanceBetween(b, a)
		total += distance
	}
	return total
}

// Manhattan distance
func distanceBetween(b point, a point) int {
	deltaX := abs(b.x - a.x)
	deltaY := abs(b.y - a.y)

	return deltaX + deltaY
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func findPairs(galaxies []point) []galaxyPair {
	pairs := make(map[point]aoc.Set[point])
	for _, i := range galaxies {
		for _, j := range galaxies {
			if i != j {
				iPairs, iPresent := pairs[i]
				jPairs, jPresent := pairs[j]
				if iPresent {
					iPairs.Add(j)
				} else if jPresent {
					jPairs.Add(i)
				} else {
					pairs[i] = aoc.Set[point]{}
				}
			}
		}
	}

	var result []galaxyPair
	for i, points := range pairs {
		for j := range points {
			result = append(result, galaxyPair{i, j})
		}
	}
	return result
}

func findGalaxies(file string, emptyGrowthFactor int) []point {
	scanner := aoc.OpenScanner(file)
	var result []point
	var grid [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	y := 0
	emptyColumns := emptyColumnsOf(grid)
	for _, line := range grid {
		x := 0
		if isEmptyLine(line) {
			y += emptyGrowthFactor
			continue
		}

		for i, r := range line {
			emptyColumn := emptyColumns[i]
			if emptyColumn {
				x += emptyGrowthFactor
			}
			if r == '#' {
				result = append(result, point{x, y})
			}
			if !emptyColumn {
				x++
			}
		}
		y++

	}
	return result

}

func emptyColumnsOf(g [][]rune) []bool {
	result := make([]bool, len(g))
	for x, _ := range g {
		result[x] = isEmptyColumn(x, g)
	}
	return result
}

func isEmptyColumn(index int, grid [][]rune) bool {
	for _, line := range grid {
		if line[index] != '.' {
			return false
		}
	}
	return true
}

func isEmptyLine(line []rune) bool {
	for _, r := range line {
		if r != '.' {
			return false
		}
	}
	return true
}
