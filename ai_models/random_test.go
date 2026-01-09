package ai_models

import (
	"GoTicTacToe/game"
	"testing"
)

/*
TestRandomAIValidMove verifies that the RandomAI model always returns
a valid board coordinate when moves are available.

The goal of this test is NOT to verify randomness or strategy,
but to ensure that:
- the AI does not crash
- the AI does not return invalid coordinates
*/
func TestRandomAIValidMove(t *testing.T) {
	// Create a standard empty 3x3 board
	b := game.NewBoard(3, 3)

	// Create a dummy player controlled by the AI
	p := &game.Player{}

	// Create the players array required by the AI interface
	players := [2]*game.Player{p, {}}

	// Instantiate the RandomAI model
	ai := RandomAI{}

	// Ask the AI to choose a move
	x, y := ai.NextMove(b, p, players)

	// The returned coordinates should be valid (non-negative)
	if x < 0 || y < 0 {
		t.Errorf("random AI returned invalid move")
	}
}
