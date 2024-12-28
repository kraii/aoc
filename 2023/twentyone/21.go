package twentyone

import (
	. "aoc"
	. "aoc/grids"
)

func possibleDestinations(file string, maxSteps int) int {
	grid := ParseGrid(file)
	start := FindPoint(grid, 'S')

	points := reachablePoints(grid, start, maxSteps)
	//printGrid(grid, points)

	return countEven(points)
}

func countEven(points map[Point]int) int {
	total := 0
	for _, dist := range points {
		if dist%2 == 0 {
			total++
		}
	}
	return total
}

func printGrid(grid [][]rune, points map[Point]int) {
	for y, row := range grid {
		for x, r := range row {
			_, reached := points[Point{X: x, Y: y}]
			if reached {
				print("O")
			} else {
				print(string(r))
			}
		}
		println()
	}
}

func reachablePoints(grid [][]rune, start Point, maxSteps int) map[Point]int {
	toVisit := MakeQueue[Point](100)
	distances := map[Point]int{start: 0}

	PushItem(toVisit, start, 0)
	for len(*toVisit) > 0 {
		current := PopItem(toVisit)
		distToCurrent := distances[current]
		distToNext := distToCurrent + 1
		if distToNext > maxSteps {
			continue
		}
		for _, move := range possibleMoves(grid, current) {
			existingDist, p := distances[move]
			if !p || distToNext < existingDist {
				distances[move] = distToNext
				PushItem(toVisit, move, distToNext)
			}
		}

	}
	return distances
}

func possibleMoves(grid [][]rune, start Point) []Point {
	moves := make([]Point, 0, 4)
	for _, direction := range Directions {
		newLoc := Move1(start, direction)
		if InRange(grid, newLoc) && grid[newLoc.Y][newLoc.X] != '#' {
			moves = append(moves, newLoc)
		}
	}
	return moves
}

const maxStepsPart2 = 26501365

func toInfinity(file string, maxSteps int) int {
	grid := ParseGrid(file)
	FindPoint(grid, 'S')

	panic("I cannot get this to work")
}
