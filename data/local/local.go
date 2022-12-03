package local

import (
	"encoding/json"

	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/data"
)

func SortedMatches() ([]data.Match, error) {
	var matches []data.Match
	err := json.Unmarshal([]byte(matchesJSON), &matches)
	if err != nil {
		return nil, err
	}

	return matches, nil
}

func GroupTables() ([]data.GroupTable, error) {
	var groups []data.GroupTable
	err := json.Unmarshal([]byte(groupsJSON), &groups)
	if err != nil {
		return nil, err
	}

	return groups, nil
}
