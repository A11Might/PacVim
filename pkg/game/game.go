package game

import (
	"sync"
	"time"

	"github.com/A11Might/PacVim/pkg/util"
)

var (
	GlobMaze Maze
	Player   *Avatar
	Ghost1   *Ghost

	MoveOrNot bool
	mutex     sync.RWMutex

	WonGame     int // 0-init, -1-lost, 1-win
	TotalPoints int64
)

func init() {
	GlobMaze.Graph = util.StringToMatrix(Map0)
	GlobMaze.Paint = util.WhiteMatrix(GlobMaze.Graph)
	GlobMaze.rows = len(GlobMaze.Graph)
	GlobMaze.cols = len(GlobMaze.Graph[0])
	Player = Born()
	Ghost1 = SpawnGhost()
	MoveOrNot = true

}

func GetPlayerPosition() (int, int) {
	return Player.X, Player.Y
}

func Lost() bool {
	if WonGame == -1 {
		return true
	}
	return false
}

func CanMove() bool {
	if MoveOrNot {
		MoveOrNot = false
		go func() {
			time.Sleep(1 * time.Second)
			MoveOrNot = true
		}()
		return true
	}
	return false
}
