package main

import (
	"fmt"
	"os"

	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/data/footballdata"
	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/data/worldcupjson"
	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	token, ok := os.LookupEnv("FOOTBALLDATA_API_TOKEN")
	var dashboard tea.Model
	if ok {
		dashboard = ui.NewDashboard(footballdata.NewClient(token))
	} else {
		dashboard = ui.NewDashboard(worldcupjson.NewClient())
	}

	p := tea.NewProgram(dashboard)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Oh no, there's been an error: %v", err)
		os.Exit(1)
	}
}
