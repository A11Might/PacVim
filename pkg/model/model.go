package model

import (
	"fmt"
	"strings"

	"github.com/A11Might/PacVim/pkg/game"
	"github.com/A11Might/PacVim/pkg/util"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
)

type Model struct {
	Player *game.Avatar
	Ghost  *game.Ghost
}

type move struct {
}

func autoMove() tea.Cmd {
	return func() tea.Msg {
		return move{}
	}
}

func (m Model) Init() tea.Cmd {
	return autoMove()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "k":
			m.Player.MoveUp()

		case "j":
			m.Player.MoveDown()

		case "h":
			m.Player.MoveLeft()

		case "l":
			m.Player.MoveRight()

		case "0":
			m.Player.ParseToBeginning()

		case "$":
			m.Player.ParseToEnd()

		case "e":
			m.Player.ParseWordEnd()

		case "b":
			m.Player.ParseWordBackward()

		case "w":
			m.Player.ParseWordForward()

		case "g":
			m.Player.ParseToUpping()

		case "G":
			m.Player.ParseToDowning()
		}
	case move:
		if game.CanMove() {
			m.Ghost.Think()
		}
		return m, autoMove()
	}

	return m, nil
}

func (m Model) View() string {
	var builder strings.Builder
	for i, line := range game.GlobMaze.Graph {
		for j, cell := range line {
			// 为啥
			//switch game.GlobMaze.Paint[i][i] {
			//case util.WallColor:
			//	builder.WriteString(termenv.String(string(cell)).Bold().String())
			//case util.WaterColor:
			//	builder.WriteString(termenv.String(string(cell)).Foreground(util.GetColor(util.Blue)).String())
			//case util.Faint:
			//	builder.WriteString(termenv.String(string(cell)).Faint().String())
			//case util.PlayerColor:
			//	builder.WriteString(termenv.String(string(cell)).Background(util.GetColor(util.BrightGreen)).String())
			//default:
			//	builder.WriteString(termenv.String(string(cell)).Foreground(util.GetColor(game.GlobMaze.Paint[i][j])).String())
			//}
			if color := game.GlobMaze.Graph[i][j].Color; color == util.WallColor {
				builder.WriteString(termenv.String(string(cell.Char)).Bold().String())
			} else if color == util.WaterColor {
				builder.WriteString(termenv.String(string(cell.Char)).Foreground(util.GetColor(util.Blue)).String())
			} else if color == util.Faint {
				builder.WriteString(termenv.String(string(cell.Char)).Faint().String())
			} else if color == util.PlayerColor {
				builder.WriteString(termenv.String(string(cell.Char)).Background(util.GetColor(util.BrightGreen)).String())
			} else {
				builder.WriteString(termenv.String(string(cell.Char)).Foreground(util.GetColor(cell.Color)).String())
			}
		}
		builder.WriteString("\n")
	}

	switch game.WonGame {
	case util.Lost:
		return "You Lost"

	case util.Win:
		return "You Win"
	}

	return builder.String() + fmt.Sprintf("\nPoints: %d", m.Player.Points)
}
