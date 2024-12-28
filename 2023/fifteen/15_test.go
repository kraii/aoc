package fifteen

import (
	"aoc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash(t *testing.T) {
	assert.Equal(t, 52, hash("HASH"))
}

func TestSumHash(t *testing.T) {
	assert.Equal(t, 1320, sumHashSequence("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"))
}

func TestSolvePart1(t *testing.T) {
	println(sumHashSequence(aoc.FileAsString("fifteen/input.txt")))
}

func TestInitializeLenses(t *testing.T) {
	assert.Equal(t, 145, initializeLenses("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"))
}

func TestSolvePart2(t *testing.T) {
	println(initializeLenses(aoc.FileAsString("fifteen/input.txt")))
}
