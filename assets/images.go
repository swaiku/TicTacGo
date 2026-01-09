package assets

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	Logo *ebiten.Image
)

func init() {

	// Ne pas charger les assets graphiques pendant les tests ou la CI
	if os.Getenv("TESTING") == "1" {
		return
	}

	logoImg, _, err := ebitenutil.NewImageFromFile("assets/static/gonnectmax_logo.png")
	if err != nil {
		log.Fatal(err)
	}
	Logo = logoImg
}
