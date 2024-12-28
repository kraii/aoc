package twentyone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExample1(t *testing.T) {
	assert.Equal(t, 16, possibleDestinations("twentyone/ex-1.txt", 6))
}

func TestSolvePart1(t *testing.T) {
	println(possibleDestinations("twentyone/input.txt", 64))
}

func TestSolvePart2Example(t *testing.T) {
	assert.Equal(t, 16733044, toInfinity("twentyone/ex-1.txt", 5000))
}

func TestSolvePart2(t *testing.T) {
	println(toInfinity("twentyone/input.txt", maxStepsPart2))
}
