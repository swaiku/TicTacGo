// Package ai_models provides AI strategies for playing Tic-Tac-Toe.
//
// This package implements the AIModel interface and provides multiple
// AI implementations with varying difficulty levels:
//   - RandomAI: Selects moves randomly (easy difficulty)
//   - MinimaxAI: Uses the Minimax algorithm for optimal play (hard difficulty)
package ai_models

import "GoTicTacToe/game"

// AIModel defines the interface for AI player strategies.
//
// Implementations must provide a NextMove method that analyzes the current
// board state and returns the coordinates of the chosen move.
type AIModel interface {
	// NextMove selects the best move for the given player.
	//
	// Parameters:
	//   - board: the current game board state
	//   - me: the player for whom to compute the move
	//   - players: all players in the game (for opponent analysis)
	//
	// Returns the (x, y) coordinates of the chosen cell.
	// Returns (-1, -1) if no valid move is available.
	NextMove(board *game.Board, me *game.Player, players []*game.Player) (x, y int)
}
