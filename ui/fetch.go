package ui

import (
	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/data"
	tea "github.com/charmbracelet/bubbletea"
)

type dataFetcher interface {
	GroupTables() ([]data.GroupTable, error)
	Name() string
}

type dataFetchMsg struct {
	groupTables []data.GroupTable
}

type dataFetchErrMsg struct{ err error }

func dataFetchCmd(fetcher dataFetcher) func() tea.Msg {
	return func() tea.Msg {
		groupTables, err := fetcher.GroupTables()
		if err != nil {
			return dataFetchErrMsg{err: err}
		}

		return dataFetchMsg{groupTables}
	}
}
