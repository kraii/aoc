package seven

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveA(t *testing.T) {
	println(solveA("seven/7-input.txt"))
}

func TestSolveB(t *testing.T) {
	println(solveB("seven/7-input.txt"))
}

func Test7aExample(t *testing.T) {
	assert.Equal(t, 6440, solveA("seven/7-example-a.txt"))
}

func Test7bExample(t *testing.T) {
	assert.Equal(t, 5905, solveB("seven/7-example-a.txt"))
}
func Test7aExample2(t *testing.T) {
	assert.Equal(t, 6592, solveA("seven/7-example-a-2.txt"))
}

func Test7bExample2(t *testing.T) {
	assert.Equal(t, 6839, solveB("seven/7-example-a-2.txt"))
}
func TestRankHandTypePart1(t *testing.T) {
	assert.Equal(t, fiveOfAKind, rankHandTypePart1("AAAAA"), "5 of a kind AAAAA")
	assert.Equal(t, fourOfAKind, rankHandTypePart1("AA8AA"), "4 of a kind AA8AA")
	assert.Equal(t, fullHouse, rankHandTypePart1("23332"), "full house 23332")
	assert.Equal(t, threeOfAKind, rankHandTypePart1("TTT98"), "3 of a kind TTT98")
	assert.Equal(t, twoPair, rankHandTypePart1("23432"), "2 pair 23432")
	assert.Equal(t, onePair, rankHandTypePart1("A23A4"), "pair A23A4")
	assert.Equal(t, highCard, rankHandTypePart1("23456"), "high card 23456")
}

func TestRankHandTypePart2(t *testing.T) {
	assert.Equal(t, fiveOfAKind, rankHandTypePart2("AAAAA"), "5 of a kind AAAAA")
	assert.Equal(t, fourOfAKind, rankHandTypePart2("AA8AA"), "4 of a kind AA8AA")
	assert.Equal(t, fullHouse, rankHandTypePart2("23332"), "full house 23332")
	assert.Equal(t, threeOfAKind, rankHandTypePart2("TTT98"), "3 of a kind TTT98")
	assert.Equal(t, twoPair, rankHandTypePart2("23432"), "2 pair 23432")
	assert.Equal(t, onePair, rankHandTypePart2("A23A4"), "pair A23A4")
	assert.Equal(t, highCard, rankHandTypePart2("23456"), "high card 23456")

	// Joker examples
	assert.Equal(t, fiveOfAKind, rankHandTypePart2("AAAAJ"), "5 of a kind AAAAJ")
	assert.Equal(t, fourOfAKind, rankHandTypePart2("AAAJ2"), "4 of a kind AAAJ2")
	assert.Equal(t, fullHouse, rankHandTypePart2("AA22J"), "Full house AA22J")
	assert.Equal(t, threeOfAKind, rankHandTypePart2("AK22J"), "3 of a kind AK22J")
	assert.Equal(t, threeOfAKind, rankHandTypePart2("Q2KJJ"), "3 of a kind Q2KJJ")
	assert.Equal(t, onePair, rankHandTypePart2("AK32J"), "1 pair AK32J")
}

func TestRankCardPart1(t *testing.T) {
	assert.Equal(t, 14, rankCardPart1('A'))
	assert.Equal(t, 2, rankCardPart1('2'))
	assert.Equal(t, 9, rankCardPart1('9'))
	assert.Equal(t, 10, rankCardPart1('T'))
}
