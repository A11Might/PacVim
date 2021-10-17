package game

import (
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

func init() {
	GlobMaze.InitMaze(Map0)
	TotalPoints = GlobMaze.totalPoints
	GlobMaze.rows = len(GlobMaze.Graph)
	GlobMaze.cols = len(GlobMaze.Graph[0])
	Player = Born()
	Ghost1 = SpawnGhost()
	MoveOrNot = true
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
