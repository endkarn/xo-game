package xo

type XoGame struct {
	CurrentTurn int
	PlayerTurn int
	PlayerOne Player
	PlayerTwo Player
	GameBoard [9]string
}

type Player struct {
	Score int
	Symbol string
}

func NewXoGame() XoGame {
	return XoGame{
		CurrentTurn:    0,
		PlayerTurn:     0,
		PlayerOne: Player{
			Score:  0,
			Symbol: "O",
		},
		PlayerTwo: Player{
			Score:  0,
			Symbol: "X",
		},
		GameBoard:      [9]string{},
	}
}

func (xoGame *XoGame) HaveWinner() (bool , string) {
	return false , "No Winner"
}

func (xoGame *XoGame) PlayerOnePlay(location int) {
	xoGame.GameBoard[location] = xoGame.PlayerOne.Symbol
}

