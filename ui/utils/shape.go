package utils

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

// CreateRoundedRect creates an Ebiten image of a rounded rectangle with anti-aliasing.
// Returns an Ebiten image of the rounded rectangle.
func CreateRoundedRect(width, height int, radius float64, color color.Color) *ebiten.Image {
	dc := gg.NewContext(width, height)

	dc.SetColor(color)

	dc.DrawRoundedRectangle(0, 0, float64(width), float64(height), radius)
	dc.Fill()

	return ebiten.NewImageFromImage(dc.Image())
}
