package seventeen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1Example(t *testing.T) {
	assert.Equal(t, 102, solvePart1("seventeen/17-ex-1.txt"))
}

func TestSolvePart1(t *testing.T) {
	println(solvePart1("seventeen/input.txt"))
}

func TestSolvePart2(t *testing.T) {
	println(solvePart2("seventeen/input.txt"))
}
