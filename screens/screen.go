package screens

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Screen is an interface that all screens must implement
// It defines the methods that a screen must implement to be used in the game.
type Screen interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type ScreenHost interface {
	SetScreen(Screen)
}

type screenHost struct {
	width, height int
	current       Screen
}

func NewScreenHost(width, height int) *screenHost {
	return &screenHost{
		width:  width,
		height: height,
	}
}

func (h *screenHost) SetScreen(s Screen) {
	h.current = s
}

func (h *screenHost) Update() error {
	if h.current != nil {
		return h.current.Update()
	}
	return nil
}

func (h *screenHost) Draw(screen *ebiten.Image) {
	if h.current != nil {
		h.current.Draw(screen)
	}
}

func (h *screenHost) Layout(_, _ int) (int, int) {
	return h.width, h.height
}
