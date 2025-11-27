package game

import "image/color"

type Player struct {
	Symbol *Symbol
	Points int
	Color  color.Color
}

func NewPlayer(sym *Symbol, color color.Color) *Player {
	return &Player{Symbol: sym, Points: 0, Color: color}
}
