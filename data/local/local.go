package local

import (
	"encoding/json"

	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/data"
)

func SortedLocalMatches() ([]data.Match, error) {
	var matches []data.Match
	err := json.Unmarshal([]byte(matchesJSON), &matches)
	if err != nil {
		return nil, err
	}

	return matches, nil
}
