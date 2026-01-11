// Package assets manages game resources including fonts, images, and symbols.
//
// All assets are loaded once at startup via init() functions and stored
// in package-level variables for efficient access throughout the application.
package assets

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Asset file paths.
const (
	// logoImagePath is the relative path to the application logo image.
	logoImagePath = "assets/static/gonnectmax_logo.png"
)

// Logo is the application logo displayed on the start screen.
//
// The image is loaded at startup and cached for efficient rendering.
var Logo *ebiten.Image

// init loads the static image assets at startup.
//
// The logo image is loaded at startup and stored in a shared variable
// to avoid reloading it multiple times during the game.
func init() {
	var err error
	Logo, _, err = ebitenutil.NewImageFromFile(logoImagePath)
	if err != nil {
		log.Fatalf("failed to load logo image from %s: %v", logoImagePath, err)
	}
}
