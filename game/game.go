package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// State du jeu
type GameState int

const (
	PLAYING GameState = iota
	GAME_END
)

type Game struct {
	State    GameState
	Board    *Board
	Players  []*Player
	Current  *Player
	Winner   *Player
	BoardImg *ebiten.Image
}

// Constructeur principal
func NewGame() *Game {
	g := &Game{}
	g.ResetHard()
	return g
}

// Réinitialisation complète (nouvelle grille, nouveaux symboles)
func (g *Game) ResetHard() {
	g.Board = NewBoard(3, 3)
	g.Players = []*Player{
		NewPlayer(newSymbol(CIRCLE), color.White),
		NewPlayer(newSymbol(CROSS), color.White),
	}

	g.Current = g.Players[0]
	g.Winner = nil
	g.State = PLAYING
	g.BoardImg = g.Board.GenerateImage()
}

// Réinitialise la grille mais garde les scores
func (g *Game) Reset() {
	size := g.Board.Size
	win := g.Board.ToWin

	g.Board = NewBoard(size, win)
	g.BoardImg = g.Board.GenerateImage()

	g.Current = g.Players[0]
	g.Winner = nil
	g.State = PLAYING
}

func (g *Game) ResetPoints() {
	for _, p := range g.Players {
		p.Points = 0
	}
}

// Rotation des joueurs
func (g *Game) nextPlayer() {
	if len(g.Players) == 0 {
		return
	}
	for i, p := range g.Players {
		if p == g.Current {
			g.Current = g.Players[(i+1)%len(g.Players)]
			return
		}
	}
	g.Current = g.Players[0]
}

func (g *Game) GetCursorBoardPos(mx, my int) (int, int) {
	cell := 480 / g.Board.Size
	return mx / cell, my / cell
}

func (g *Game) PlayMove(x, y int) bool {
	ok := g.Board.Play(g.Current, x, y)
	if !ok {
		return false
	}

	return true
}

// Vérifie victoire
func (g *Game) CheckWin() bool {
	p := g.Board.CheckWin()
	if p != nil {
		g.Winner = p
		g.Winner.Points++
		g.State = GAME_END
		return true
	}
	return false
}

// Vérifie égalité
func (g *Game) CheckDraw() bool {
	if g.Board.CheckDraw() {
		g.Winner = nil
		g.State = GAME_END
		return true
	}
	return false
}
