package eight

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1Example1(t *testing.T) {
	assert.Equal(t, 2, solvePart1("eight/example1.txt"))
}

func TestSolvePart1Example2(t *testing.T) {
	assert.Equal(t, 6, solvePart1("eight/example2.txt"))
}

func TestSolvePart1(t *testing.T) {
	println(solvePart1("eight/8-input.txt"))
}

// part 2

func TestSolvePart2Example3(t *testing.T) {
	assert.Equal(t, 6, solvePart2("eight/example3.txt"))
}

func TestSolvePart2(t *testing.T) {
	println(solvePart2("eight/8-input.txt"))
}

func TestLCM(t *testing.T) {
	assert.Equal(t, 12, lcmAll([]int{2, 12}))
	assert.Equal(t, 12, lcmAll([]int{12, 2}))
}
