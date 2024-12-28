package twelve

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountPossible(t *testing.T) {
	assert.Equal(t, 10, countPossible("?###????????", []int{3, 2, 1}))
	assert.Equal(t, 1, countPossible("#.#.###", []int{1, 1, 3}))
	assert.Equal(t, 4, countPossible(".??..??...?##.", []int{1, 1, 3}))
	fmt.Printf("mappy %+v\n", cache)
}

func TestCountPossible2(t *testing.T) {
	assert.Equal(t, 4, countPossible(".??..??...?##.", []int{1, 1, 3}))
}

func TestNoWorkingSprings(t *testing.T) {
	assert.Equal(t, true, noWorkingSprings("###.", 3))
	assert.Equal(t, false, noWorkingSprings("##..", 3))
}

func TestNextSpringIsNotWorking(t *testing.T) {
	assert.Equal(t, true, chainIsNotTooLong("###.", 3))
	assert.Equal(t, false, chainIsNotTooLong("####", 3))
	assert.Equal(t, true, chainIsNotTooLong("....", 4))
}

func TestCheckMySlicing(t *testing.T) {
	assert.Equal(t, "###", "###????????"[:3])
	assert.Equal(t, "????????", "###????????"[3:])
}

func TestPart1(t *testing.T) {
	println(solve("twelve/12-input.txt", 1))
}

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 21, solve("twelve/12-example.txt", 1))
}

func TestPart2Example(t *testing.T) {
	assert.Equal(t, 525152, solve("twelve/12-example.txt", 5))
}

func TestPart2(t *testing.T) {
	println(solve("twelve/12-input.txt", 5))
}
