package thirteen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	valleys := readFile("thirteen/13-example.txt")
	assert.Equal(t, 2, len(valleys))
	assert.Equal(t, "#.##..#", valleys[0].vertical[0])
	assert.Equal(t, "#...##..#", valleys[1].horizontal[0])
	for i, v := range valleys {
		for j, h := range v.horizontal {
			assert.NotEmpty(t, h, "Valley %d row %d", i, j)
		}
	}
}

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 405, solve("thirteen/13-example.txt", 0))
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 400, solve("thirteen/13-example.txt", 1))
}

func TestPart1Example2(t *testing.T) {
	assert.Equal(t, 709, solve("thirteen/13-example-2.txt", 0))
}

func TestPart1(t *testing.T) {
	println(solve("thirteen/13-input.txt", 0))
}

func TestPart2(t *testing.T) {
	println(solve("thirteen/13-input.txt", 1))
}

func TestCountBeforeReflection(t *testing.T) {
	valleys := readFile("thirteen/13-example.txt")
	assert.Equal(t, 0, countBeforeReflection(valleys[0].horizontal, 0))
	assert.Equal(t, 5, countBeforeReflection(valleys[0].vertical, 0))
	assert.Equal(t, 0, countBeforeReflection(valleys[1].vertical, 0))
	assert.Equal(t, 4, countBeforeReflection(valleys[1].horizontal, 0))
}
