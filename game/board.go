// Package game implements the core Tic-Tac-Toe game logic including
// board state, player management, win detection, and move validation.

package game

// Board represents the game grid containing player tokens.
//
// The board uses a column-major layout where Cells[x][y] accesses
// column x, row y. Width represents the number of columns and Height
// the number of rows. ToWin defines how many consecutive symbols
// are required to win (supports N-in-a-row variants).
type Board struct {
	Cells  [][]*Player // 2D grid of player references (nil = empty cell)
	Width  int         // Number of columns
	Height int         // Number of rows
	ToWin  int         // Required consecutive symbols to win
}

// Direction represents a 2D step vector (dx, dy) used for line scanning
// during win detection.
type Direction struct {
	DX int // Horizontal step (-1, 0, or 1)
	DY int // Vertical step (-1, 0, or 1)
}

// winDirections contains the four directions to scan for winning alignments.
//
// Only four directions are needed because we scan from each cell as the
// starting point, so reverse directions would be redundant:
//   - Horizontal (→): checks left-to-right alignments
//   - Vertical (↓): checks top-to-bottom alignments
//   - Diagonal down-right (↘): checks descending diagonal
//   - Diagonal up-right (↗): checks ascending diagonal
var winDirections = [...]Direction{
	{DX: 1, DY: 0},  // horizontal (→)
	{DX: 0, DY: 1},  // vertical (↓)
	{DX: 1, DY: 1},  // diagonal down-right (↘)
	{DX: 1, DY: -1}, // diagonal up-right (↗)
}

// Win detection constants.
const (
	// initialStreakCount is the starting count when considering the origin cell.
	initialStreakCount = 1

	// firstStep is the first step index away from the starting cell.
	firstStep = 1
)

// NewBoard allocates and returns a new empty board with the given dimensions.
//
// All cells are initialized to nil (empty). The toWin parameter specifies
// how many consecutive symbols are required to win.
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
//
// Returns true if the move was valid and the cell was empty.
// Returns false if the coordinates are out of bounds or the cell is occupied.
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

// CheckWin scans the board for a winning alignment of ToWin consecutive symbols.
//
// The algorithm iterates through each cell as a potential starting point,
// then scans in four directions (horizontal, vertical, two diagonals) to
// count consecutive symbols belonging to the same player.
//
// Returns the winning player if found, or nil if no winner exists.
// If ToWin is invalid (<= 0 or larger than the smallest board dimension),
// it is clamped to the smallest dimension for robustness.
func (b *Board) CheckWin() *Player {
	target := b.effectiveToWin()

	for x := 0; x < b.Width; x++ {
		for y := 0; y < b.Height; y++ {
			startPlayer := b.Cells[x][y]
			if startPlayer == nil {
				continue
			}

			for _, dir := range winDirections {
				if b.checkLineWin(x, y, dir, startPlayer, target) {
					return startPlayer
				}
			}
		}
	}

	return nil
}

// checkLineWin checks if there are 'target' consecutive symbols starting
// from (x, y) in the given direction, all belonging to the same player.
func (b *Board) checkLineWin(x, y int, dir Direction, player *Player, target int) bool {
	count := initialStreakCount

	for step := firstStep; step < target; step++ {
		nx := x + dir.DX*step
		ny := y + dir.DY*step

		if !b.inBounds(nx, ny) {
			return false
		}
		if b.Cells[nx][ny] != player {
			return false
		}
		count++
	}

	return count == target
}

// CheckDraw returns true if the board is full (no empty cells remain).
//
// Note: The game loop should call CheckWin first; this method does not
// verify whether a winner exists.
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
//
// This is primarily used by AI models to enumerate valid moves.
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
//
// Players are referenced (not cloned) because they are treated as immutable
// identity objects. Only the board state (cell assignments) is duplicated.
// This is essential for AI algorithms like Minimax that simulate moves.
func (b *Board) Clone() *Board {
	clone := NewBoard(b.Width, b.Height, b.ToWin)

	for x := 0; x < b.Width; x++ {
		for y := 0; y < b.Height; y++ {
			clone.Cells[x][y] = b.Cells[x][y]
		}
	}
	return clone
}

// isValidPosition returns true if (x, y) is within the board boundaries.
func (b *Board) isValidPosition(x, y int) bool {
	return b.inBounds(x, y)
}

// inBounds returns true if (x, y) is within the board limits.
func (b *Board) inBounds(x, y int) bool {
	return x >= 0 && y >= 0 && x < b.Width && y < b.Height
}

// effectiveToWin returns a valid, clamped value for ToWin.
//
// If ToWin is not usable (<=0 or exceeds the smallest dimension),
// it is clamped to min(Width, Height). This ensures robustness
// when used with board variants or unexpected inputs.
func (b *Board) effectiveToWin() int {
	target := b.ToWin
	minDim := b.Width
	if b.Height < minDim {
		minDim = b.Height
	}

	if target <= 0 || target > minDim {
		return minDim
	}
	return target
}
