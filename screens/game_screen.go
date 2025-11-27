package screens

import (
	"GoTicTacToe/assets"
	"GoTicTacToe/game"
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type GameScreen struct {
	host      ScreenHost
	game      *game.Game
	gameImage *ebiten.Image
}

const (
	boardSize = 3
	cellSize  = 160
)

func NewGameScreen(h ScreenHost) *GameScreen {
	g := game.NewGame()

	return &GameScreen{
		host:      h,
		game:      g,
		gameImage: ebiten.NewImage(boardSize*cellSize, boardSize*cellSize),
	}
}

func (gs *GameScreen) Update() error {
	switch gs.game.State {

	case game.PLAYING:
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			boardX, boardY := gs.game.GetCursorBoardPos(x, y)

			if gs.game.PlayMove(boardX, boardY) {
				// Win
				if gs.game.CheckWin() {
					gs.game.State = game.GAME_END
				} else if gs.game.CheckDraw() {
					gs.game.State = game.GAME_END
					gs.game.Winner = nil
				}
			}
		}

	case game.GAME_END:
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			gs.game.Reset()
		}
	}

	// Hotkeys
	if inpututil.KeyPressDuration(ebiten.KeyEscape) == 60 {
		os.Exit(0)
	}
	if inpututil.KeyPressDuration(ebiten.KeyR) == 60 {
		gs.game.Reset()
		gs.game.ResetPoints()
	}

	return nil
}

func (gs *GameScreen) Draw(screen *ebiten.Image) {
	gs.gameImage.Clear()

	// Draw board grid
	screen.DrawImage(gs.game.BoardImg, nil)

	// Draw symbols
	for y := 0; y < gs.game.Board.Size; y++ {
		for x := 0; x < gs.game.Board.Size; x++ {
			player := gs.game.Board.Cells[x][y]
			if player != nil {
				gs.drawSymbol(x, y, player, screen)
			}
		}
	}

	// Debug info
	gs.drawDebug(screen)

	// Score
	gs.drawScore(screen)

	// Winning message
	if gs.game.State == game.GAME_END {
		gs.drawEndMessage(screen)
	}
}

//
// === Drawing Helpers ===
//

func (gs *GameScreen) drawSymbol(gridX, gridY int, p *game.Player, screen *ebiten.Image) {
	padding := cellSize * 0.1
	usable := cellSize - 2*padding
	scale := usable / float64(cellSize)

	px := float64(gridX)*cellSize + padding
	py := float64(gridY)*cellSize + padding

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(px, py)

	screen.DrawImage(p.Symbol.Image, op)
}

func (gs *GameScreen) drawDebug(screen *ebiten.Image) {
	mx, my := ebiten.CursorPosition()
	msg := fmt.Sprintf("TPS: %.2f | FPS: %.2f | Mouse: %d,%d",
		ebiten.ActualTPS(), ebiten.ActualFPS(), mx, my)

	opts := &text.DrawOptions{}
	opts.GeoM.Translate(5, 560)
	opts.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, msg, assets.NormalFont, opts)
}

func (gs *GameScreen) drawScore(screen *ebiten.Image) {
	msg := fmt.Sprintf("O: %d | X: %d",
		gs.game.Players[1].Points,
		gs.game.Players[0].Points,
	)

	opts := &text.DrawOptions{}
	opts.PrimaryAlign = text.AlignCenter
	opts.GeoM.Translate(240, 570)
	opts.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, msg, assets.NormalFont, opts)
}

func (gs *GameScreen) drawEndMessage(screen *ebiten.Image) {
	var msg string
	if gs.game.Winner != nil {
		msg = fmt.Sprintf("%s wins!", gs.game.Winner.Symbol)
	} else {
		msg = "It's a draw!"
	}

	opts := &text.DrawOptions{}
	opts.PrimaryAlign = text.AlignCenter
	opts.SecondaryAlign = text.AlignCenter
	opts.GeoM.Translate(240, 300)
	opts.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, msg, assets.BigFont, opts)
}
