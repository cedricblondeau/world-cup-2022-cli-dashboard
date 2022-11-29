package data

import "time"

type EventType string

const EventTypeGoal = "Goal"
const EventTypePenalyKick = "Penalty Kick"
const EventTypeYellowCard = "Yellow Card"
const EventTypeSeconYellowCard = "Second Yellow Card"
const EventTypeRedCard = "Red Card"
const EventTypeSubIn = "Substitution In"
const EventTypeSubOut = "Substitution Out"
const EventTypePenaltyGoal = "Penalty Goal"
const EventTypeOwnGoal = "Own Goal"

type Stage string

const StageGroup Stage = "Group"
const StageLast16 Stage = "1/8"
const StageQuarter Stage = "1/4"
const StageSemi Stage = "1/2"
const StageThird Stage = "3rd"
const StageFinal Stage = "Final"

type Status string

const StatusScheduled = "Scheduled"
const StatusLive = "Live"
const StatusFinished = "Finished"

type GroupTable struct {
	Letter string
	Table  []GroupTableTeam
}

type GroupTableTeam struct {
	Code              string
	Points            int
	Wins              int
	Draws             int
	Losses            int
	MatchesPlayed     int
	GoalsFor          int
	GoalsAgainst      int
	GoalsDifferential int
}

type Match struct {
	ID             int
	HomeTeamCode   string
	AwayTeamCode   string
	Date           time.Time
	Venue          string
	HomeTeamScore  uint64
	AwayTeamScore  uint64
	WinnerTeamCode string
	Minute         string
	HomeTeamEvents []Event
	AwayTeamEvents []Event
	Status         Status
	HomeTeamLineup []Player
	AwayTeamLineup []Player

	// Stage is a Stage when the type is known - otherwise a string
	Stage string
}

type Event struct {
	// MatchEventType is a MatchEventTypeGoal when the type is known - otherwise a string
	Type string

	Minute string
	Player string
}

type TeamInfo struct {
	Name        string
	Group       string
	FirstColor  string
	SecondColor string
}

type Player struct {
	Name        string
	ShirtNumber int
}
