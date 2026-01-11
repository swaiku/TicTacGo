package game

import (
	"GoTicTacToe/assets"
	"image/color"
)

// Player represents a participant in the game.
//
// A player can be either human-controlled or AI-controlled (IsAI flag).
// The Symbol and Color fields are used for rendering, while Points
// tracks the player's cumulative score across multiple rounds.
type Player struct {
	Symbol *assets.Symbol // Visual symbol rendered on the board
	Points int            // Score accumulated across rounds
	Color  color.Color    // Display color used in the UI
	Name   string         // Display name (optional)
	IsAI   bool           // True if controlled by an AI model
}

// NewPlayer creates and returns a new player instance.
//
// The player's score is initialized to zero. The Name and IsAI fields
// can be set by the caller after creation if needed.
func NewPlayer(sym *assets.Symbol, c color.Color) *Player {
	return &Player{
		Symbol: sym,
		Points: 0,
		Color:  c,
	}
}

// Opponent returns the first opponent of this player from the given list.
//
// This method assumes a two-player game and returns the first player
// in the list that is not the receiver. Returns nil if no opponent is found.
//
// For games with more than two players, use Opponents instead.
func (p *Player) Opponent(players []*Player) *Player {
	for _, candidate := range players {
		if candidate != p {
			return candidate
		}
	}
	return nil
}

// Opponents returns all players in the list that are not this player.
//
// This is useful for multi-player scenarios where there may be more than
// one opponent. The returned slice preserves the order from the input list.
func (p *Player) Opponents(players []*Player) []*Player {
	opponents := make([]*Player, 0, len(players)-1)
	for _, candidate := range players {
		if candidate != p {
			opponents = append(opponents, candidate)
		}
	}
	return opponents
}
