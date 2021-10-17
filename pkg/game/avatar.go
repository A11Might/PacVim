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
	Points      int
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
		LetterUnder: GlobMaze.Graph[a][b].Char,
		IsPlayer:    true,
		Points:      0,
		Portrait:    util.PlayerPortrait,
		Lives:       3,
		Color:       util.PlayerColor,
	}

	// 在地图上出生
	GlobMaze.Graph[a][b].Char = player.Portrait
	GlobMaze.Graph[a][b].Color = player.Color

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

		// 正常行走
		if GlobMaze.Graph[a][b].Point == 1 {
			t.Points++
			GlobMaze.Graph[a][b].Point--
		}

		// 移动，将走过的字符变成绿色
		WriteAtWithColor(t.X, t.Y, t.LetterUnder, util.Green)
		t.X = a
		t.Y = b
		t.LetterUnder = curChar
		WriteAtWithColor(t.X, t.Y, t.Portrait, util.PlayerColor)

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
		t.ColorUnder = GlobMaze.Graph[a][b].Color
		WriteAtWithColor(a, b, t.Portrait, util.Ghost)
	}

	return true
}

func (t *Avatar) MoveUp() bool {
	if !IsValid(t.X-1, t.Y) {
		return false
	}

	return t.MoveTo(t.X-1, t.Y)
}

func (t *Avatar) MoveDown() bool {
	if !IsValid(t.X+1, t.Y) {
		return false
	}

	return t.MoveTo(t.X+1, t.Y)
}

func (t *Avatar) MoveLeft() bool {
	if !IsValid(t.X, t.Y-1) {
		return false
	}

	return t.MoveTo(t.X, t.Y-1)
}

func (t *Avatar) MoveRight() bool {
	if !IsValid(t.X, t.Y+1) {
		return false
	}

	return t.MoveTo(t.X, t.Y+1)
}

// ParseWordEnd for e
func (t *Avatar) ParseWordEnd() bool {
	// 向右走
	if CharAt(t.X, t.Y+1) == ' ' {
		t.MoveRight()
	}
	nextChar := CharAt(t.X, t.Y+1)

	for nextChar == ' ' {
		if !t.MoveRight() {
			return false
		}
		nextChar = CharAt(t.X, t.Y+1)
	}

	for true {
		if t.endForSymbol() {
			t.MoveRight()
			return true
		} else {
			if !t.MoveRight() {
				return false
			}
			if t.endForSymbol() {
				break
			}

		}
	}
	return true
}

// endForSymbol for case: 从左往右 a. or ,.
func (t *Avatar) endForSymbol() bool {
	if !util.IsAlphanumeric(CharAt(t.X, t.Y+1)) {
		return true
	}
	return false
}

// ParseWordEndForE for E
func (t *Avatar) ParseWordEndForE() bool {
	// 走到下一个单词的开头
	for t.LetterUnder != ' ' {
		if !t.MoveRight() {
			return false
		}
	}
	for t.LetterUnder == ' ' {
		if !t.MoveRight() {
			return false
		}
	}

	// 走到单词的末尾
	for CharAt(t.X, t.Y+1) != ' ' {
		if !t.MoveRight() {
			return false
		}
	}

	return true
}

// ParseWordBackward for b
func (t *Avatar) ParseWordBackward() bool {
	// 向左走
	if CharAt(t.X, t.Y-1) == ' ' {
		t.MoveLeft()
	}
	nextChar := CharAt(t.X, t.Y-1)

	for nextChar == ' ' {
		if !t.MoveLeft() {
			return false
		}
		nextChar = CharAt(t.X, t.Y-1)
	}

	for true {
		if t.beginForSymbol() {
			t.MoveLeft()
			return true
		} else {
			if !t.MoveLeft() {
				return false
			}
			if t.beginForSymbol() {
				break
			}
		}
	}
	return true
}

// beginForSymbol for case: 从右往左 .a or .,
func (t *Avatar) beginForSymbol() bool {
	if !util.IsAlphanumeric(CharAt(t.X, t.Y-1)) {
		return true
	}
	return false
}

