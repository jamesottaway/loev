package loev_test

import (
	"loev/lib"
	"testing"
)

func TestNewDeckHasSixteenCards(t *testing.T) {
	deck := loev.NewDeck()
	cards := make([]loev.Card, 0)

	for !deck.Empty() {
		card, remainder := deck.Draw()
		deck = remainder
		cards = append(cards, card)
	}

	if len(cards) != 16 {
		t.Errorf("Expected 16 cards, but got %d", len(cards))
	}
}

func TestDrawReturnsTheRemainingDeck(t *testing.T) {
	deck := loev.NewDeck()
	_, next := deck.Draw()

	if deck.Length() != next.Length()+1 {
		t.Errorf("Original deck had %d cards before draw, but %d after draw", deck.Length(), next.Length())
	}
}

func TestShuffleReturnsDeckWithTheSameCardsButInDifferentOrder(t *testing.T) {
	deck := loev.NewDeck()
	shuffled := deck.Shuffle()

	unshuffledCards := make([]loev.Card, 0)
	shuffledCards := make([]loev.Card, 0)

	for !deck.Empty() {
		card, remaining := deck.Draw()
		deck = remaining
		unshuffledCards = append(unshuffledCards, card)
	}

	for !shuffled.Empty() {
		card, remaining := shuffled.Draw()
		shuffled = remaining
		shuffledCards = append(shuffledCards, card)
	}

	matching := 0
	for i, card := range shuffledCards {
		if card == unshuffledCards[i] {
			matching += 1
		}
	}

	if matching >= len(shuffledCards)/2 {
		t.Errorf("Expected less that %d cards to match after shuffling, but was %d", len(shuffledCards)/2, matching)
	}
}

func TestShuffleIsNonDeterministicAndReturnsDifferentOrderEachTime(t *testing.T) {
	deck := loev.NewDeck()
	shuffleOne := deck.Shuffle()
	shuffleTwo := deck.Shuffle()

	firstShuffle := make([]loev.Card, 0)
	secondShuffle := make([]loev.Card, 0)

	for !shuffleOne.Empty() {
		card, remaining := shuffleOne.Draw()
		shuffleOne = remaining
		firstShuffle = append(firstShuffle, card)
	}

	for !shuffleTwo.Empty() {
		card, remaining := shuffleTwo.Draw()
		shuffleTwo = remaining
		secondShuffle = append(secondShuffle, card)
	}

	matching := 0
	for i, card := range firstShuffle {
		if card == secondShuffle[i] {
			matching += 1
		}
	}

	if matching >= len(secondShuffle)/2 {
		t.Errorf("Expected less that %d cards to match after shuffling, but was %d", len(secondShuffle)/2, matching)
	}
}
