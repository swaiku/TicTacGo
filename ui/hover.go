package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type HoverMode int

const (
	HoverFade      HoverMode = iota // Fade alpha (0 -> 1)
	HoverColorLerp                  // Interpolate color normal -> hover
	HoverSolid                      // Direct hover color (no interpolation)
)

// Lerp performs a linear interpolation between two floats.
func lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}

// LerpColor interpolates between two colors.
func LerpColor(c1, c2 color.Color, t float64) color.Color {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()

	return color.RGBA{
		R: uint8(lerp(float64(r1>>8), float64(r2>>8), t)),
		G: uint8(lerp(float64(g1>>8), float64(g2>>8), t)),
		B: uint8(lerp(float64(b1>>8), float64(b2>>8), t)),
		A: uint8(lerp(float64(a1>>8), float64(a2>>8), t)),
	}
}

// ApplyHoverColor sets the ebiten.ColorScale based on hover amount and style.
func ApplyHoverColor(op *ebiten.DrawImageOptions, style WidgetStyle, hover float64) {
	switch style.HoverMode {

	case HoverFade:
		// Normal background + fade alpha depending on hover
		op.ColorScale.ScaleWithColor(style.BackgroundNormal)
		op.ColorScale.ScaleAlpha(float32(hover))

	case HoverColorLerp:
		// Smooth interpolation between normal and hover color
		col := LerpColor(style.BackgroundNormal, style.BackgroundHover, hover)
		op.ColorScale.ScaleWithColor(col)

	case HoverSolid:
		// Switch to hover color completely when hovered
		if hover >= 1 {
			op.ColorScale.ScaleWithColor(style.BackgroundHover)
		} else {
			op.ColorScale.ScaleWithColor(style.BackgroundNormal)
		}
	}
}
