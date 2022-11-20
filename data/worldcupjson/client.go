package worldcupjson

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/data"
)

type mockableHttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	httpClient mockableHttpClient
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

func NewMockClient() *Client {
	return &Client{
		httpClient: &mockHttpClient{},
	}
}

func (c *Client) Name() string {
	return "worldcupjson.net"
}

func (c *Client) Matches() ([]data.Match, error) {
	b, err := httpGetBytes(c.httpClient, "https://worldcupjson.net/matches?details=true")
	if err != nil {
		return nil, err
	}

	var parsedMatches []parsedMatch
	if err := json.Unmarshal(b, &parsedMatches); err != nil {
		return nil, err
	}

	var matches []data.Match
	for _, parsedMatch := range parsedMatches {
		date, err := time.Parse(time.RFC3339, parsedMatch.Datetime)
		if err != nil {
			return nil, err
		}

		homeTeamEvents := make([]data.Event, len(parsedMatch.HomeTeamEvents))
		for i, event := range parsedMatch.HomeTeamEvents {
			homeTeamEvents[i] = data.Event{
				Type:   eventType(event.TypeOfEvent),
				Player: event.Player,
				Minute: event.Time,
			}
		}

		awayTeamEvents := make([]data.Event, len(parsedMatch.AwayTeamEvents))
		for i, event := range parsedMatch.AwayTeamEvents {
			awayTeamEvents[i] = data.Event{
				Type:   eventType(event.TypeOfEvent),
				Player: event.Player,
				Minute: event.Time,
			}
		}

		matches = append(matches, data.Match{
			HomeTeamCode:   parsedMatch.HomeTeam.Country,
			AwayTeamCode:   parsedMatch.AwayTeam.Country,
			Date:           date.UTC(),
			Stage:          stage(parsedMatch.StageName),
			Status:         status(parsedMatch.Status),
			Venue:          parsedMatch.Venue + " (" + parsedMatch.Location + ")",
			HomeTeamScore:  uint64(parsedMatch.HomeTeam.Goals),
			AwayTeamScore:  uint64(parsedMatch.AwayTeam.Goals),
			HomeTeamEvents: homeTeamEvents,
			AwayTeamEvents: awayTeamEvents,
			Minute:         parsedMatch.Minute,
		})
	}

	return matches, nil
}

func (c *Client) GroupTables() ([]data.GroupTable, error) {
	b, err := httpGetBytes(c.httpClient, "https://worldcupjson.net/teams")
	if err != nil {
		return nil, err
	}

	var p parsedTeams
	if err := json.Unmarshal(b, &p); err != nil {
		return nil, err
	}

	groupTables := make([]data.GroupTable, len(p.Groups))
	for i, group := range p.Groups {
		table := make([]data.GroupTableTeam, len(group.Teams))
		for j, team := range group.Teams {
			table[j] = data.GroupTableTeam{
				Code:   team.Country,
				Points: team.GroupPoints,
			}
		}

		groupTables[i] = data.GroupTable{
			Letter: group.Letter,
			Table:  table,
		}
	}

	return groupTables, nil
}

func eventType(eventTypeStr string) string {
	eventTypeMappings := map[string]data.EventType{
		"substitution-in":    data.EventTypeSubIn,
		"substitution-out":   data.EventTypeSubOut,
		"yellow-card":        data.EventTypeYellowCard,
		"yellow-card-second": data.EventTypeSeconYellowCard,
		"red-card":           data.EventTypeRedCard,
		"goal":               data.EventTypeGoal,
		"penalty-kick":       data.EventTypePenalyKick,
		"goal-penalty":       data.EventTypePenaltyGoal,
		"goal-own":           data.EventTypeOwnGoal,
	}

	if eventType, ok := eventTypeMappings[eventTypeStr]; ok {
		return string(eventType)
	}

	return eventTypeStr
}

func status(status string) data.Status {
	statusMappings := map[string]data.Status{
		"future_scheduled": data.StatusScheduled,
		"in progress":      data.StatusLive,
		"completed":        data.StatusFinished,
	}

	if status, ok := statusMappings[status]; ok {
		return status
	}

	return data.StatusScheduled
}

func stage(stageStr string) string {
	stageMappings := map[string]data.Stage{
		"Final":                    data.StageFinal,
		"Play-off for third place": data.StageThird,
		"Semi-final":               data.StageSemi,
		"Semi-finals":              data.StageSemi,
		"Quarter-final":            data.StageQuarter,
		"Quarter-finals":           data.StageQuarter,
		"Round of 16":              data.StageLast16,
		"First stage":              data.StageGroup,
	}

	if stage, ok := stageMappings[stageStr]; ok {
		return string(stage)
	}

	return stageStr
}

func httpGetBytes(client mockableHttpClient, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
