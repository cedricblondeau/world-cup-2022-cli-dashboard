package ui

import (
	"sort"

	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/data"
	tea "github.com/charmbracelet/bubbletea"
)

type dataFetcher interface {
	GroupTables() ([]data.GroupTable, error)
	Matches() ([]data.Match, error)
	Name() string
}

type dataFetchMsg struct {
	groupTables []data.GroupTable
	matches     []data.Match
}

type dataFetchErrMsg struct{ err error }

func dataFetchCmd(fetcher dataFetcher) func() tea.Msg {
	return func() tea.Msg {
		groupTables, err := fetcher.GroupTables()
		if err != nil {
			return dataFetchErrMsg{err: err}
		}

		matches, err := fetcher.Matches()
		if err != nil {
			return dataFetchErrMsg{err: err}
		}

		sort.Slice(matches, func(i, j int) bool {
			return matches[i].Date.Before(matches[j].Date)
		})

		return dataFetchMsg{groupTables, matches}
	}
}
