package utils

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

// CreateRoundedRect generates an Ebiten image of a rounded rectangle.
//
// The rectangle is rendered with anti-aliasing using the gg library.
// This is typically used for button and container backgrounds.
//
// Parameters:
//   - width, height: dimensions of the image in pixels
//   - radius: corner radius in pixels (0 for sharp corners)
//   - fillColor: the solid fill color for the rectangle
//
// Returns an Ebiten image ready for rendering.
func CreateRoundedRect(width, height int, radius float64, fillColor color.Color) *ebiten.Image {
	dc := gg.NewContext(width, height)
	dc.SetColor(fillColor)
	dc.DrawRoundedRectangle(0, 0, float64(width), float64(height), radius)
	dc.Fill()

	return ebiten.NewImageFromImage(dc.Image())
}
