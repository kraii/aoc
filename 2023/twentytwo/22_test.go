package twentytwo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const inputFile = "twentytwo/input.txt"
const example1 = "twentytwo/22-e.txt"
const example2 = "twentytwo/22-e2.txt"

func TestParse(t *testing.T) {
	bricks := parse(example1)

	assert.Equal(t, vector{1, 2, 1}, bricks[0].end)
	assert.Equal(t, vector{0, 0, 4}, bricks[3].start)
}

func TestLowestZ(t *testing.T) {
	assert.Equal(t, 8, lowest(&brick{start: vector{1, 1, 8}, end: vector{1, 1, 9}}))
}

func TestIntersect(t *testing.T) {
	a := brick{start: vector{1, 2, 1}, end: vector{1, 2, 1}}
	b := brick{start: vector{1, 2, 2}, end: vector{1, 2, 2}}

	assert.Equal(t, true, intersects(&a, &b))
}

func TestIntersect2(t *testing.T) {
	a := brick{start: vector{0, 0, 1}, end: vector{0, 1, 1}}
	b := brick{start: vector{1, 1, 1}, end: vector{1, 1, 1}}
	c := brick{start: vector{0, 0, 2}, end: vector{0, 0, 2}}

	assert.Equal(t, true, intersects(&a, &c))
	assert.Equal(t, false, intersects(&a, &b))
}

func TestFall(t *testing.T) {
	bricks := parse(example1)
	fall(bricks)

	assert.Equal(t, 0, bricks[0].start.z)
	assert.Equal(t, 0, bricks[0].end.z)

	assert.Equal(t, 1, bricks[1].start.z)
	assert.Equal(t, 1, bricks[1].end.z)
}

func TestFall2(t *testing.T) {
	bricks := parse(example2)
	fall(bricks)

	for _, b := range bricks {
		fmt.Printf("%s=%d,", b.label, lowest(b))
	}
	println()

	assert.Equal(t, 0, bricks[0].start.z)
	assert.Equal(t, 0, bricks[0].end.z)

	assert.Equal(t, 1, bricks[2].start.z)
	assert.Equal(t, 1, bricks[2].end.z)
}

func TestFallVerySimple(t *testing.T) {
	bricks := []*brick{
		{start: vector{1, 2, 1}, end: vector{1, 2, 1}},
		{start: vector{1, 2, 2}, end: vector{1, 2, 2}},
	}
	fall(bricks)

	assert.Equal(t, 0, bricks[0].start.z)
	assert.Equal(t, 0, bricks[0].end.z)

	assert.Equal(t, 1, bricks[1].start.z)
	assert.Equal(t, 1, bricks[1].end.z)
}

func TestSolvePart1Example(t *testing.T) {
	assert.Equal(t, 5, solvePart1(example1))
}

func TestSolvePart1Example2(t *testing.T) {
	assert.Equal(t, 3, solvePart1(example2))
}

func TestSolvePart1(t *testing.T) {
	println(solvePart1(inputFile))
}

func TestSolvePart2Example(t *testing.T) {
	assert.Equal(t, 7, solvePart2(example1))
}

func TestSolvePart2(t *testing.T) {
	println(solvePart2(inputFile)) // 109531
}
