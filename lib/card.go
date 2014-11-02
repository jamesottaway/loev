package loev

type Card struct {
	Number int
}

func Guard() Card {
	return Card{Number: 1}
}

func Priest() Card {
	return Card{Number: 2}
}

func Baron() Card {
	return Card{Number: 3}
}

func Handmaid() Card {
	return Card{Number: 4}
}

func Prince() Card {
	return Card{Number: 5}
}

func King() Card {
	return Card{Number: 6}
}

func Countess() Card {
	return Card{Number: 7}
}

func Princess() Card {
	return Card{Number: 8}
}
