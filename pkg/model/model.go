package model

import (
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
		for j, chr := range line {
			if color := game.GlobMaze.Paint[i][j]; color == util.Wall {
				builder.WriteString(termenv.String(string(chr)).Bold().String())
			} else if color == util.Water {
				builder.WriteString(termenv.String(string(chr)).Foreground(util.GetColor(util.Blue)).String())
			} else if color == util.Faint {
				builder.WriteString(termenv.String(string(chr)).Faint().String())
			} else if color == util.White {
				builder.WriteString(string(chr))
			} else if color == util.Player {
				builder.WriteString(termenv.String(string(chr)).Background(util.GetColor(util.BrightGreen)).String())
			} else {
				builder.WriteString(termenv.String(string(chr)).Foreground(util.GetColor(game.GlobMaze.Paint[i][j])).String())
			}
		}
		builder.WriteString("\n")
	}

	if game.Lost() {
		return "You Lost"
	}
	return builder.String()
}
