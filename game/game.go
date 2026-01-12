package game

import (
	"GoTicTacToe/assets"
	"image/color"
)

// GameState represents the current phase of a game.
type GameState int

const (
	// StatePlaying indicates a game is currently in progress.
	StatePlaying GameState = iota
	// StateGameEnd indicates the game has finished (either win or draw).
	StateGameEnd
)

// Default board configuration for a standard 3x3 Tic-Tac-Toe game.
const (
	DefaultBoardWidth  = 3 // Default number of columns
	DefaultBoardHeight = 3 // Default number of rows
	DefaultToWin       = 3 // Default symbols needed to win (3-in-a-row)
)

// Game orchestrates the game state, players, and board interactions.
// It manages turn order, win/draw detection, and score tracking.
type Game struct {
	State   GameState // Current game phase (playing or ended)
	Board   *Board    // The game board containing cell states
	Players []*Player // All players participating in the game
	Current *Player   // The player whose turn it currently is
	Winner  *Player   // The winner of the current round (nil if draw or ongoing)

	boardWidth  int // Configured board width for resets
	boardHeight int // Configured board height for resets
	toWin       int // Configured win condition for resets
}

// NewGame creates a new Game with default 3x3 configuration and two players.
func NewGame() *Game {
	return NewGameWithConfig(DefaultBoardWidth, DefaultBoardHeight, DefaultToWin, nil)
}

// NewGameWithConfig creates a Game with custom board dimensions and players.
// If players is nil or empty, default players (Circle and Cross) are created.
func NewGameWithConfig(boardWidth, boardHeight, toWin int, players []*Player) *Game {
	g := &Game{}
	g.ResetHardWithPlayers(boardWidth, boardHeight, toWin, players)
	return g
}

// ResetHard performs a complete reset with default configuration.
// This clears all scores and reinitializes with default players.
func (g *Game) ResetHard() {
	g.ResetHardWithPlayers(DefaultBoardWidth, DefaultBoardHeight, DefaultToWin, nil)
}

// ResetHardWithPlayers performs a complete reset with custom configuration.
// All player scores are reset to zero.
func (g *Game) ResetHardWithPlayers(boardWidth, boardHeight, toWin int, players []*Player) {
	g.boardWidth = boardWidth
	g.boardHeight = boardHeight
	g.toWin = toWin
	g.Board = NewBoard(boardWidth, boardHeight, toWin)

	if len(players) == 0 {
		players = g.createDefaultPlayers()
	}

	g.Players = players
	g.resetAllPlayerScores()

	g.Current = g.Players[0]
	g.Winner = nil
	g.State = StatePlaying
}

// createDefaultPlayers returns the standard two-player setup.
func (g *Game) createDefaultPlayers() []*Player {
	return []*Player{
		NewPlayer(assets.NewSymbol(assets.CircleSymbol), color.RGBA{R: 255, G: 0, B: 0, A: 255}),
		NewPlayer(assets.NewSymbol(assets.CrossSymbol), color.RGBA{R: 0, G: 0, B: 255, A: 255}),
	}
}

// resetAllPlayerScores sets all player scores to zero.
func (g *Game) resetAllPlayerScores() {
	for _, player := range g.Players {
		player.Points = 0
	}
}

// Reset clears the board and restarts the game while preserving player scores.
// Use this between rounds in a multi-round match.
func (g *Game) Reset() {
	if g.Board == nil {
		g.Board = NewBoard(g.boardWidth, g.boardHeight, g.toWin)
	} else {
		g.Board.Clear()
	}

	g.Current = g.Players[0]
	g.Winner = nil
	g.State = StatePlaying
}

// ResetPoints sets all player scores to zero without affecting the current game state.
func (g *Game) ResetPoints() {
	g.resetAllPlayerScores()
}

// NextPlayer advances the turn to the next player in rotation.
// Players are cycled in the order they appear in the Players slice.
func (g *Game) NextPlayer() {
	if len(g.Players) == 0 {
		return
	}

	for i, player := range g.Players {
		if player == g.Current {
			g.Current = g.Players[(i+1)%len(g.Players)]
			return
		}
	}

	// Fallback: current player not found in list, reset to first
	g.Current = g.Players[0]
}

// PlayMove attempts to execute a move at coordinates (x, y) for the current player.
// Returns true if the move was valid and executed successfully.
// After a valid move, the game checks for win/draw conditions and advances the turn.
func (g *Game) PlayMove(x, y int) bool {
	if !g.Board.Play(g.Current, x, y) {
		return false
	}

	if g.CheckWin() {
		return true
	}

	if g.CheckDraw() {
		return true
	}

	g.NextPlayer()
	return true
}

// CheckWin checks if the current board state contains a winning alignment.
// If a winner is found, updates the game state, increments their score,
// and sets the State to StateGameEnd.
func (g *Game) CheckWin() bool {
	winner := g.Board.CheckWin()
	if winner == nil {
		return false
	}

	g.Winner = winner
	g.Winner.Points++
	g.State = StateGameEnd
	return true
}

// CheckDraw checks if the game is a draw (board full with no winner).
// If a draw is detected, sets the State to StateGameEnd with no winner.
func (g *Game) CheckDraw() bool {
	if !g.Board.CheckDraw() {
		return false
	}

	g.Winner = nil
	g.State = StateGameEnd
	return true
}

// IsPlaying returns true if the game is currently in progress.
func (g *Game) IsPlaying() bool {
	return g.State == StatePlaying
}

// IsGameEnd returns true if the game has finished.
func (g *Game) IsGameEnd() bool {
	return g.State == StateGameEnd
}
