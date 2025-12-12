package screens

import (
	"GoTicTacToe/assets"
	"GoTicTacToe/ui"
	uiutils "GoTicTacToe/ui/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type StartScreen struct {
	host       ScreenHost
	buttons    []*ui.Button
	background *ebiten.Image
}

const (
	buttonWidth  = float64(200)
	buttonHeight = float64(64)
	buttonRadius = float64(10)
	buttonSpacing = float64(20)
	buttonYOffset = float64(140)
)

func NewStartScreen(h ScreenHost) *StartScreen {
	s := &StartScreen{host: h}

	s.buttons = []*ui.Button{
		ui.NewButton("Local", -buttonWidth - buttonSpacing, buttonYOffset, uiutils.AnchorCenter,
			buttonWidth, buttonHeight, buttonRadius, uiutils.NormalWidgetStyle,
			func() {
				s.host.SetScreen(NewGameScreen(s.host, GameConfig{
					Mode: LocalVsLocal,
				}))
			},
		),

		ui.NewButton("Play vs AI", 0, buttonYOffset, uiutils.AnchorCenter,
			buttonWidth, buttonHeight, buttonRadius,
			uiutils.NormalWidgetStyle,
			func() {
				h.SetScreen(NewAIScreen(h, GameConfig{
					Mode: LocalVsAI,
				}))
			},
		),

		ui.NewButton(
			"Multiplayer \n (not implemented)",
			buttonWidth + buttonSpacing, buttonYOffset,
			uiutils.AnchorCenter,
			buttonWidth, buttonHeight, buttonRadius,
			uiutils.DangerWidgetStyle,
			func() {},
		),
	}
	return s
}

func (s *StartScreen) Update() error {
	for _, btn := range s.buttons {
		btn.Update()
	}
	return nil
}

func (s *StartScreen) Draw(screen *ebiten.Image) {
	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()
	if s.background == nil {
		topColor := color.RGBA{R: 0x1C, G: 0x25, B: 0x41, A: 0xFF}
		bottomColor := color.RGBA{R: 0x0B, G: 0x13, B: 0x2B, A: 0xFF}
		s.background = uiutils.CreateGradientBackground(w, h, topColor, bottomColor)
	}
	screen.DrawImage(s.background, nil)

	logo := assets.Logo
	origLogoW := float64(logo.Bounds().Dx())
	origLogoH := float64(logo.Bounds().Dy())

	op := &ebiten.DrawImageOptions{}

	op.Filter = ebiten.FilterLinear

	targetWidth := float64(w) * 0.5

	scaleFactor := targetWidth / origLogoW

	op.GeoM.Scale(scaleFactor, scaleFactor)

	scaledLogoW := origLogoW * scaleFactor
	scaledLogoH := origLogoH * scaleFactor

	posX := (float64(w) - scaledLogoW) / 2
	posY := (float64(h) - scaledLogoH) / 2 * 0.4

	op.GeoM.Translate(posX, posY)

	screen.DrawImage(logo, op)

	// Buttons
	for _, btn := range s.buttons {
		btn.Draw(screen)
	}
}
