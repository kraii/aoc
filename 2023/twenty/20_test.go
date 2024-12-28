package twenty

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcPulsesExample1(t *testing.T) {
	assert.Equal(t, 8*4, calcPulses("twenty/ex-input.txt", 1))
	assert.Equal(t, 32000000, calcPulses("twenty/ex-input.txt", 1000))
}

func TestCalcPulsesExample2(t *testing.T) {
	assert.Equal(t, 11687500, calcPulses("twenty/ex-input2.txt", 1000))
}

func TestSolvePart1(t *testing.T) {
	println(calcPulses("twenty/input.txt", 1000))
}

func TestSolvePart2(t *testing.T) {
	println(findRxLowSend("twenty/input.txt"))
}

func TestGraphviz(t *testing.T) {
	createGraphImage()
}
