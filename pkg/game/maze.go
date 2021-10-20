package game

import (
	"strings"

	"github.com/A11Might/PacVim/pkg/util"
	"github.com/beefsack/go-astar"
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
	X     int
	Y     int
}

func (c *Cell) PathNeighbors() []astar.Pather {
	neighbors := make([]astar.Pather, 0)
	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		a, b := c.X+offset[0], c.Y+offset[1]
		if a >= 0 && a < Rows && b >= 0 && b < Cols {
			neighbors = append(neighbors, GlobMaze.Graph[a][b])
		}
	}
	return neighbors
}

func (c *Cell) PathNeighborCost(to astar.Pather) float64 {
	if to.(*Cell).Char == '#' {
		return 1000
	}
	return 1
}

// PathEstimatedCost 使用曼哈顿距离作为启发式算法
func (c *Cell) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*Cell)
	absX := toT.X - c.X
	if absX < 0 {
		absX = -absX
	}
	absY := toT.Y - c.Y
	if absY < 0 {
		absY = -absY
	}
	r := float64(absX + absY)

	return r
}

func (m *Maze) InitMaze(str string) {
	lines := strings.Split(str, "\n")
	m.Graph = make([][]*Cell, len(lines))
	totalPoints := 0
	for i := range lines {
		m.Graph[i] = make([]*Cell, len(lines[i]))
		for j, chr := range lines[i] {
			m.Graph[i][j] = new(Cell)
			m.Graph[i][j].X = i
			m.Graph[i][j].Y = j
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
				// 初始化字符分数
				m.Graph[i][j].Point = 1
				totalPoints++
			}
		}
	}

	m.rows = len(m.Graph)
	m.cols = len(m.Graph[0])
	Rows = m.rows
	Cols = m.cols
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
