package game

import "testing"

/*
TestBoardPlayAndAvailableMoves verifies the basic behavior of the Board:
- a valid move can be played
- the corresponding cell is updated
- the number of available moves decreases accordingly
*/
func TestBoardPlayAndAvailableMoves(t *testing.T) {
	// Create a standard 3x3 board where 3 aligned symbols are needed to win
	b := NewBoard(3, 3)

	// Create a test player
	p := &Player{Name: "P1"}

	// Play a valid move in the center of the board
	ok := b.Play(p, 1, 1)
	if !ok {
		t.Fatalf("expected move to be valid")
	}

	// The cell (1,1) should now contain the player pointer
	if b.Cells[1][1] != p {
		t.Errorf("cell not updated correctly")
	}

	// After one move on a 3x3 board, 8 moves should remain available
	moves := b.AvailableMoves()
	if len(moves) != 8 {
		t.Errorf("expected 8 available moves, got %d", len(moves))
	}
}

//==================================================================================

/*
TestBoardInvalidMove verifies that the Board correctly rejects invalid moves,
such as attempting to play on an already occupied cell.
*/
func TestBoardInvalidMove(t *testing.T) {
	// Create a new empty board
	b := NewBoard(3, 3)

	// Create a test player
	p := &Player{Name: "P1"}

	// Play a move in the top-left corner
	b.Play(p, 0, 0)

	// Attempt to play in the same cell again
	ok := b.Play(p, 0, 0)

	// The second move should be rejected
	if ok {
		t.Errorf("expected move to be rejected")
	}
}

//==================================================================================
