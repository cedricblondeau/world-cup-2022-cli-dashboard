package ui

import (
	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/data"
	tea "github.com/charmbracelet/bubbletea"
)

type dataFetcher interface {
	GroupTables() ([]data.GroupTable, error)
	SortedMatches() ([]data.Match, error)
	Name() string
}

type dataFetchMsg struct {
	groupTablesByLetter map[string]data.GroupTable
	sortedMatches       []data.Match
}

type dataFetchErrMsg struct{ err error }

func dataFetchCmd(fetcher dataFetcher) func() tea.Msg {
	return func() tea.Msg {
		groupTables, err := fetcher.GroupTables()
		if err != nil {
			return dataFetchErrMsg{err: err}
		}
		groupTablesByLetter := make(map[string]data.GroupTable, len(groupTables))
		for _, g := range groupTables {
			groupTablesByLetter[g.Letter] = g
		}

		sortedMatches, err := fetcher.SortedMatches()
		if err != nil {
			return dataFetchErrMsg{err: err}
		}

		return dataFetchMsg{groupTablesByLetter, sortedMatches}
	}
}
