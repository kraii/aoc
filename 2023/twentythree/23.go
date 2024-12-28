package twentythree

import (
	. "aoc/grids"
	"slices"
)

func solvePart1(file string) int {
	grid := ParseGrid(file)

	start := Point{X: findStartOrEnd(grid[0]), Y: 0}
	end := Point{X: findStartOrEnd(grid[len(grid)-1]), Y: len(grid) - 1}

	return search(grid, start, end, 0, Make2DSlice[bool](len(grid[0]), len(grid)))
}

func search(grid [][]rune, cur Point, end Point, distance int, visited [][]bool) int {
	if cur == end {
		return distance
	}
	visited[cur.Y][cur.X] = true
	maxDist := 0
	for _, point := range findAvailableMoves(grid, cur) {
		if !visited[point.Y][point.X] {
			maxDist = max(maxDist, search(grid, point, end, distance+1, visited))
		}
	}
	visited[cur.Y][cur.X] = false
	return maxDist
}

var directionMap = map[rune]Direction{
	'^': Up,
	'>': Right,
	'v': Down,
	'<': Left,
}

func findAvailableMoves(grid [][]rune, pos Point) []Point {
	currentTile := grid[pos.Y][pos.X]
	if currentTile == '.' {
		return availableMovesSimple(grid, pos)
	}

	direction, pres := directionMap[currentTile]
	if !pres {
		message := "wot is" + string(currentTile)
		panic(message)
	}

	newPos := Move1(pos, direction)
	if traversable(grid, newPos) {
		return []Point{newPos}
	}
	return nil
}

func availableMovesSimple(grid [][]rune, pos Point) []Point {
	var moves []Point
	for _, direction := range Directions {
		newPos := Move1(pos, direction)
		if traversable(grid, newPos) {
			moves = append(moves, newPos)
		}
	}
	return moves
}

func traversable(grid [][]rune, newPos Point) bool {
	return InRange(grid, newPos) && grid[newPos.Y][newPos.X] != '#'
}

func findStartOrEnd(row []rune) int {
	for x, r := range row {
		if r == '.' {
			return x
		}
	}
	panic("Couldn't find start or end")
}

func solvePart2(file string) int {
	grid := ParseGrid(file)

	start := Point{X: findStartOrEnd(grid[0]), Y: 0}
	end := Point{X: findStartOrEnd(grid[len(grid)-1]), Y: len(grid) - 1}

	g := buildGraph(grid)
	shortenGraph(g)

	return searchGraph(g, start, end, make(map[Point]int, len(grid)*len(grid[0])))
}

// Remove nodes that are in a "corridor" i.e. you can only go forward or back
// Add their weights to the nodes that are actual junctions
func shortenGraph(g graph) {
	for pos, neighbors := range g {
		// simple and can be removed
		if len(neighbors) == 2 {
			a, b := neighbors[0], neighbors[1]
			totalWeight := a.weight + b.weight
			removeEdge(g, pos, a.dest, b.dest, totalWeight)
			removeEdge(g, pos, b.dest, a.dest, totalWeight)
			delete(g, pos)
		}
	}
}

func removeEdge(g graph, removed Point, from Point, to Point, weight int) {
	edges := g[from]

	found := false
	for i, existing := range edges {
		if existing.dest == to {
			existing.weight = max(weight, existing.weight)
			found = true
		}
		if removed == existing.dest {
			edges = slices.Delete(edges, i, i+1)
			g[from] = edges
		}
	}

	if !found {
		g[from] = append(edges, &edge{to, weight})
	}
}

type edge struct {
	dest   Point
	weight int
}

type graph map[Point][]*edge

func buildGraph(grid [][]rune) graph {
	result := make(graph, len(grid)*len(grid[0]))

	for y, row := range grid {
		for x, r := range row {
			if r == '#' {
				continue
			}
			current := Point{X: x, Y: y}
			neighbors := availableMovesSimple(grid, current)
			edges := make([]*edge, len(neighbors))
			for i, neighbor := range neighbors {
				edges[i] = &edge{neighbor, 1}
			}
			result[current] = edges
		}
	}

	return result
}

func searchGraph(g graph, start Point, end Point, visited map[Point]int) int {
	if start == end {
		total := 0
		for _, d := range visited {
			total += d
		}
		return total
	}
	maxLen := 0
	for _, e := range g[start] {
		if visited[e.dest] == 0 {
			visited[e.dest] = e.weight
			maxLen = max(maxLen, searchGraph(g, e.dest, end, visited))
			visited[e.dest] = 0
		}
	}
	return maxLen
}
