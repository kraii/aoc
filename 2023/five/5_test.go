package five

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveAExample1(t *testing.T) {
	assert.Equal(t, 35, solveA("/five/example-a.txt"))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 84470622, solveA("five/5.txt"))
}

func TestSolveBExample1(t *testing.T) {
	assert.Equal(t, 46, solveBWithBruteForce("/five/example-a.txt"))
}

/* Commented because my potato laptop can't cope with it (unlike my gaming pc)
func TestSolveB(t *testing.T) {
	assert.Equal(t, 26714516, solveBWithBruteForce("five/5.txt"))
}
*/

func TestSolveBBinary(t *testing.T) {
	assert.Equal(t, 26714516, solveBWithBinarySearch("five/5.txt"))
}
