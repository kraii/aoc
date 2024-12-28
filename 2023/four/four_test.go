package four

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveAExample1(t *testing.T) {
	assert.Equal(t, 13, solveA("/four/a-example.txt"))
}

func TestSolveA(t *testing.T) {
	println(solveA("/four/input.txt"))
}

func TestSolveBExample1(t *testing.T) {
	assert.Equal(t, 30, solveB("/four/b-example.txt"))
}

func TestSolveB(t *testing.T) {
	println(solveB("/four/input.txt"))
}
