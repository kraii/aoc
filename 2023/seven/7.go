package seven

import (
	"aoc"
	"cmp"
	"slices"
	"strconv"
	"strings"
)

type camelPlay struct {
	hand string
	bid  int
}

const (
	fiveOfAKind  = 7
	fourOfAKind  = 6
	fullHouse    = 5
	threeOfAKind = 4
	twoPair      = 3
	onePair      = 2
	highCard     = 1
)

func solveA(file string) int {
	return solve(file, buildCompare(rankHandTypePart1, rankCardPart1))
}

func solve(file string, cmp func(a camelPlay, b camelPlay) int) int {
	plays := parse(file)
	slices.SortFunc(plays, cmp)
	rank := 0
	total := 0
	for _, play := range plays {
		rank++
		println(play.hand, play.bid)
		total += play.bid * rank

	}
	return total
}

func solveB(file string) int {
	return solve(file, buildCompare(rankHandTypePart2, rankCardPart2))
}

func buildCompare(handTypeRank func(hand string) int, cardRank func(card uint8) int) func(a camelPlay, b camelPlay) int {

	return func(a camelPlay, b camelPlay) int {
		r := cmp.Compare(handTypeRank(a.hand), handTypeRank(b.hand))
		if r == 0 {
			for i := 0; i < 5; i++ {
				ai, bi := a.hand[i], b.hand[i]
				if ai != bi {
					return cmp.Compare(cardRank(ai), cardRank(bi))
				}
			}
		}
		return r
	}
}

func rankHandTypePart1(hand string) int {
	set := newCardSet(hand)
	distinctCards := len(set)

	switch distinctCards {
	case 1:
		return fiveOfAKind
	case 2:
		if maxRepeatCount(set, false) == 4 {
			return fourOfAKind
		} else {
			return fullHouse
		}
	case 3:
		if maxRepeatCount(set, false) == 3 {
			return threeOfAKind
		} else {
			return twoPair
		}
	case 4:
		return onePair
	default:
		return highCard
	}
}

func rankHandTypePart2(hand string) int {
	set := newCardSet(hand)
	distinctCards := len(set)
	numJokers := set['J']
	maxRepeat := maxRepeatCount(set, false)
	maxRepeatNotJoker := maxRepeatCount(set, true)

	switch distinctCards {
	case 1:
		return fiveOfAKind
	case 2:
		if numJokers+maxRepeatNotJoker == 5 {
			return fiveOfAKind
		} else if maxRepeat == 4 {
			return fourOfAKind
		} else {
			return fullHouse
		}
	case 3:
		if numJokers+maxRepeatNotJoker == 4 {
			return fourOfAKind
		} else if numJokers == 1 && maxRepeat == 2 {
			return fullHouse
		}
		if maxRepeat == 3 || (numJokers+maxRepeatNotJoker == 3) {
			return threeOfAKind
		} else {
			return twoPair
		}
	case 4:
		if numJokers+maxRepeatNotJoker == 3 {
			return threeOfAKind
		}
		return onePair
	default:
		if numJokers+maxRepeatNotJoker == 2 {
			return onePair
		}
		return highCard
	}
}

func maxRepeatCount(set cardSet, ignoreJoker bool) int {
	highCount := 0
	for k, c := range set {
		if ignoreJoker && k == 'J' {

			continue
		}
		highCount = max(highCount, c)
	}

	return highCount
}

func rankCardPart1(card uint8) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		return int(card - '0')
	}
}

func rankCardPart2(card uint8) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'T':
		return 10
	case 'J':
		return 1
	default:
		return int(card - '0')
	}
}

type cardSet map[rune]int

func newCardSet(hand string) cardSet {
	set := make(cardSet, 5)
	for _, card := range hand {
		set[card] = set[card] + 1
	}
	return set
}

func parse(file string) []camelPlay {
	scanner := aoc.OpenScanner(file)
	var plays []camelPlay
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		plays = append(plays, camelPlay{
			hand: parts[0],
			bid:  toInt(parts[1]),
		})
	}
	return plays
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return result
}
