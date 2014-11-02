package loev

type Action struct {
	Player Player
	Card   Card
	Target Player
	Turn   Turn
}

type Result struct{}
