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

func TestShuffleReturnsDeckWithTheSameCardsButInDifferentOrder() (t *testing.T) {
	deck := loev.NewDeck()
	shuffled := deck.Shuffle()

	for !deck.Empty() {
	}

	for !shuffled.Empty() {
	}
}

func TestShuffleIsNonDeterministicAndReturnsDifferentOrderEachTime() (t *testing.T) {
	deck := loev.NewDeck()
	shuffleOne := deck.Shuffle()
	shuffleTwo := deck.Shuffle()

	for !shuffleOne.Empty() {
	}

	for !shuffleTwo.Empty() {
	}
}
