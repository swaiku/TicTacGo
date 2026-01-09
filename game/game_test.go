package game

import "testing"

/*
TestWinDetection verifies that a simple winning condition
(3 symbols aligned in a row) is correctly detected by the game.

This test uses Board.Play directly to isolate win detection logic
without involving turn switching.
*/
func TestWinDetection(t *testing.T) {
	g := NewGame()
	p := g.Players[0]

	// Player places three symbols in the top row
	g.Board.Play(p, 0, 0)
	g.Board.Play(p, 1, 0)
	g.Board.Play(p, 2, 0)

	// The game should detect a win
	if !g.CheckWin() {
		t.Fatalf("expected win to be detected")
	}

	// The winner should be the player who placed the symbols
	if g.Winner != p {
		t.Errorf("wrong winner detected")
	}
}

//==================================================================================

/*
TestDrawDetection verifies that a full board with no winner
is correctly detected as a draw.

The move sequence fills the board completely while alternating players,
ensuring no winning alignment exists.
*/
func TestDrawDetection(t *testing.T) {
	g := NewGame()
	p1 := g.Players[0]
	p2 := g.Players[1]

	// Sequence of moves that fills the board without any win
	moves := []struct {
		p *Player
		x int
		y int
	}{
		{p1, 0, 0}, {p2, 1, 0}, {p1, 2, 0},
		{p1, 0, 1}, {p2, 1, 1}, {p1, 2, 1},
		{p2, 0, 2}, {p1, 1, 2}, {p2, 2, 2},
	}

	for _, m := range moves {
		g.Board.Play(m.p, m.x, m.y)
	}

	// The game should detect a draw
	if !g.CheckDraw() {
		t.Fatalf("expected draw to be detected")
	}
}

//==================================================================================

/*
TestPlayMoveSwitchesPlayer ensures that calling PlayMove
correctly switches the current player after a valid move. (also ensures valid moves)
*/
func TestPlayMoveSwitchesPlayer(t *testing.T) {
	g := NewGame()
	first := g.Current

	ok := g.PlayMove(0, 0)
	if !ok {
		t.Fatalf("valid move rejected")
	}

	// The current player should have changed
	if g.Current == first {
		t.Errorf("expected current player to switch")
	}
}

//==================================================================================

/*
TestResetHard verifies that ResetHard completely resets the game:
- board is cleared
- winner is cleared
- points are reset
- game state returns to PLAYING
*/
func TestResetHard(t *testing.T) {
	g := NewGame()

	// Modify the game state
	g.Players[0].Points = 5
	g.Board.Play(g.Players[0], 0, 0)

	g.ResetHard()

	if g.State != PLAYING {
		t.Errorf("expected state PLAYING after ResetHard")
	}

	if g.Winner != nil {
		t.Errorf("winner should be nil after ResetHard")
	}

	if g.Players[0].Points != 0 {
		t.Errorf("points should be reset")
	}

	// All cells should be empty again
	if len(g.Board.AvailableMoves()) != 9 {
		t.Errorf("board should be empty after ResetHard")
	}
}

//==================================================================================

/*
TestResetKeepsPoints ensures that Reset only clears the board
but does NOT reset player scores.
*/
func TestResetKeepsPoints(t *testing.T) {
	g := NewGame()
	g.Players[0].Points = 3

	g.Board.Play(g.Players[0], 0, 0)
	g.Reset()

	if g.Players[0].Points != 3 {
		t.Errorf("Reset should not reset points")
	}

	if len(g.Board.AvailableMoves()) != 9 {
		t.Errorf("board should be cleared on Reset")
	}
}

//==================================================================================

/*
TestResetPoints verifies that ResetPoints correctly resets
the score of all players without modifying the board.
*/
func TestResetPoints(t *testing.T) {
	g := NewGame()
	g.Players[0].Points = 2
	g.Players[1].Points = 4

	g.ResetPoints()

	for i, p := range g.Players {
		if p.Points != 0 {
			t.Errorf("player %d points not reset", i)
		}
	}
}

//==================================================================================

/*
TestNextPlayerWrapsAround ensures that when the current player
is the last one in the list, NextPlayer wraps back to the first player.
*/
func TestNextPlayerWrapsAround(t *testing.T) {
	g := NewGame()

	g.Current = g.Players[1]
	g.NextPlayer()

	if g.Current != g.Players[0] {
		t.Errorf("NextPlayer should wrap to first player")
	}
}

//==================================================================================

/*
TestNextPlayerFallback verifies a defensive behavior:
if Current is not found in the Players slice,
NextPlayer should safely reset it to the first player.
*/
func TestNextPlayerFallback(t *testing.T) {
	g := NewGame()

	// Assign an invalid current player
	g.Current = &Player{Name: "ghost"}

	g.NextPlayer()

	if g.Current != g.Players[0] {
		t.Errorf("fallback should reset current to first player")
	}
}

//==================================================================================

/*
TestPlayMoveInvalidDoesNotSwitchPlayer ensures that an invalid move:
- is rejected
- does NOT switch the current player
- playing in the same spot is tested here
*/
func TestPlayMoveInvalidDoesNotSwitchPlayer(t *testing.T) {
	g := NewGame()
	start := g.Current

	ok := g.PlayMove(0, 0)
	if !ok {
		t.Fatalf("first move should be valid")
	}

	// Try to play in the same cell again
	ok = g.PlayMove(0, 0)
	if ok {
		t.Errorf("move should be invalid")
	}

	// Current player should not change
	if g.Current == start {
		t.Errorf("current player should not change on invalid move")
	}
}

//==================================================================================

/*
TestDiagonalWinViaPlayMove tests a full game flow using PlayMove:
- players alternate turns
- a diagonal win is created
- the game state transitions to GAME_END
*/
func TestDiagonalWinViaPlayMove(t *testing.T) {
	g := NewGame()
	p1 := g.Players[0]

	// Sequence of moves leading to a diagonal win for p1
	moves := []struct {
		x, y int
	}{
		{0, 0}, // p1
		{0, 1}, // p2
		{1, 1}, // p1
		{0, 2}, // p2
		{2, 2}, // p1 -> win
	}

	for _, m := range moves {
		g.PlayMove(m.x, m.y)
	}

	if g.Winner != p1 {
		t.Errorf("expected diagonal win for p1")
	}

	if g.State != GAME_END {
		t.Errorf("game should be ended after win")
	}
}

//==================================================================================
