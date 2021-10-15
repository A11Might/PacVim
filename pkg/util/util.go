package util

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/muesli/termenv"
)

const (
	Faint = iota
	White
	Red
	Green
	BrightGreen
	Blue
	Yellow
	Magenta
	Cyan
	Wall
	Water
	Player
	Ghost = Red
	Char
	Nothing

	WallPortrait   = '#'
	WaterPortrait  = '~'
	PlayerPortrait = 'P'
	GhostPortrait  = 'G'

	Win        = 1
	Lost       = -1
	Gaming     = 0
	GhostSpeed = time.Second / 3
)

func init() {
	rand.Seed(time.Now().Unix())
}

func GetRandInt(n int) int {
	return rand.Intn(n)
}

func GetColor(color int) termenv.ANSIColor {
	switch color {
	case Red:
		return termenv.ANSIRed

	case Green:
		return termenv.ANSIGreen

	case BrightGreen:
		return termenv.ANSIBrightGreen

	case Blue:
		return termenv.ANSIBlue

	case Yellow:
		return termenv.ANSIYellow

	case Cyan:
		return termenv.ANSICyan

	case Magenta:
		return termenv.ANSIMagenta
	}

	return termenv.ANSIBlack
}

func MatrixToString(matrix [][]rune) string {
	var builder strings.Builder
	for _, line := range matrix {
		builder.WriteString(string(line) + "\n")
	}
	return builder.String()
}

func StringToMatrix(str string) [][]rune {
	lines := strings.Split(str, "\n")
	matrix := make([][]rune, len(lines))
	for i := range lines {
		matrix[i] = make([]rune, len(lines[i]))
		for j, chr := range lines[i] {
			matrix[i][j] = chr
		}
	}
	return matrix
}

func WhiteMatrix(matrix [][]rune) [][]int {
	rows, cols := len(matrix), len(matrix[0])
	ret := make([][]int, rows)
	for i := 0; i < rows; i++ {
		ret[i] = make([]int, cols)
	}

	// 初始化墙和水的颜色
	for i, line := range matrix {
		for j, chr := range line {
			switch chr {
			case WallPortrait:
				ret[i][j] = Wall

			case WaterPortrait:
				ret[i][j] = Water

			default:
				ret[i][j] = Faint
			}
		}
	}

	return ret
}

func PointMatrix(matrix [][]rune) ([][]int, int) {
	rows, cols := len(matrix), len(matrix[0])
	ret := make([][]int, rows)
	for i := 0; i < rows; i++ {
		ret[i] = make([]int, cols)
	}

	totalPoints := 0
	for i, line := range matrix {
		for j, chr := range line {
			a := string(chr)
			fmt.Sprintf(a)
			if chr != ' ' && chr != WallPortrait && chr != WaterPortrait {
				ret[i][j] = 1
				totalPoints++
			}
		}
	}

	return ret, totalPoints
}

func MinFloat(nums ...float64) float64 {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums[0]
}

func IsAlphanumeric(chr rune) bool {
	if unicode.IsLetter(chr) || unicode.IsDigit(chr) {
		return true
	}
	return false
}
