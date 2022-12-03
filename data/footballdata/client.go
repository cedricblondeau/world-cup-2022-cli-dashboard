package footballdata

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/data"
	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/data/local"
)

type mockableHttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	httpClient mockableHttpClient
	token      string
}

func NewClient(token string) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
		token: token,
	}
}

func NewMockClient() *Client {
	return &Client{
		httpClient: &mockHttpClient{},
	}
}

func (c *Client) Name() string {
	return "football-data.org"
}

func (c *Client) GroupTables() ([]data.GroupTable, error) {
	b, err := httpGetBytes(c.httpClient, "https://api.football-data.org/v4/competitions/WC/standings", c.token)
	if err != nil {
		return nil, err
	}

	var p parsedStandings
	if err := json.Unmarshal(b, &p); err != nil {
		return nil, err
	}

	groupTables := make([]data.GroupTable, len(p.Groups))
	for i, parsedGroup := range p.Groups {
		table := make([]data.GroupTableTeam, len(parsedGroup.Table))
		for j, team := range parsedGroup.Table {
			table[j] = data.GroupTableTeam{
				Code:              team.Team.TLA,
				Points:            team.Points,
				MatchesPlayed:     team.PlayedGames,
				Wins:              team.Won,
				Draws:             team.Draw,
				Losses:            team.Lost,
				GoalsFor:          team.GoalsFor,
				GoalsAgainst:      team.GoalsAgainst,
				GoalsDifferential: team.GoalDifference,
			}
		}

		groupTables[i] = data.GroupTable{
			Letter: strings.TrimPrefix(parsedGroup.Group, "GROUP_"),
			Table:  table,
		}
	}

	return groupTables, nil
}

func (c *Client) SortedMatches() ([]data.Match, error) {
	b, err := httpGetBytes(c.httpClient, "https://api.football-data.org/v4/competitions/WC/matches", c.token)
	if err != nil {
		return nil, err
	}

	var p parsedMatches
	if err := json.Unmarshal(b, &p); err != nil {
		return nil, err
	}

	var matches []data.Match
	for _, parsedMatch := range p.Matches {
		date, err := time.Parse(time.RFC3339, parsedMatch.UTCDate)
		if err != nil {
			return nil, err
		}

		matches = append(matches, data.Match{
			ID:            parsedMatch.ID,
			HomeTeamCode:  parsedMatch.HomeTeam.TLA,
			AwayTeamCode:  parsedMatch.AwayTeam.TLA,
			Date:          date.UTC(),
			Stage:         stage(parsedMatch.Stage),
			Status:        status(parsedMatch.Status),
			Venue:         "N/A",
			HomeTeamScore: uint64(parsedMatch.Score.FullTime.Home),
			AwayTeamScore: uint64(parsedMatch.Score.FullTime.Away),
		})
	}

	sort.Slice(matches, func(i, j int) bool {
		if matches[i].Date.Equal(matches[j].Date) {
			return matches[i].ID < matches[j].ID
		}
		return matches[i].Date.Before(matches[j].Date)
	})

	localMatches, err := local.SortedLocalMatches()
	if err != nil {
		return nil, err
	}

	for i, localMatch := range localMatches {
		matches[i] = localMatch
	}

	return matches, nil
}

func status(status string) data.Status {
	statusMappings := map[string]data.Status{
		// Scheduled
		"SCHEDULED": data.StatusScheduled,
		"TIMED":     data.StatusScheduled,
		"CANCELLED": data.StatusScheduled,
		"POSTPONED": data.StatusScheduled,

		// Live
		"SUSPENDED": data.StatusLive,
		"IN_PLAY":   data.StatusLive,
		"PAUSED":    data.StatusLive,

		// Finished
		"AWARDED":  data.StatusFinished,
		"FINISHED": data.StatusFinished,
	}

	if status, ok := statusMappings[status]; ok {
		return status
	}

	return data.StatusScheduled
}

func stage(stageStr string) string {
	stageMappings := map[string]data.Stage{
		"FINAL":          data.StageFinal,
		"THIRD_PLACE":    data.StageThird,
		"SEMI_FINALS":    data.StageSemi,
		"QUARTER_FINALS": data.StageQuarter,
		"LAST_16":        data.StageLast16,
		"GROUP_STAGE":    data.StageGroup,
	}

	if stage, ok := stageMappings[stageStr]; ok {
		return string(stage)
	}

	return stageStr
}

func httpGetBytes(client mockableHttpClient, url string, token string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Auth-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
