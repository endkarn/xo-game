package xo

import "testing"

func Test_HaveWinner_Input_Turn_0_Should_Be_False_And_NoWinner(t *testing.T){
	expectedHaveWinner, expectedWinner := false, "No Winner"
	xoGame := NewXoGame()

	actualHaveWinner , actualWinner := xoGame.HaveWinner()

	if expectedHaveWinner != actualHaveWinner {
		t.Errorf("Expected %v but got %v ",expectedHaveWinner,actualHaveWinner)
	}
	if expectedWinner != actualWinner {
		t.Errorf("Expected %v but got %v ",expectedHaveWinner,actualHaveWinner)
	}

}