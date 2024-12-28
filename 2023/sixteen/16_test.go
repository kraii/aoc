package sixteen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 46, solvePart1("sixteen/16-ex-1.txt"))
}

func TestSolvePart1(t *testing.T) {
	println(solvePart1("sixteen/input.txt"))
}

func TestSolvePart2(t *testing.T) {
	println(findBestConfig("sixteen/input.txt"))
}
