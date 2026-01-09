package ui

import (
	"GoTicTacToe/game"
	"GoTicTacToe/ui/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// BoardView is the visual component responsible for rendering the
// tic-tac-toe board and handling user interaction.
type BoardView struct {
	Widget // Embeds Widget: inherits size, position, anchor, AbsPosition(), etc.

	logicBoard  *game.Board      // Reference to the logical board
	OnCellClick func(cx, cy int) // Callback triggered when a cell is clicked
}

// NewBoardView creates a new visual board component.
func NewBoardView(
	board *game.Board,
	x, y, size float64,
	style utils.WidgetStyle,
	onClick func(cx, cy int),
) *BoardView {

	view := &BoardView{
		Widget: Widget{
			OffsetX: x,
			OffsetY: y,
			Width:   size,
			Height:  size,
			Anchor:  utils.AnchorCenter,
			Style:   style,
		},
		logicBoard:  board,
		OnCellClick: onClick,
	}

	// Pre-generate the grid once (static background)
	view.Widget.image = view.createGridImage()

	return view
}

// createGridImage renders the static background grid (lines + background)
// and returns the resulting image.
func (v *BoardView) createGridImage() *ebiten.Image {
	img := ebiten.NewImage(int(v.Width), int(v.Height))

	// Fill board background
	img.Fill(v.Style.BackgroundNormal)

	cellSize := v.Width / float64(v.logicBoard.Size)
	thickness := v.Style.BorderWidth
	lineColor := v.Style.BorderColor

	// Draw vertical and horizontal grid lines
	for i := 1; i < v.logicBoard.Size; i++ {
		offset := float64(i) * cellSize

		// Vertical line
		vert := ebiten.NewImage(int(thickness), int(v.Height))
		vert.Fill(lineColor)

		opv := &ebiten.DrawImageOptions{}
		opv.GeoM.Translate(offset-thickness/2, 0)
		img.DrawImage(vert, opv)

		// Horizontal line
		hori := ebiten.NewImage(int(v.Width), int(thickness))
		hori.Fill(lineColor)

		oph := &ebiten.DrawImageOptions{}
		oph.GeoM.Translate(0, offset-thickness/2)
		img.DrawImage(hori, oph)
	}

	return img
}

// Update handles mouse click detection and cell coordinate translation.
func (v *BoardView) Update() {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		vx, vy := v.AbsPosition()

		// Check if click is inside the board boundaries
		if float64(mx) >= vx && float64(mx) <= vx+v.Width &&
			float64(my) >= vy && float64(my) <= vy+v.Height {

			cellSize := v.Width / float64(v.logicBoard.Size)

			// Convert pixel coordinates → board grid coordinates
			gridX := int((float64(mx) - vx) / cellSize)
			gridY := int((float64(my) - vy) / cellSize)

			// Trigger callback
			if v.OnCellClick != nil {
				v.OnCellClick(gridX, gridY)
			}
		}
	}
}

// Draw renders the grid and the X/O symbols for each cell.
func (v *BoardView) Draw(screen *ebiten.Image) {
	vx, vy := v.AbsPosition()

	// --- Draw static grid background ---
	opGrid := &ebiten.DrawImageOptions{}
	opGrid.GeoM.Translate(vx, vy)
	screen.DrawImage(v.Widget.image, opGrid)

	cellSize := v.Width / float64(v.logicBoard.Size)
	padding := cellSize * 0.1          // 10% padding inside each cell
	usableSize := cellSize - 2*padding // Space available for the symbol

	// --- Draw all symbols (X and O) ---
	for x := 0; x < v.logicBoard.Size; x++ {
		for y := 0; y < v.logicBoard.Size; y++ {
			p := v.logicBoard.Cells[x][y]
			if p == nil || p.Symbol.Image == nil {
				continue
			}

			symbolImg := p.Symbol.Image
			srcW, srcH := symbolImg.Bounds().Dx(), symbolImg.Bounds().Dy()

			// Determine the scaling factor based on largest dimension
			maxDim := float64(srcW)
			if srcH > srcW {
				maxDim = float64(srcH)
			}
			scale := usableSize / maxDim

			// Draw the scaled symbol inside the cell
			opSym := &ebiten.DrawImageOptions{}
			opSym.Filter = ebiten.FilterLinear // Smooth scaling

			// Scale first
			opSym.GeoM.Scale(scale, scale)

			// Position inside the cell with padding
			drawX := vx + float64(x)*cellSize + padding
			drawY := vy + float64(y)*cellSize + padding

			opSym.GeoM.Translate(drawX, drawY)

			// Change the color of the image for the player color
			opSym.ColorScale.ScaleWithColor(p.Color)

			screen.DrawImage(symbolImg, opSym)
		}
	}
}