// ParseWordBackwardForB for B
func (t *Avatar) ParseWordBackwardForB() bool {
	// 走到上一个单词的末尾
	for t.LetterUnder != ' ' {
		if !t.MoveLeft() {
			return false
		}
	}
	for t.LetterUnder == ' ' {
		if !t.MoveLeft() {
			return false
		}
	}

	// 走到单词的开头
	for CharAt(t.X, t.Y-1) != ' ' {
		if !t.MoveLeft() {
			return false
		}
	}

	return true
}

// ParseWordForward for w
func (t *Avatar) ParseWordForward() bool {
	// 向右走
	for true {
		if CharAt(t.X, t.Y+1) == ' ' {
			if !t.MoveRight() {
				return false
			}
		} else if t.endCondition() {
			t.MoveRight()
			return true
		} else {
			if !t.MoveRight() {
				return false
			}
		}
	}
	return true
}

func (t *Avatar) endCondition() bool {
	if (t.LetterUnder == ' ' && CharAt(t.X, t.Y+1) != ' ') ||
		(util.IsAlphanumeric(t.LetterUnder) && (!util.IsAlphanumeric(CharAt(t.X, t.Y+1)) && CharAt(t.X, t.Y+1) != ' ')) ||
		(!util.IsAlphanumeric(t.LetterUnder) && util.IsAlphanumeric(CharAt(t.X, t.Y+1))) ||
		(!util.IsAlphanumeric(t.LetterUnder) && !util.IsAlphanumeric(CharAt(t.X, t.Y+1))) {
		return true
	}
	return false
}

// ParseWordForwardForW for W
func (t *Avatar) ParseWordForwardForW() bool {
	// 走到下一个单词的开头之前
	for t.LetterUnder != ' ' {
		if !t.MoveRight() {
			return false
		}
	}
	for t.LetterUnder == ' ' {
		if !t.MoveRight() {
			return false
		}
	}

	return true
}

// ParseToEnd for $
func (t *Avatar) ParseToEnd() bool {
	// 注地图不是规则矩形
	// 从后往前找到第一堵墙
	a, b := t.X, GlobMaze.cols-1
	for IsValid(a, b) {
		b--
	}

	// 再往前，进图
	// for case: # ####    #### ### #####
	//								  <--
	for !IsValid(a, b) {
		b--
	}

	t.MoveTo(a, b)
	return true
}

// ParseToBeginning for 0
func (t *Avatar) ParseToBeginning() bool {
	// 同理
	a, b := t.X, 0
	for IsValid(a, b) {
		b++
	}
	b++

	// for case: #####    ###  ####### #
	//  		 -->
	for !IsValid(a, b) {
		b++
	}

	t.MoveTo(a, b)
	return true
}

// ParseToBeginningFor6 for ^
func (t *Avatar) ParseToBeginningFor6() bool {
	t.ParseToBeginning()
	// 踩到空，则移到第一个单词的词首
	if t.LetterUnder == ' ' {
		t.ParseWordForward()
	}
	return true
}

// ParseToUpping for gg
func (t *Avatar) ParseToUpping() bool {
	// 同理
	a, b := 0, t.Y
	for IsValid(a, b) {
		a++
	}

	// for case: #  |
	//			 #  |
	//				v
	//			 #
	for !IsValid(a, b) {
		a++
	}

	t.MoveTo(a, b)
	return true
}

// ParseToDowning for G
func (t *Avatar) ParseToDowning() bool {
	// 同理
	a, b := GlobMaze.rows-1, t.Y
	for IsValid(a, b) {
		a--
	}

	// for case: #
	//			    ^
	//		     #	|
	//			 #  |
	for !IsValid(a, b) {
		a--
	}

	t.MoveTo(a, b)
	return true
}

func randPosition() (int, int) {
	return rand.Intn(GlobMaze.rows), rand.Intn(GlobMaze.cols)
}
