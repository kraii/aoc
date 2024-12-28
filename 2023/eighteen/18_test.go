package eighteen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 62, solve("eighteen/18-ex.txt", false))
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 952408144115, solve("eighteen/18-ex.txt", true))
}

func TestPart1(t *testing.T) {
	println(solve("eighteen/input.txt", false))
}

func TestPart2(t *testing.T) {
	println(solve("eighteen/input.txt", true))
}
