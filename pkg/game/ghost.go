package game

import (
	"math"

	"github.com/A11Might/PacVim/pkg/util"
)

type Ghost struct {
	Avatar
}

func SpawnGhost() *Ghost {
	a, b := randPosition()
	for !IsValid(a, b) || CharAt(a, b) == ' ' || CharAt(a, b) == util.PlayerPortrait {
		a, b = randPosition()
	}

	ghost := &Ghost{
		Avatar: Avatar{
			X:           a,
			Y:           b,
			LetterUnder: GlobMaze.Graph[a][b].Char,
			ColorUnder:  GlobMaze.Graph[a][b].Color,
			IsPlayer:    false,
			Portrait:    util.GhostPortrait,
			Color:       util.Ghost},
	}

	// 在地图上出生
	GlobMaze.Graph[a][b].Char = ghost.Portrait
	GlobMaze.Graph[a][b].Color = ghost.Color

	return ghost
}

func (g *Ghost) Think() {
	//time.Sleep(1 * time.Second)
	up := eval(g.X-1, g.Y)
	down := eval(g.X+1, g.Y)
	left := eval(g.X, g.Y-1)
	right := eval(g.X, g.Y+1)
	minVal := util.MinFloat(up, down, left, right)
	switch minVal {
	case up:
		g.MoveTo(g.X-1, g.Y)

	case down:
		g.MoveTo(g.X+1, g.Y)

	case left:
		g.MoveTo(g.X, g.Y-1)

	case right:
		g.MoveTo(g.X, g.Y+1)
	}
}

func eval(a, b int) float64 {
	if !IsValid(a, b) {
		return 1000
	}
	playerX, playerY := GetPlayerPosition()
	return math.Sqrt(math.Pow(float64(playerX-a), 2) + math.Pow(float64(playerY-b), 2))
}
