package main

import (
	"fmt"
	"os"

	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/data/local"
	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	dashboard := ui.NewDashboard(&local.Client{})
	p := tea.NewProgram(dashboard)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Oh no, there's been an error: %v", err)
		os.Exit(1)
	}
}
