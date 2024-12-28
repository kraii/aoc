package twentytwo

import (
	. "aoc"
	"cmp"
	"slices"
	"strings"
)

type vector struct {
	x, y, z int
}

type brick struct {
	label      string
	start, end vector
}

func parse(file string) []*brick {
	scanner := OpenScanner(file)
	var bricks []*brick
	i := 0
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, "~")
		bricks = append(bricks, &brick{toBrickId(i), parseVec(split[0]), parseVec(split[1])})
		i++
	}
	return bricks
}

func parseVec(s string) vector {
	split := strings.Split(s, ",")
	x, y, z := split[0], split[1], split[2]
	return vector{ToInt(x), ToInt(y), ToInt(z)}
}

func lowest(b *brick) int {
	return min(b.start.z, b.end.z)
}

func highest(b *brick) int {
	return max(b.start.z, b.end.z)
}

func fall(bricks []*brick) {
	slices.SortFunc(bricks, func(a, b *brick) int { return cmp.Compare(lowest(a), lowest(b)) })

	settled := false

	for !settled {
		settled = true // maybe

		for _, b := range bricks {
			startZ := lowest(b)
			bBelow := findBelow(b, bricks)
			if bBelow == nil {
				if startZ != 0 {
					settled = false
					lowerBy(b, startZ)
				}
				continue
			}

			settledHeight := highest(bBelow) + 1
			if startZ != settledHeight {
				settled = false
				diff := startZ - settledHeight
				lowerBy(b, diff)
			}
		}
	}
}

func lowerBy(b *brick, diff int) {
	b.start.z -= diff
	b.end.z -= diff
}

func findBelow(start *brick, bricks []*brick) *brick {
	var found *brick

	for _, b := range bricks {
		if b == start {
			continue
		}
		if below(start, b) && intersects(start, b) {
			if found == nil || highest(b) > highest(found) {
				found = b
			}
		}

	}

	return found
}

func below(start *brick, other *brick) bool {
	return highest(other) < lowest(start)
}

func intersects(b1 *brick, b2 *brick) bool {
	return max(b1.start.x, b2.start.x) <= min(b1.end.x, b2.end.x) && max(b1.start.y, b2.start.y) <= min(b1.end.y, b2.end.y)
}

func solvePart1(file string) int {
	bricks := parse(file)
	fall(bricks)

	result := 0
	supportedBy := mapSupportedBy(bricks)
	for _, b := range bricks {
		countOnlySupport := 0
		for _, other := range bricks {
			if b == other {
				continue
			}
			if supports(b, other) {
				if len(supportedBy[other.label]) < 2 {
					countOnlySupport++
				}
			}
		}
		if countOnlySupport == 0 {
			result++
		}
	}
	return result
}

func solvePart2(file string) int {
	bricks := parse(file)
	fall(bricks)

	result := 0
	supportMap := mapSupportedBy(bricks)

	for _, toRemove := range bricks {
		willChain := make(Set[string])
		for _, other := range bricks {
			_, seen := willChain[other.label]
			if toRemove == other || seen || lowest(other) == 0 {
				continue
			}

			supporting := supportMap[other.label]

			if willFall(toRemove, supporting, willChain) {
				willChain.Add(other.label)
			}
		}
		result += len(willChain)
	}
	return result
}

func willFall(remove *brick, supporting []string, allReadyFalling Set[string]) bool {
	supportsRemoved := 0
	for _, s := range supporting {
		if s == remove.label || allReadyFalling.Contains(s) {
			supportsRemoved++
		}
	}

	return len(supporting)-supportsRemoved == 0
}

func mapSupportedBy(bricks []*brick) map[string][]string {
	result := make(map[string][]string)

	for _, b := range bricks {
		for _, other := range bricks {
			if b == other {
				continue
			}

			if supportedBy(b, other) {
				supports := result[b.label]
				result[b.label] = append(supports, other.label)
			}
		}
	}

	return result
}

func toBrickId(i int) string {
	return string(rune(i + 97))
}

func supports(support *brick, supported *brick) bool {
	return lowest(supported) == highest(support)+1 && intersects(support, supported)
}

func supportedBy(b *brick, other *brick) bool {
	return lowest(b) == highest(other)+1 && intersects(b, other)
}
