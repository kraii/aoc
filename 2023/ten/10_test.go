package ten

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckPoint(t *testing.T) {
	assert.Equal(t, true, point{1, 2} == (point{1, 2}))
	assert.Equal(t, false, point{3, 4} == (point{3, 7}))
	assert.Equal(t, false, point{5, 4} == (point{3, 4}))
}

func TestFindStart(t *testing.T) {
	p := findStart(parse("ten/10-example-input.txt"))
	assert.Equal(t, p.x, 1)
	assert.Equal(t, p.y, 1)
}

func TestFindValidMoveFromStart(t *testing.T) {
	p := findValidMoveFromStart(parse("ten/10-example-input.txt"), point{1, 1})
	assert.Equal(t, p.x, 2)
	assert.Equal(t, p.y, 1)
}

func TestSolvePart1Example(t *testing.T) {
	assert.Equal(t, 4, solvePart1("ten/10-example-input.txt"))
}

func TestSolvePart1Example2(t *testing.T) {
	assert.Equal(t, 8, solvePart1("ten/10-example-input-2.txt"))
}

func TestSolvePart1(t *testing.T) {
	println(solvePart1("ten/10-input.txt"))
}

func TestSolvePart2Example3(t *testing.T) {
	assert.Equal(t, 4, solvePart2("ten/10-example-input-3.txt"))
}

func TestSolvePart2(t *testing.T) {
	println(solvePart2("ten/10-input.txt"))
}
