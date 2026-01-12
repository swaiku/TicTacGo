package utils

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// CreateGradientBackground generates an Ebiten image with a vertical gradient.
//
// The gradient interpolates from the top color to the bottom color,
// providing a smooth transition across the height of the image.
// This is commonly used for screen backgrounds.
//
// Parameters:
//   - width, height: dimensions of the image in pixels
//   - top: color at the top of the gradient
//   - bottom: color at the bottom of the gradient
//
// Returns an Ebiten image ready for rendering.
func CreateGradientBackground(width, height int, top, bottom color.Color) *ebiten.Image {
	img := ebiten.NewImage(width, height)
	bounds := img.Bounds()

	gradientHeight := float64(bounds.Dy())
	if gradientHeight == 0 {
		return img
	}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		// Calculate interpolation factor (0 at top, 1 at bottom)
		t := float64(y-bounds.Min.Y) / gradientHeight
		rowColor := LerpColor(top, bottom, t)

		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			img.Set(x, y, rowColor)
		}
	}

	return img
}
