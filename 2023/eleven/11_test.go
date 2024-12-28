package eleven

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1Example(t *testing.T) {
	assert.Equal(t, 374, solve("eleven/11e1.txt", 2))
}

func TestSolvePart1(t *testing.T) {
	println(solve("eleven/11-input.txt", 2))
}

func TestSolvePart2Example(t *testing.T) {
	assert.Equal(t, 1030, solve("eleven/11e1.txt", 10))
}

func TestSolvePart2(t *testing.T) {
	println(solve("eleven/11-input.txt", 1e6))
}
func TestDoubleCheckEquality(t *testing.T) {
	assert.Equal(t, true, point{1, 1} == point{1, 1})
	a := [2]point{{1, 1}, {1, 1}}
	b := [2]point{{1, 1}, {1, 1}}
	assert.Equal(t, true, a == b)
}

func TestDistanceBetween(t *testing.T) {
	a, b := point{4, 0}, point{10, 9}
	assert.Equal(t, 15, distanceBetween(a, b))

	a, b = point{0, 0}, point{1, 1}
	assert.Equal(t, 2, distanceBetween(a, b))
}
