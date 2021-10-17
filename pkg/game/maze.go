package game

import (
	"strings"

	"github.com/A11Might/PacVim/pkg/util"
)

type Maze struct {
	Graph       [][]*Cell
	rows        int
	cols        int
	totalPoints int
}

type Cell struct {
	Char  rune
	Color int
	Point int
}

func (m *Maze) InitMaze(str string) {
	lines := strings.Split(str, "\n")
	m.Graph = make([][]*Cell, len(lines))
	totalPoints := 0
	for i := range lines {
		m.Graph[i] = make([]*Cell, len(lines[i]))
		for j, chr := range lines[i] {
			m.Graph[i][j] = new(Cell)
			m.Graph[i][j].Char = chr
			// 初始化墙和水的颜色
			switch chr {
			case util.WallPortrait:
				m.Graph[i][j].Color = util.WallColor

			case util.WaterPortrait:
				m.Graph[i][j].Color = util.WaterColor

			case ' ':

			default:
				m.Graph[i][j].Color = util.Faint
				m.Graph[i][j].Point = 1
				totalPoints++
			}
		}
	}

	m.totalPoints = totalPoints
}

func IsValid(a, b int) bool {
	if GlobMaze.Graph[a][b].Char == '#' {
		return false
	}
	return true
}

func CharAt(a, b int) rune {
	return GlobMaze.Graph[a][b].Char
}

func WriteAtWithColor(a, b int, curChar rune, color int) {
	GlobMaze.Graph[a][b].Char = curChar
	GlobMaze.Graph[a][b].Color = color
}

func WriteAt(a, b int, curChar rune) {
	GlobMaze.Graph[a][b].Char = curChar
}
