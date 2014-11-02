package main

import "loev/lib"

func main() {
	p1 := loev.Player{
		Name:  "James",
		Hand:  make([]loev.Card, 0, 2),
		Brain: loev.Brain{},
	}
	p2 := loev.Player{
		Name:  "Nick",
		Hand:  make([]loev.Card, 0, 2),
		Brain: loev.Brain{},
	}
	deck := loev.NewDeck()
	game := loev.NewGame([]loev.Player{p1, p2}, deck.Shuffle())
	game = game.Start()
	for !game.Complete() {
		game = game.Tick()
	}
}
