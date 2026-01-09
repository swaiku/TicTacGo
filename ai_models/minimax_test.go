package ai_models

import (
	"GoTicTacToe/game"
	"testing"
)

/*
TestMinimaxReturnsValidMove verifies that the MinimaxAI model
always returns a valid board coordinate on an empty board.

This test does NOT attempt to validate the optimality of the move
(minimax correctness), which would require complex scenario-based tests.
Instead, it ensures that:
- the algorithm terminates
- no panic occurs
- the returned move is within board bounds
*/
func TestMinimaxReturnsValidMove(t *testing.T) {
	// Create a standard empty 3x3 board
	b := game.NewBoard(3, 3)

	// Create players: one controlled by the AI, one opponent
	p1 := &game.Player{Name: "AI"}
	p2 := &game.Player{Name: "Human"}

	// Players array required by the AI interface
	players := [2]*game.Player{p1, p2}

	// Instantiate the Minimax AI model
	ai := MinimaxAI{}

	// Ask the AI to compute the next move
	x, y := ai.NextMove(b, p1, players)

	// The returned move must be inside the board boundaries
	if x < 0 || x >= 3 || y < 0 || y >= 3 {
		t.Errorf("invalid move returned: %d,%d", x, y)
	}
}
