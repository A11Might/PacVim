package game

import (
	"math/rand"

	"github.com/A11Might/PacVim/pkg/util"
)

type Avatar struct {
	X           int
	Y           int
	LetterUnder rune
	ColorUnder  int
	IsPlayer    bool
	Points      int64
	Portrait    rune
	Lives       int64
	Color       int
}

func Born() *Avatar {
	a, b := randPosition()
	for !IsValid(a, b) || CharAt(a, b) == ' ' || CharAt(a, b) == '~' {
		a, b = randPosition()
	}

	player := &Avatar{
		X:           a,
		Y:           b,
		LetterUnder: GlobMaze.Graph[a][b],
		IsPlayer:    true,
		Points:      0,
		Portrait:    util.PlayerPortrait,
		Lives:       3,
		Color:       util.Player,
	}

	// 在地图上出生
	GlobMaze.Graph[a][b] = player.Portrait
	GlobMaze.Paint[a][b] = player.Color

	return player
}

func (t *Avatar) MoveTo(a, b int) bool {
	if !IsValid(a, b) {
		return false
	}

	// 目标位置的字符
	curChar := CharAt(a, b)
	if t.IsPlayer { // 如果是玩家
		// 遇到河流
		if curChar == util.WaterPortrait {
			WonGame = -1
			return false
		}

		// 遇到幽灵
		if curChar == util.GhostPortrait {
			WonGame = -1
			return false
		}

		//// 正常行走
		//if curChar != ' ' {
		//	t.Points++
		//}

		// 移动，将走过的字符变成绿色
		WriteAtWithColor(t.X, t.Y, t.LetterUnder, util.Green)
		t.X = a
		t.Y = b
		t.LetterUnder = curChar
		WriteAtWithColor(t.X, t.Y, t.Portrait, util.Player)

		// 判断游戏是否结束
		if t.Points >= TotalPoints {
			WonGame = 1
		}
	} else { // 如果是幽灵
		playerX, playerY := GetPlayerPosition()
		// 幽灵抓住玩家
		if playerX == a && playerY == b {
			WonGame = -1
			return false
		}

		// 恢复脚下字符，并记录下一步字符
		WriteAtWithColor(t.X, t.Y, t.LetterUnder, t.ColorUnder)
		t.LetterUnder = CharAt(a, b)

		// 走过去
		t.X = a
		t.Y = b
		t.ColorUnder = GlobMaze.Paint[a][b]
		WriteAtWithColor(a, b, t.Portrait, util.Ghost)
	}

	return true
}

func (t *Avatar) MoveUp() bool {
	if !IsValid(t.X-1, t.Y) {
		return false
	}

	t.MoveTo(t.X-1, t.Y)
	return true
}

func (t *Avatar) MoveDown() bool {
	if !IsValid(t.X+1, t.Y) {
		return false
	}

	t.MoveTo(t.X+1, t.Y)
	return true
}

func (t *Avatar) MoveLeft() bool {
	if !IsValid(t.X, t.Y-1) {
		return false
	}

	t.MoveTo(t.X, t.Y-1)
	return true
}

func (t *Avatar) MoveRight() bool {
	if !IsValid(t.X, t.Y+1) {
		return false
	}

	t.MoveTo(t.X, t.Y+1)
	return true
}

func randPosition() (int, int) {
	return rand.Intn(GlobMaze.rows), rand.Intn(GlobMaze.cols)
}
