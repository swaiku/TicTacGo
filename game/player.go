package game

import (
	"GoTicTacToe/assets"
	"image/color"
)

// Player represents a participant in the game with their visual identity and score.
type Player struct {
	Symbol *assets.Symbol // Visual symbol displayed on the board (X, O, etc.)
	Points int            // Accumulated score across multiple rounds
	Color  color.Color    // Display color for the player's symbol and UI elements
	Name   string         // Display name shown in the UI
	IsAI   bool           // True if this player is controlled by an AI model
}

// NewPlayer creates a new Player with the given symbol and color.
// Points start at zero, and the player defaults to human-controlled.
func NewPlayer(symbol *assets.Symbol, playerColor color.Color) *Player {
	return &Player{
		Symbol: symbol,
		Points: 0,
		Color:  playerColor,
	}
}

// Opponent returns the first player in the list that is not this player.
// Returns nil if no opponent is found (e.g., single-player scenario).
// Note: For games with more than 2 players, this returns only the first opponent.
func (p *Player) Opponent(players []*Player) *Player {
	for _, candidate := range players {
		if candidate != p {
			return candidate
		}
	}
	return nil
}

// Opponents returns all players in the list that are not this player.
// Useful for multi-player scenarios where there may be more than one opponent.
func (p *Player) Opponents(players []*Player) []*Player {
	var opponents []*Player
	for _, candidate := range players {
		if candidate != p {
			opponents = append(opponents, candidate)
		}
	}
	return opponents
}
