package seventeen

import (
	. "aoc"
	. "aoc/grids"
	"strconv"
)

func solvePart1(file string) int {
	weights := parse(file)
	//return dijkstra(start, dest)
	return aStarSearch(weights, Point{X: 0, Y: 0}, Point{X: len(weights[0]) - 1, Y: len(weights) - 1}, possibleMoves)
}

func solvePart2(file string) int {
	weights := parse(file)
	//return dijkstra(start, dest)
	return aStarSearch(weights, Point{X: 0, Y: 0}, Point{X: len(weights[0]) - 1, Y: len(weights) - 1}, possibleMovesUltra)
}

type move struct {
	pos   Point
	dir   Direction
	count int // how many times we moved in this direction
}

func aStarSearch(weights [][]int, start Point, end Point, moveFinder func(prev move) []move) int {
	firstMove := move{start, Right, 0}
	pq := PriorityQueue[ScoredVertex]{ScoredVertex{firstMove, 0}}
	pq.Init()
	seen := map[move]int{firstMove: 0}

	for {
		if len(pq) == 0 {
			panic("Got lost")
		}
		current := pq.PopHeap()
		if current.v.pos == end {
			return seen[current.v]
		}
		moves := moveFinder(current.v)
		for _, m := range moves {
			_, visited := seen[m]
			if !visited && InRange(weights, m.pos) {
				total := seen[current.v] + score(weights, m.pos)

				n := ScoredVertex{m, total}
				pq.PushHeap(n)
				seen[m] = total
			}
		}
	}

	return -1
}

func score(weights [][]int, p Point) int {
	return weights[p.Y][p.X]
}

type ScoredVertex struct {
	v     move
	score int
}

func (sv ScoredVertex) Priority() int {
	return sv.score
}

func possibleMoves(prev move) []move {
	left := TurnLeft(prev.dir)
	right := TurnRight(prev.dir)
	moves := []move{
		{pos: Move1(prev.pos, left), dir: left, count: 1},
		{pos: Move1(prev.pos, right), dir: right, count: 1},
	}
	if prev.count < 3 {
		moves = append(moves, move{pos: Move1(prev.pos, prev.dir), dir: prev.dir, count: prev.count + 1})
	}

	return moves
}

func possibleMovesUltra(prev move) []move {
	left := TurnLeft(prev.dir)
	right := TurnRight(prev.dir)
	var moves []move
	if prev.count == 0 || prev.count >= 4 {
		moves = append(moves,
			move{pos: Move1(prev.pos, left), dir: left, count: 1},
			move{pos: Move1(prev.pos, right), dir: right, count: 1},
		)
	}
	if prev.count < 10 {
		moves = append(moves, move{pos: Move1(prev.pos, prev.dir), dir: prev.dir, count: prev.count + 1})
	}

	return moves
}

func parse(file string) [][]int {
	scanner := OpenScanner(file)
	var result [][]int
	for scanner.Scan() {
		line := scanner.Text()
		nodes := make([]int, len(line))

		for i, r := range line {
			v, err := strconv.Atoi(string(r))
			if err != nil {
				panic(err)
			}
			nodes[i] = v
		}

		result = append(result, nodes)
	}
	return result
}
