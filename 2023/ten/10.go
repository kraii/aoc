package ten

import (
	grid2 "aoc/grids"
	"fmt"
)

type grid [][]rune

type point struct {
	x, y int
}

func solvePart1(file string) int {
	g := grid2.ParseGrid(file)
	route := traverse(g)
	fmt.Printf("Route %+v\n", route)
	return len(route) / 2
}

func solvePart2(file string) int {
	g := grid2.ParseGrid(file)
	route := traverse(g)
	area := calculateShoelace(route)
	// https://en.wikipedia.org/wiki/Pick%27s_theorem
	interiorPoints := area - len(route)/2 + 1
	fmt.Printf("Area %d, vertices %d\n", area, len(route))
	return interiorPoints
}

// https://en.wikipedia.org/wiki/Shoelace_formula
func calculateShoelace(route []point) int {
	r := 0
	for i := range route[:(len(route) - 2)] {
		r += (route[i].x - route[i+1].x) * (route[i].y + route[i+1].y)
	}
	// Wrap around for end of route
	r += (route[len(route)-1].x - route[0].x) * (route[0].y + route[len(route)-1].y)
	return abs(r) / 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func traverse(g grid) []point {
	start := findStart(g)
	visited := []point{start}

	prev := start
	current := findValidMoveFromStart(g, start)
	for current != start {
		next := chooseNext(g, prev, current)
		prev, current = current, next
		visited = append(visited, current)
	}
	return visited
}

func findValidMoveFromStart(g grid, start point) point {
	for i := 1; i < 5; i++ {
		location := move(start, direction(i))
		if !g.inBounds(location) {
			continue
		}
		moves, validMove := pipeTypes[g.value(location)]
		if !validMove {
			continue
		}
		for _, m := range moves {
			if move(location, m) == start {
				return location
			}
		}
	}
	panic("I can't find anywhere to start :(")
}

//func (p point) eq(other point) bool {
//	return p.x == other.x && p.y == other.y
//}

func (g grid) value(p point) rune {
	return g[p.y][p.x]
}

func (g grid) inBounds(p point) bool {
	maxX, maxY := len(g[0]), len(g)
	return p.x > 0 && p.x < maxX && p.y >= 0 && p.y < maxY
}

func move(p point, d direction) point {
	switch d {
	case up:
		return point{p.x, p.y - 1}
	case down:
		return point{p.x, p.y + 1}
	case left:
		return point{p.x - 1, p.y}
	case right:
		return point{p.x + 1, p.y}
	}
	panic(fmt.Sprintf("What the heck direction is %d", d))
}

type direction int

const (
	up    direction = 1
	right direction = 2
	down  direction = 3
	left  direction = 4
)

var pipeTypes = map[rune][]direction{
	'|': {up, down},
	'-': {left, right},
	'L': {up, right},
	'J': {up, left},
	'7': {down, left},
	'F': {down, right},
}

func chooseNext(g grid, prev point, current point) point {
	value := g.value(current)
	possibleMoves := pipeTypes[value]
	for _, dir := range possibleMoves {
		move := move(current, dir)
		if move != prev {
			return move
		}
	}
	fmt.Printf("No moves found at current %+v, prev %+v\n", prev, current)
	panic("I got lost :(")
}

func findStart(g grid) point {
	for y, row := range g {
		for x, r := range row {
			if r == 'S' {
				return point{x, y}
			}
		}
	}
	panic("Where's me S at")
}
