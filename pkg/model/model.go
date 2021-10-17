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

		case "^":
			m.Player.ParseToBeginningFor6()

		case "$":
			m.Player.ParseToEnd()

		case "e":
			m.Player.ParseWordEnd()

		case "E":
			m.Player.ParseWordEndForE()

		case "b":
			m.Player.ParseWordBackward()

		case "B":
			m.Player.ParseWordBackwardForB()

		case "w":
			m.Player.ParseWordForward()

		case "W":
			m.Player.ParseWordForwardForW()

		case "g":
			m.Player.ParseToUpping()

		case "G":
			m.Player.ParseToDowning()

		case "ctrl+c":
			return m, tea.Quit
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
	for _, line := range game.GlobMaze.Graph {
		for _, cell := range line {
			style := termenv.String(string(cell.Char))
			switch cell.Color {
			case util.WallColor:
				builder.WriteString(style.Bold().String())
			case util.WaterColor:
				builder.WriteString(style.Foreground(util.GetColor(util.Blue)).String())
			case util.Faint:
				builder.WriteString(style.Faint().String())
			case util.PlayerColor:
				builder.WriteString(style.Background(util.GetColor(util.BrightGreen)).String())
			default:
				builder.WriteString(style.Foreground(util.GetColor(cell.Color)).String())
			}
		}
		builder.WriteString("\n")
	}
	builder.WriteString(fmt.Sprintf("\nPoints: %d/%d", m.Player.Points, game.TotalPoints))

	switch game.WonGame {
	case util.Lost:
		return "You Lost\nPress Ctrl + C to quit game"

	case util.Win:
		return "You Win\nPress Ctrl + C to quit game"
	}

	return builder.String()
}
