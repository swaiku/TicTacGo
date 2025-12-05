package screens

import "github.com/hajimehoshi/ebiten/v2"

func GetWindowSize() (float64, float64) {
	w, h := ebiten.WindowSize()
	return float64(w), float64(h)
}
