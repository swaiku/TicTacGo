package game

// Board represents the game grid and contains player tokens.
//   - Cells is a Size x Size matrix of *Player.
//   - ToWin defines how many aligned symbols are required to win
//     (used for variants such as 4-in-a-row or 5-in-a-row).
type Board struct {
	Cells [][]*Player
	Size  int
	ToWin int
}

// NewBoard allocates a new empty board of a given size.
// All cells start as nil (empty).
func NewBoard(size, toWin int) *Board {
	b := &Board{
		Size:  size,
		ToWin: toWin,
		Cells: make([][]*Player, size),
	}

	for i := range b.Cells {
		b.Cells[i] = make([]*Player, size)
	}
	return b
}

// Play attempts to place player p at grid coordinates (x, y).
// Returns true if the move is valid and the cell was empty.
func (b *Board) Play(p *Player, x, y int) bool {
	// Out-of-bounds protection
	if x < 0 || y < 0 || x >= b.Size || y >= b.Size {
		return false
	}
	// Cell already filled
	if b.Cells[x][y] != nil {
		return false
	}

	b.Cells[x][y] = p
	return true
}

// CheckWin verifies if a player has won.
// It checks all rows, columns, and the two main diagonals.
func (b *Board) CheckWin() *Player {
	n := b.Size

	// Check all rows and columns
	for i := 0; i < n; i++ {

		// Check row i
		if p := same(b.Cells[i]); p != nil {
			return p
		}

		// Check column i
		col := make([]*Player, n)
		for y := 0; y < n; y++ {
			col[y] = b.Cells[y][i]
		}
		if p := same(col); p != nil {
			return p
		}
	}

	// Check main diagonal
	diag1 := make([]*Player, n)
	for i := 0; i < n; i++ {
		diag1[i] = b.Cells[i][i]
	}
	if p := same(diag1); p != nil {
		return p
	}

	// Check anti-diagonal
	diag2 := make([]*Player, n)
	for i := 0; i < n; i++ {
		diag2[i] = b.Cells[i][n-1-i]
	}
	if p := same(diag2); p != nil {
		return p
	}

	return nil
}

// same returns the player occupying the whole row/column/diagonal
// if all entries are identical and non-nil.
// Otherwise, it returns nil.
func same(line []*Player) *Player {
	first := line[0]
	if first == nil {
		return nil
	}
	for _, p := range line {
		if p != first {
			return nil
		}
	}
	return first
}

// CheckDraw returns true if the board is full and no winner exists.
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

// Clear resets all cells to nil (empty board).
func (b *Board) Clear() {
	for x := range b.Cells {
		for y := range b.Cells[x] {
			b.Cells[x][y] = nil
		}
	}
}
