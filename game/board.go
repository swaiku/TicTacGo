package game

// Board represents the game grid and contains player tokens.
// - Cells is a Size x Size matrix of *Player.
// - ToWin defines how many aligned symbols are required to win
//   (used for variants such as 4-in-a-row or 5-in-a-row).
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

// CheckWin verifies if a player has won for any streak of length ToWin
// horizontally, vertically, or diagonally.
func (b *Board) CheckWin() *Player {
	target := b.ToWin
	if target <= 0 || target > b.Size {
		target = b.Size
	}

	directions := [][2]int{
		{1, 0},  // horizontal
		{0, 1},  // vertical
		{1, 1},  // diagonal
		{1, -1}, // anti-diagonal
	}

	for x := 0; x < b.Size; x++ {
		for y := 0; y < b.Size; y++ {
			start := b.Cells[x][y]
			if start == nil {
				continue
			}

			for _, dir := range directions {
				count := 1
				for step := 1; step < target; step++ {
					nx := x + dir[0]*step
					ny := y + dir[1]*step
					if nx < 0 || ny < 0 || nx >= b.Size || ny >= b.Size {
						break
					}
					if b.Cells[nx][ny] != start {
						break
					}
					count++
				}
				if count == target {
					return start
				}
			}
		}
	}

	return nil
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
