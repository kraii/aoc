package three

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveAExample(t *testing.T) {
	assert.Equal(t, 4361, solveA("three/smol.txt"))
}

func TestSolveAExample2(t *testing.T) {
	assert.Equal(t, 596+954+9, solveA("three/smol2.txt"))
}

func TestSolveA(t *testing.T) {
	println(solveA("three/input.txt"))
}

func TestSymbool(t *testing.T) {
	assert.Equal(t, isSymbol('.'), false, ".")
	assert.Equal(t, isSymbol('#'), true, "#")
}

func TestSolveBExample(t *testing.T) {
	assert.Equal(t, 467835, solveB("three/gear-check.txt"))
}

func TestSolveB(t *testing.T) {
	println(solveB("three/input.txt"))
}

func TestGetNumber(t *testing.T) {
	number, i := getNumber("..123...", 3)

	assert.Equal(t, 123, number)
	assert.Equal(t, 5, i)

	number, i = getNumber("456....", 0)

	assert.Equal(t, 456, number)
	assert.Equal(t, 3, i)

	number, i = getNumber("456....", 2)

	assert.Equal(t, 456, number)
	assert.Equal(t, 3, i)

}
