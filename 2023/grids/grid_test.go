package grids

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointComparable(t *testing.T) {
	assert.Equal(t, Point{0, 0}, Point{0, 0})
	assert.NotEqual(t, Point{1, 0}, Point{0, 0})
}
