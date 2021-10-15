package main

import (
	"fmt"
	"os"

	"github.com/A11Might/PacVim/pkg/game"
	"github.com/A11Might/PacVim/pkg/model"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() model.Model {
	return model.Model{
		Player: game.Player,
		Ghost:  game.Ghost1,
	}
}
