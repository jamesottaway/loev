package loev

type Game struct {
	next      int
	players   []Player
	actions   []Action
	deck      Deck
	burntCard Card
}

func NewGame(players []Player, deck Deck) Game {
	return Game{
		next:    0,
		players: players,
		actions: make([]Action, 0, deck.Length()),
		deck:    deck,
	}
}

func (game Game) Start() Game {
	burn, deck := game.deck.Draw()
	players := make([]Player, 0, len(game.players))
	for _, player := range game.players {
		card, d := deck.Draw()
		deck = d
		players = append(players, player.Deal(card))
	}

	return Game{
		next:      0,
		players:   players,
		actions:   game.actions,
		deck:      deck,
		burntCard: burn,
	}
}

func (game Game) Tick() Game {
	card, deck := game.deck.Draw()
	player := game.players[game.next].Deal(card)
	action := player.Brain.Move(game.stateFor(player))
	game, result, err := game.tick(action)
	if err != nil {
		panic("invalid action")
	}
	player.Brain.Notify(action, result)
	players := game.players
	players[game.next] = player
	newGame := Game{
		next:      game.next + 1,
		players:   players,
		actions:   append(game.actions, action),
		deck:      deck,
		burntCard: game.burntCard,
	}
	return newGame
}

func (game Game) tick(action Action) (Game, Result, error) {
	return game, Result{}, nil
}

func (game Game) stateFor(player Player) State {
	return State{
		Hand:    player.Hand,
		Actions: game.actions,
	}
}

func (game Game) Validate(action Action) bool {
	return true
}

func (game Game) Complete() bool {
	return false
}
