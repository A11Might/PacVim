package game

import (
	"math/rand"
	"time"

	"github.com/A11Might/PacVim/pkg/util"
)

var (
	GlobMaze Maze
	Player   *Avatar
	Ghost1   *Ghost

	MoveOrNot bool

	WonGame     int // 0-init, -1-lost, 1-win
	TotalPoints int
)

func InitGame() {
	GlobMaze.InitMaze(MapIndex[rand.Intn(10)])
	TotalPoints = GlobMaze.totalPoints
	Player = Born()
	//Ghost1 = SpawnGhost()
	Ghost1 = new(Ghost)
	MoveOrNot = false // 幽灵是否移动，用来控制幽灵移动速度的标志位
}

func GetPlayerPosition() (int, int) {
	return Player.X, Player.Y
}

func CanMove() bool {
	if MoveOrNot {
		MoveOrNot = false
		go func() {
			time.Sleep(util.GhostSpeed)
			MoveOrNot = true
		}()
		return true
	}
	return false
}
