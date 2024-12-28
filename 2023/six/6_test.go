package six

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveA(t *testing.T) {
	assert.Equal(t, 1084752, solveA())
}

func TestSolveBExample(t *testing.T) {
	assert.Equal(t, 71503, solveBruteForce(71530, 940200))
}

/*
Input was
---
Time:        40     70     98     79
Distance:   215   1051   2147   1005
---
*/
func TestSolveB(t *testing.T) {
	//it's only 10ms on my machine, but I feel like we were supposed to use maths lol
	assert.Equal(t, 28228952, solveBruteForce(40709879, 215105121471005))

	assert.Equal(t, 28228952, solveB(40709879, 215105121471005))
}
