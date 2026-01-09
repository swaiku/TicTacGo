// Package game implements the core Tic-Tac-Toe game logic including
// board state, player management, win detection, and move validation.
package game

// Move represents a single move on the game board as grid coordinates.
// X is the column index (0-based, left to right).
// Y is the row index (0-based, top to bottom).
type Move struct {
	X int
	Y int
}

// NewMove creates a Move at the specified grid coordinates.
func NewMove(x, y int) Move {
	return Move{X: x, Y: y}
}

// IsValid checks if the move coordinates are within the given board dimensions.
func (m Move) IsValid(boardWidth, boardHeight int) bool {
	return m.X >= 0 && m.X < boardWidth && m.Y >= 0 && m.Y < boardHeight
}
