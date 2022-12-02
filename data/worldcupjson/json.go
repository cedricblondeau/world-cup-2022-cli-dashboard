package worldcupjson

type parsedTeams struct {
	Groups []struct {
		Letter string `json:"letter"`
		Teams  []struct {
			Country           string `json:"country"`
			GamesPlayed       int    `json:"games_played"`
			Wins              int    `json:"wins"`
			Draws             int    `json:"draws"`
			Losses            int    `json:"losses"`
			GoalsFor          int    `json:"goals_for"`
			GoalsAgainst      int    `json:"goals_against"`
			GoalsDifferential int    `json:"goal_differential"`
			GroupPoints       int    `json:"group_points"`
		}
	} `json:"groups"`
}

type parsedMatchTeam struct {
	Country string `json:"country"`
	Goals   uint   `json:"goals"`
}

type parsedEvent struct {
	TypeOfEvent string `json:"type_of_event"`
	Player      string `json:"player"`
	Time        string `json:"time"`
	ExtraInfo   string `json:"extra_info"`
}

type parsedSubEventInfo struct {
	PlayerOn  string `json:"player_on"`
	PlayerOff string `json:"player_off"`
}

type parsedMatch struct {
	ID             int             `json:"id"`
	Venue          string          `json:"venue"`
	Location       string          `json:"location"`
	Status         string          `json:"status"`
	StageName      string          `json:"stage_name"`
	HomeTeam       parsedMatchTeam `json:"home_team"`
	AwayTeam       parsedMatchTeam `json:"away_team"`
	Datetime       string          `json:"datetime"`
	Minute         string          `json:"time"` // Remove and use new field
	HomeTeamEvents []parsedEvent   `json:"home_team_events"`
	AwayTeamEvents []parsedEvent   `json:"away_team_events"`
	HomeTeamLineup parsedLineup    `json:"home_team_lineup"`
	AwayTeamLineup parsedLineup    `json:"away_team_lineup"`
}

type parsedLineup struct {
	StartingEleven []parsedPlayer `json:"starting_eleven"`
}

type parsedPlayer struct {
	Name        string `json:"name"`
	ShirtNumber int    `json:"shirt_number"`
}
