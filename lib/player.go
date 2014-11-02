package loev

type Player struct {
	Name string
	Hand []Card
	Brain
}

func (p Player) Deal(card Card) Player {
	return Player{
		Name:  p.Name,
		Hand:  append(p.Hand, card),
		Brain: p.Brain,
	}
}

func (p Player) String() string {
	return p.Name
}
