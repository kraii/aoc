package fourteen

import (
	"aoc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTilt1Column(t *testing.T) {
	p := platform{{'O'}, {'O'}, {'.'}, {'O'}, {'.'}, {'O'}, {'.'}, {'.'}, {'#'}, {'#'}}
	p.print()
	tiltNegative(p, north)

	println("------")
	p.print()
	expected := platform{{'O'}, {'O'}, {'O'}, {'O'}, {'.'}, {'.'}, {'.'}, {'.'}, {'#'}, {'#'}}
	assert.Equal(t, expected, p)
}

func TestTilt1ColumnSouth(t *testing.T) {
	p := platform{{'O'}, {'O'}, {'.'}, {'O'}, {'.'}, {'O'}, {'.'}, {'.'}, {'#'}, {'#'}}
	p.print()
	tiltPositive(p, south)

	println("------")
	p.print()
	expected := platform{{'.'}, {'.'}, {'.'}, {'.'}, {'O'}, {'O'}, {'O'}, {'O'}, {'#'}, {'#'}}
	assert.Equal(t, expected, p)
}

func TestTilt(t *testing.T) {
	p := parse("fourteen/14-ex-1.txt")
	p.print()
	tiltNegative(p, north)
	println()
	p.print()

	expected := parse("fourteen/14-ex-1-expected.txt")
	assert.Equal(t, expected, p)
}

func TestSolvePart1Example(t *testing.T) {
	assert.Equal(t, 136, solvePart1("fourteen/14-ex-1.txt"))
}

func TestSolvePart1(t *testing.T) {
	println(solvePart1("fourteen/14-input.txt"))
}

func TestSpin1(t *testing.T) {
	p := parse("fourteen/14-ex-1.txt")
	spin(p, 1)

	expected := parse("fourteen/14-ex-1-expected-1-cycle.txt")
	assert.Equal(t, expected, p)
}

func TestSpin2(t *testing.T) {
	p := parse("fourteen/14-ex-1.txt")
	spin(p, 2)

	expected := parse("fourteen/14-ex-1-expected-2-cycle.txt")
	assert.Equal(t, expected, p)
}

func TestSpin3(t *testing.T) {
	p := parse("fourteen/14-ex-1.txt")
	spin(p, 3)

	expected := parse("fourteen/14-ex-1-expected-3-cycle.txt")
	assert.Equal(t, expected, p)
}

func TestCycleMan(t *testing.T) {
	p := parse("fourteen/14-ex-1.txt")
	spin(p, 1e10)

}

func TestContains(t *testing.T) {
	a := make(aoc.Set[string])
	assert.Equal(t, false, a.Contains("b"))
}

func TestSolvePart2Example(t *testing.T) {
	assert.Equal(t, 64, solvePart2("fourteen/14-ex-1.txt", 1e9))
}

func TestSolvePart2(t *testing.T) {
	println(solvePart2("fourteen/14-input.txt", 1e9))
}
