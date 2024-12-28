package nine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNext(t *testing.T) {
	assert.Equal(t, 18, next([]int{0, 3, 6, 9, 12, 15}))
	assert.Equal(t, 28, next([]int{1, 3, 6, 10, 15, 21}))
	assert.Equal(t, 68, next([]int{10, 13, 16, 21, 30, 45}))

	assert.Equal(t, 9622, next([]int{11, 12, 8, -2, -17, -34, -48, -52, -37, 8, 96, 242, 463, 778, 1208, 1776, 2507, 3428, 4568, 5958, 7631}))
}

func TestPrev(t *testing.T) {
	assert.Equal(t, -3, prev([]int{0, 3, 6, 9, 12, 15}))
	assert.Equal(t, 0, prev([]int{1, 3, 6, 10, 15, 21}))
	assert.Equal(t, 5, prev([]int{10, 13, 16, 21, 30, 45}))
}

func TestSolvePart1Example(t *testing.T) {
	assert.Equal(t, 114, solve("nine/9-example.txt", next))
}

func TestSolvePart2Example(t *testing.T) {
	assert.Equal(t, 2, solve("nine/9-example.txt", prev))
}

func TestSolvePart1(t *testing.T) {
	println(solve("nine/9-input.txt", next))
}

func TestSolvePart2(t *testing.T) {
	println(solve("nine/9-input.txt", prev))
}
