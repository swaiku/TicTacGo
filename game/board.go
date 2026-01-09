package game

// Board represents the game grid and tracks which player occupies each cell.
// The grid uses column-major indexing: Cells[x][y] where x is the column and y is the row.
type Board struct {
	Cells  [][]*Player // 2D grid of players; nil indicates an empty cell
	Width  int         // Number of columns in the grid
	Height int         // Number of rows in the grid
	ToWin  int         // Number of aligned symbols required to win
}

// winCheckDirections defines the four directions to check for winning alignments:
// horizontal, vertical, diagonal down-right, and diagonal up-right.
var winCheckDirections = [][2]int{
	{1, 0},  // Horizontal: check cells to the right
	{0, 1},  // Vertical: check cells downward
	{1, 1},  // Diagonal: check cells down-right
	{1, -1}, // Diagonal: check cells up-right
}

// NewBoard creates an empty board with the specified dimensions.
// All cells are initialized to nil (empty).
func NewBoard(width, height, toWin int) *Board {
	cells := make([][]*Player, width)
	for x := range cells {
		cells[x] = make([]*Player, height)
	}

	return &Board{
		Width:  width,
		Height: height,
		ToWin:  toWin,
		Cells:  cells,
	}
}

// Play attempts to place the player's mark at grid coordinates (x, y).
// Returns true if the move was valid and the cell was empty.
// Returns false if the coordinates are out of bounds or the cell is already occupied.
func (b *Board) Play(player *Player, x, y int) bool {
	if !b.isValidPosition(x, y) {
		return false
	}
	if b.Cells[x][y] != nil {
		return false
	}

	b.Cells[x][y] = player
	return true
}

// isValidPosition checks if coordinates are within board boundaries.
func (b *Board) isValidPosition(x, y int) bool {
	return x >= 0 && y >= 0 && x < b.Width && y < b.Height
}

// CheckWin scans the board for a winning alignment of ToWin consecutive symbols.
// Returns the winning player, or nil if no winner exists yet.
//
// The algorithm checks every cell as a potential starting point and looks
// in four directions (horizontal, vertical, and both diagonals) for a
// consecutive sequence of the same player's marks.
func (b *Board) CheckWin() *Player {
	requiredToWin := b.effectiveToWin()

	for x := 0; x < b.Width; x++ {
		for y := 0; y < b.Height; y++ {
			startPlayer := b.Cells[x][y]
			if startPlayer == nil {
				continue
			}

			if b.hasWinningLineFrom(x, y, startPlayer, requiredToWin) {
				return startPlayer
			}
		}
	}

	return nil
}

// effectiveToWin returns the ToWin value clamped to valid bounds.
// ToWin cannot exceed the smallest board dimension.
func (b *Board) effectiveToWin() int {
	toWin := b.ToWin
	minDimension := b.Width
	if b.Height < minDimension {
		minDimension = b.Height
	}

	if toWin <= 0 || toWin > minDimension {
		return minDimension
	}
	return toWin
}

// hasWinningLineFrom checks if there's a winning line starting from (x, y)
// in any of the four check directions.
func (b *Board) hasWinningLineFrom(x, y int, player *Player, requiredToWin int) bool {
	for _, dir := range winCheckDirections {
		consecutiveCount := 1

		for step := 1; step < requiredToWin; step++ {
			nextX := x + dir[0]*step
			nextY := y + dir[1]*step

			if !b.isValidPosition(nextX, nextY) {
				break
			}
			if b.Cells[nextX][nextY] != player {
				break
			}
			consecutiveCount++
		}

		if consecutiveCount >= requiredToWin {
			return true
		}
	}
	return false
}

// CheckDraw returns true if the board is completely filled with no empty cells.
// Note: This should be called after CheckWin to ensure no winner exists.
func (b *Board) CheckDraw() bool {
	for x := range b.Cells {
		for y := range b.Cells[x] {
			if b.Cells[x][y] == nil {
				return false
			}
		}
	}
	return true
}

// Clear resets all cells to nil, returning the board to its initial empty state.
func (b *Board) Clear() {
	for x := range b.Cells {
		for y := range b.Cells[x] {
			b.Cells[x][y] = nil
		}
	}
}

// AvailableMoves returns a slice of all empty cell positions on the board.
// Useful for AI move selection and validation.
func (b *Board) AvailableMoves() []Move {
	moves := make([]Move, 0, b.Width*b.Height)

	for x := 0; x < b.Width; x++ {
		for y := 0; y < b.Height; y++ {
			if b.Cells[x][y] == nil {
				moves = append(moves, Move{X: x, Y: y})
			}
		}
	}
	return moves
}

// Clone creates a deep copy of the board.
// The new board shares Player pointers but has independent cell storage.
// Useful for AI algorithms that need to simulate moves without affecting the game state.
func (b *Board) Clone() *Board {
	clone := NewBoard(b.Width, b.Height, b.ToWin)

	for x := 0; x < b.Width; x++ {
		for y := 0; y < b.Height; y++ {
			clone.Cells[x][y] = b.Cells[x][y]
		}
	}
	return clone
}

// IsFull returns true if no empty cells remain on the board.
func (b *Board) IsFull() bool {
	return len(b.AvailableMoves()) == 0
}

// CellAt returns the player occupying the cell at (x, y), or nil if empty.
// Returns nil for out-of-bounds coordinates.
func (b *Board) CellAt(x, y int) *Player {
	if !b.isValidPosition(x, y) {
		return nil
	}
	return b.Cells[x][y]
}
