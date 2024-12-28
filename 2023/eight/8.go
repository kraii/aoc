package eight

import (
	"aoc"
	"aoc/maths"
	"fmt"
	"strings"
)

type mapNode struct {
	id    string
	left  string
	right string
}

func solvePart1(file string) int {
	turns, nodes := parse(file)

	current := "AAA"
	steps := 0
	i := 0

	for current != "ZZZ" {
		node := nodes[current]
		direction := turns[i]

		current = chooseNext(direction, node)

		steps++
		i++
		if i >= len(turns) {
			i = 0
		}
	}

	return steps
}

func solvePart2(file string) int {
	turns, nodes := parse(file)

	currentPositions := nodesEndingWithA(nodes)
	pathLengths := make([]int, len(currentPositions))
	steps := 0
	turnIndex := 0

	for stillTravelling(pathLengths) {
		direction := turns[turnIndex]

		for i := 0; i < len(currentPositions); i++ {
			node := nodes[currentPositions[i]]
			if pathLengths[i] > 0 {
				continue
			}

			next := chooseNext(direction, node)
			if strings.HasSuffix(next, "Z") {
				pathLengths[i] = steps + 1
			} else {
				currentPositions[i] = next
			}
		}

		steps++

		turnIndex++
		if turnIndex >= len(turns) {
			turnIndex = 0
		}
	}

	fmt.Printf("%+v\n", pathLengths)
	return maths.LcmAll(pathLengths)
}

func stillTravelling(paths []int) bool {
	for _, l := range paths {
		if l == 0 {
			return true
		}
	}
	return false
}

func chooseNext(direction rune, node mapNode) string {
	var next string
	if direction == 'L' {
		next = node.left
	} else {
		next = node.right
	}
	return next
}

func nodesEndingWithA(nodes map[string]mapNode) []string {
	var result []string
	for k := range nodes {
		if strings.HasSuffix(k, "A") {
			result = append(result, k)
		}
	}
	return result
}

func parse(file string) ([]rune, map[string]mapNode) {
	scanner := aoc.OpenScanner(file)
	scanner.Scan()
	turnsLine := scanner.Text()
	scanner.Scan() // skip empty line

	nodes := make(map[string]mapNode)

	for scanner.Scan() {
		nodeLine := scanner.Text()
		fields := strings.Fields(nodeLine)
		nodeId := fields[0]
		left := strings.Trim(fields[2], ",()")
		right := strings.Trim(fields[3], ",()")
		nodes[nodeId] = mapNode{nodeId, left, right}
	}

	return []rune(turnsLine), nodes
}
