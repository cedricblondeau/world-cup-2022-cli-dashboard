package playerstats

import "github.com/cedricblondeau/world-cup-2022-cli-dashboard/data"

type PlayerStats struct {
	GoalsByPlayer       map[string]int64
	YellowCardsByPlayer map[string]int64
	RedCardsByPlayer    map[string]int64
}

func PlayerStatsByTeam(matches []data.Match) map[string]PlayerStats {
	eventsByTeam := make(map[string][]data.Event)
	for _, m := range matches {
		eventsByTeam[m.HomeTeamCode] = append(eventsByTeam[m.HomeTeamCode], m.HomeTeamEvents...)
		eventsByTeam[m.AwayTeamCode] = append(eventsByTeam[m.AwayTeamCode], m.AwayTeamEvents...)
	}

	playerStatsByTeam := make(map[string]PlayerStats)
	for team, events := range eventsByTeam {
		playerStatsByTeam[team] = playerStats(events)
	}
	return playerStatsByTeam
}

func playerStats(events []data.Event) PlayerStats {
	playerStats := PlayerStats{
		GoalsByPlayer:       make(map[string]int64),
		YellowCardsByPlayer: make(map[string]int64),
		RedCardsByPlayer:    make(map[string]int64),
	}

	for _, e := range events {
		if e.Canceled {
			continue
		}

		if e.Type == data.EventTypeGoal || e.Type == data.EventTypePenaltyKickGoal {
			playerStats.GoalsByPlayer[e.Player]++
			continue
		}

		if e.Type == data.EventTypeYellowCard || e.Type == data.EventTypeSeconYellowCard {
			playerStats.YellowCardsByPlayer[e.Player]++
			continue
		}

		if e.Type == data.EventTypeRedCard {
			playerStats.RedCardsByPlayer[e.Player]++
			continue
		}
	}

	return playerStats
}
