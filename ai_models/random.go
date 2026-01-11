package ai_models

import (
	"GoTicTacToe/game"
	"math/rand"
)

// invalidMoveCoord is returned when no valid move is available.
const invalidMoveCoord = -1

// RandomAI implements the AIModel interface with a random move selection strategy.
//
// This AI simply picks a random empty cell from the available moves,
// providing an "easy" difficulty level suitable for casual play or testing.
type RandomAI struct{}

// NextMove selects a random available move on the board.
//
// Returns (-1, -1) if the board has no empty cells.
func (RandomAI) NextMove(board *game.Board, me *game.Player, players []*game.Player) (int, int) {
	moves := board.AvailableMoves()
	if len(moves) == 0 {
		return invalidMoveCoord, invalidMoveCoord
	}

	selected := moves[rand.Intn(len(moves))]
	return selected.X, selected.Y
}
