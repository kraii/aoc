package four

import (
	"aoc"
	"slices"
	"strconv"
	"strings"
)

type scratchCard struct {
	id             int
	winningNumbers []string
	numbers        []string
}

func solveB(file string) int {
	total := 0
	originals := readCards(file)
	cardCount := len(originals)
	// Allocating a stupid big capacity stops us from reallocating. Could probably refine the slice logic but...
	scoreQueue := make([]*scratchCard, cardCount, cardCount*100000)
	scoreMemo := makeMemo(cardCount)
	copy(scoreQueue, originals)

	for len(scoreQueue) > 0 {
		total++
		card := scoreQueue[0]
		scoreQueue = scoreQueue[1:]

		cardScore := score(scoreMemo, card)

		for i := card.id; i <= card.id-1+cardScore && i < cardCount; i++ {
			scoreQueue = append(scoreQueue, originals[i])
		}

	}
	return total
}

func makeMemo(length int) []int {
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = -1
	}
	return result
}

func score(memoScores []int, card *scratchCard) int {
	cardIndex := card.id - 1
	prev := memoScores[cardIndex]
	if prev >= 0 {
		return prev
	}

	score := 0
	for _, n := range card.numbers {
		if slices.Contains(card.winningNumbers, n) {
			score++
		}
	}
	memoScores[cardIndex] = score
	return score
}

func readCards(file string) []*scratchCard {
	var cards []*scratchCard
	scanner := aoc.OpenScanner(file)

	for scanner.Scan() {
		cards = append(cards, parseCard(scanner.Text()))
	}

	return cards
}

func parseCard(line string) *scratchCard {
	split := strings.Split(line, ":")
	title, withoutTitle := split[0], split[1]
	split = strings.Split(withoutTitle, "|")

	id, err := strconv.Atoi(strings.Fields(title)[1])
	if err != nil {
		panic(err)
	}
	winningNumbers := strings.Fields(split[0])
	cardNumbers := strings.Fields(split[1])

	return &scratchCard{
		id:             id,
		winningNumbers: winningNumbers,
		numbers:        cardNumbers,
	}
}
