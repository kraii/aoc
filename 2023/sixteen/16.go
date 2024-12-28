package sixteen

import (
	"aoc"
	"fmt"
)

type grid [][]rune
type direction int

const (
	up    direction = 1
	right direction = 2
	down  direction = 3
	left  direction = 4
)

type point struct {
	x, y int
}

type visit struct {
	pos point
	dir direction
}

func solvePart1(file string) int {
	g := parse(file)
	return countEnergized(g, point{0, 0}, right)
}

func countEnergized(g grid, startPos point, startDir direction) int {
	visits := make(aoc.Set[visit])
	traverse(g, visits, startPos, startDir)
	visitedPoints := make(aoc.Set[point])
	for v, _ := range visits {
		visitedPoints.Add(v.pos)
	}
	return len(visitedPoints)
}

func findBestConfig(file string) int {
	g := parse(file)
	maxX, maxY := len(g[0])-1, len(g)-1

	best := 0
	for y := range g {
		fromLeft := countEnergized(g, point{0, y}, right)
		if fromLeft > best {
			best = fromLeft
		}
		fromRight := countEnergized(g, point{maxX, y}, left)
		if fromRight > best {
			best = fromRight
		}
	}

	for x := range g[0] {
		fromTop := countEnergized(g, point{x, 0}, down)
		if fromTop > best {
			best = fromTop
		}
		fromBottom := countEnergized(g, point{x, maxY}, up)
		if fromBottom > best {
			best = fromBottom
		}
	}

	return best
}

func printState(g grid, visited aoc.Set[point]) {
	for y, row := range g {
		for x, r := range row {
			if visited.Contains(point{x, y}) {
				print("#")
			} else {
				print(string(r))
			}
		}
		println()
	}
}

func parse(file string) grid {
	scanner := aoc.OpenScanner(file)
	var g grid
	for scanner.Scan() {
		line := scanner.Text()
		g = append(g, []rune(line))
	}

	return g
}

func traverse(g grid, visits aoc.Set[visit], startPos point, startDir direction) {
	pos := startPos
	dir := startDir

	for pos.x >= 0 && pos.y >= 0 && pos.x < len(g[0]) && pos.y < len(g) {
		tile := g[pos.y][pos.x]
		if visits.Contains(visit{pos, dir}) {
			return
		}
		visits.Add(visit{pos, dir})

		switch tile {
		case '.':
		case '\\':
			dir = turnBackslash(dir)
		case '/':
			dir = turnForwardSlash(dir)
		case '|':
			if dir == left || dir == right {
				traverse(g, visits, move(up, pos), up)
				dir = down
			}
		case '-':
			if dir == up || dir == down {
				traverse(g, visits, move(left, pos), left)
				dir = right
			}

		default:
			panic("we're lost AF")
		}
		pos = move(dir, pos)
	}
}

// /
func turnForwardSlash(dir direction) direction {
	switch dir {
	case up:
		return right
	case down:
		return left
	case left:
		return down
	case right:
		return up
	default:
		panic(wrongDir(dir))
	}
}

// \
func turnBackslash(dir direction) direction {
	switch dir {
	case up:
		return left
	case down:
		return right
	case left:
		return up
	case right:
		return down
	default:
		panic(wrongDir(dir))
	}
}

func move(dir direction, pos point) point {
	switch dir {
	case up:
		return point{pos.x, pos.y - 1}
	case down:
		return point{pos.x, pos.y + 1}
	case left:
		return point{pos.x - 1, pos.y}
	case right:
		return point{pos.x + 1, pos.y}

	default:
		panic(wrongDir(dir))
	}
}

func wrongDir(dir direction) string {
	return fmt.Sprintf("where the hell are we going %d", dir)
}
