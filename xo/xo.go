package xo

type XoGame struct {
	CurrentTurn int
	PlayerTurn int
	PlayerOneScore int
	PlayerTwoScore int
	GameBoard [3][3]string
}

func NewXoGame() XoGame {
	return XoGame{
		CurrentTurn:    0,
		PlayerTurn:     0,
		PlayerOneScore: 0,
		PlayerTwoScore: 0,
		GameBoard:      [3][3]string{},
	}
}

func (xoGame XoGame) HaveWinner() (bool , string) {
	return false , "No Winner"
}