package game

type Maze struct {
	Graph [][]rune
	Paint [][]int
	Point [][]int
	rows  int
	cols  int
	score int
}

func IsValid(a, b int) bool {
	if string(GlobMaze.Graph[a][b]) == "#" {
		return false
	}
	return true
}

func CharAt(a, b int) rune {
	return GlobMaze.Graph[a][b]
}

func WriteAtWithColor(a, b int, curChar rune, color int) {
	GlobMaze.Graph[a][b] = curChar
	GlobMaze.Paint[a][b] = color
}

func WriteAt(a, b int, curChar rune) {
	GlobMaze.Graph[a][b] = curChar
}
