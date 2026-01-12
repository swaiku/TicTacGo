package utils

import "image/color"

// lerp performs linear interpolation between two values.
//
// Returns a + (b - a) * t, where t is typically in the range [0, 1].
// When t=0, returns a. When t=1, returns b.
func lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}

// LerpColor performs linear interpolation between two colors.
//
// The interpolation is performed independently on each RGBA channel.
// Parameter t should be in the range [0, 1]:
//   - t=0 returns c1
//   - t=1 returns c2
//   - t=0.5 returns the midpoint color
//
// This is useful for smooth color transitions in hover animations.
func LerpColor(c1, c2 color.Color, t float64) color.Color {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()

	// RGBA() returns 16-bit values, shift to 8-bit for RGBA struct
	return color.RGBA{
		R: uint8(lerp(float64(r1>>8), float64(r2>>8), t)),
		G: uint8(lerp(float64(g1>>8), float64(g2>>8), t)),
		B: uint8(lerp(float64(b1>>8), float64(b2>>8), t)),
		A: uint8(lerp(float64(a1>>8), float64(a2>>8), t)),
	}
}
