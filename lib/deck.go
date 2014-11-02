package loev

type Deck struct {
	cards []Card
}

func NewDeck() Deck {
	return Deck{
		cards: []Card{
			Guard(), Guard(), Guard(), Guard(), Guard(),
			Priest(), Priest(),
			Baron(), Baron(),
			Handmaid(), Handmaid(),
			Prince(), Prince(),
			King(),
			Countess(),
			Princess(),
		},
	}
}

func (deck Deck) Shuffle() Deck {
	return deck
}

func (deck Deck) Draw() (Card, Deck) {
	return deck.cards[0], Deck{cards: deck.cards[1:]}
}

func (deck Deck) Length() int {
	return len(deck.cards)
}

func (deck Deck) Empty() bool {
	return len(deck.cards) == 0
}
