package twentythree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const example = "twentythree/ex.txt"
const input = "twentythree/input.txt"

func TestExample1(t *testing.T) {
	assert.Equal(t, 94, solvePart1(example))
}

func TestFindStart(t *testing.T) {
	assert.Equal(t, 1, findStartOrEnd([]rune("#.#####################")))
}

func TestSolvePart1(t *testing.T) {
	println(solvePart1(input))
}

func TestExample1Part2(t *testing.T) {
	assert.Equal(t, 154, solvePart2(example))
}

func TestSolvePart2(t *testing.T) {
	println(solvePart2(input))
}
