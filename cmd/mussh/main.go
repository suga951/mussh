package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"mussh/internal/ui"
)

func main() {
	fmt.Print("\033[H\033[2J")

	p := tea.NewProgram(ui.New())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running termify: ", err)
		os.Exit(1)
	}
}
