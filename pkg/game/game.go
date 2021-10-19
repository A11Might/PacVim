package game

import (
	"math/rand"
)

var (
	GlobMaze Maze
	Player   *Avatar
	Ghost1   *Ghost

	WonGame     int // 0-init, -1-lost, 1-win
	TotalPoints int
)

func InitGame() {
	GlobMaze.InitMaze(MapIndex[rand.Intn(10)])
	TotalPoints = GlobMaze.totalPoints
	Player = Born()
	Ghost1 = SpawnGhost()
	//Ghost1 = new(Ghost) // 不造幽灵
}

func GetPlayerPosition() (int, int) {
	return Player.X, Player.Y
}
