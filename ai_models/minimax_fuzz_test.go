package ai_models

import (
	"GoTicTacToe/game"
	"testing"
)

/*
FuzzMinimaxDoesNotCrash performs fuzz testing on the MinimaxAI algorithm.

The goal of this test is NOT to verify the quality of the moves,
but to ensure that the Minimax implementation:
- never panics
- always terminates
- behaves safely for a range of board sizes

The fuzzing engine generates random board sizes within a constrained range
to explore unexpected execution paths.
*/
func FuzzMinimaxDoesNotCrash(f *testing.F) {

	// Seed corpus: start fuzzing with a valid board size
	f.Add(3)

	// Fuzz function executed with randomly generated inputs
	f.Fuzz(func(t *testing.T, size int) {

		// Ignore sizes that would produce invalid or unsupported boards
		// This keeps the fuzzing focused on realistic game scenarios
		if size < 3 || size > 5 {
			return
		}

		// Create a board with the fuzz-generated size
		b := game.NewBoard(size, size)

		// Create AI player and opponent
		p1 := &game.Player{Name: "AI"}
		p2 := &game.Player{Name: "Human"}

		// Players array required by the AI interface
		players := [2]*game.Player{p1, p2}

		// Instantiate the Minimax AI model
		ai := MinimaxAI{}

		// Ask the AI to compute a move
		x, y := ai.NextMove(b, p1, players)

		// The AI should never return absurd values or crash
		// Even in edge cases, coordinates must remain within safe bounds
		if x < -1 || y < -1 {
			t.Fatalf("invalid move")
		}
	})
}
