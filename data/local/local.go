package local

import (
	"encoding/json"

	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/data"
)

func LocalMatchesByID() (map[int]data.Match, error) {
	var matches []data.Match
	err := json.Unmarshal([]byte(matchesJSON), &matches)
	if err != nil {
		return nil, err
	}

	matchesByID := make(map[int]data.Match)
	for _, m := range matches {
		matchesByID[m.ID] = m
	}

	return matchesByID, nil
}
