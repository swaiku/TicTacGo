package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const boardThickness = 3

type Board struct {
	Cells [][]*Player
	Size  int
	ToWin int
}

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

// Dessin de la grille
func (b *Board) GenerateImage() *ebiten.Image {
	img := ebiten.NewImage(480, 480)
	img.Fill(color.RGBA{20, 20, 20, 255})

	cell := 480 / b.Size

	// Lignes verticales et horizontales
	lineColor := color.RGBA{200, 200, 200, 255}
	for i := 1; i < b.Size; i++ {

		// vertical
		v := ebiten.NewImage(boardThickness, 480)
		v.Fill(lineColor)
		opv := &ebiten.DrawImageOptions{}
		opv.GeoM.Translate(float64(i*cell), 0)
		img.DrawImage(v, opv)

		// horizontal
		h := ebiten.NewImage(480, boardThickness)
		h.Fill(lineColor)
		oph := &ebiten.DrawImageOptions{}
		oph.GeoM.Translate(0, float64(i*cell))
		img.DrawImage(h, oph)
	}

	return img
}

// Place un symbole
func (b *Board) Play(p *Player, x, y int) bool {
	if x < 0 || y < 0 || x >= b.Size || y >= b.Size {
		return false
	}
	if b.Cells[x][y] != nil {
		return false
	}

	b.Cells[x][y] = p
	return true
}

// Vérifie victoire
func (b *Board) CheckWin() *Player {
	n := b.Size

	// Rows/Cols
	for i := 0; i < n; i++ {
		// Row
		if p := same(b.Cells[i]); p != nil {
			return p
		}

		// Col
		col := make([]*Player, n)
		for y := 0; y < n; y++ {
			col[y] = b.Cells[y][i]
		}
		if p := same(col); p != nil {
			return p
		}
	}

	// Diag 1
	diag1 := make([]*Player, n)
	for i := 0; i < n; i++ {
		diag1[i] = b.Cells[i][i]
	}
	if p := same(diag1); p != nil {
		return p
	}

	// Diag 2
	diag2 := make([]*Player, n)
	for i := 0; i < n; i++ {
		diag2[i] = b.Cells[i][n-1-i]
	}
	if p := same(diag2); p != nil {
		return p
	}

	return nil
}

func same(row []*Player) *Player {
	first := row[0]
	if first == nil {
		return nil
	}
	for _, p := range row {
		if p != first {
			return nil
		}
	}
	return first
}

// Check draw
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
