package xo

import "testing"

func Test_HaveWinner_Input_Turn_0_Should_Be_False_And_NoWinner_And_BoardEmpty(t *testing.T){
	expectedHaveWinner, expectedWinner := false, "No Winner"
	expectedBoard := [9]string{
		"","","",
		"","","",
		"","",""}
	xoGame := NewXoGame()

	actualHaveWinner , actualWinner := xoGame.HaveWinner()
	actualBoard := xoGame.GameBoard

	if expectedBoard != actualBoard {
		t.Errorf("Expected %v but got %v ",expectedBoard,actualBoard)
	}
	if expectedHaveWinner != actualHaveWinner {
		t.Errorf("Expected %v but got %v ",expectedHaveWinner,actualHaveWinner)
	}
	if expectedWinner != actualWinner {
		t.Errorf("Expected %v but got %v ",expectedHaveWinner,actualHaveWinner)
	}
}

func Test_HaveWinner_Input_Turn_1_PlayerOnePlay_0_1_Should_Be_False_And_NoWinner(t *testing.T){
	expectedHaveWinner, expectedWinner := false, "No Winner"
	expectedBoard := [9]string{
		"O","","",
		"","","",
		"","",""}
	xoGame := NewXoGame()

	xoGame.PlayerOnePlay(0)
	actualHaveWinner , actualWinner := xoGame.HaveWinner()
	actualBoard := xoGame.GameBoard

	if expectedBoard != actualBoard {
		t.Errorf("Expected %v but got %v ",expectedBoard,actualBoard)
	}
	if expectedHaveWinner != actualHaveWinner {
		t.Errorf("Expected %v but got %v ",expectedHaveWinner,actualHaveWinner)
	}
	if expectedWinner != actualWinner {
		t.Errorf("Expected %v but got %v ",expectedHaveWinner,actualHaveWinner)
	}
}