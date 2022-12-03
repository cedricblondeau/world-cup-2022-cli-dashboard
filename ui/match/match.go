package match

import (
	"fmt"
	"math"
	"strconv"

	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/data"
	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/ui/bigtext"
	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/ui/flags"
	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/ui/playerstats"
	"github.com/charmbracelet/lipgloss"
)

type MatchParams struct {
	BigText           *bigtext.BigText
	PlayerStatsByTeam map[string]playerstats.PlayerStats
	Match             data.Match
	Width             int
}

var (
	scoreDigitStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FAFAFA"}).
			Height(11).
			Align(lipgloss.Left)

	eventTypeStyle = lipgloss.NewStyle().
			Background(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FAFAFA"}).
			Foreground(lipgloss.AdaptiveColor{Light: "#FAFAFA", Dark: "#000000"})

	canceledEventStyle = lipgloss.NewStyle().Strikethrough(true)

	shirtNumberStyle = lipgloss.NewStyle().Width(2)

	yellowCard = lipgloss.NewStyle().SetString("■").Foreground(lipgloss.Color("#FFFF00")).String()
	redCard    = lipgloss.NewStyle().SetString("■").Foreground(lipgloss.Color("#FF0000")).String()
)

func Match(params MatchParams) string {
	threeColsWidth := int(math.Floor(float64(params.Width) / 3))
	fourColsWidth := int(math.Floor(float64(params.Width) / 4))

	s := "\n"

	s += lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().Width(threeColsWidth).Align(lipgloss.Left).SetString(params.Match.Venue).String(),
		lipgloss.NewStyle().Width(threeColsWidth).Align(lipgloss.Center).SetString(renderStatus(params.Match)).String(),
		lipgloss.NewStyle().Width(threeColsWidth).Align(lipgloss.Right).SetString(params.Match.Date.Local().Format("Jan 2 3:04 PM")).String(),
	)

	s += "\n\n\n"

	s += lipgloss.JoinHorizontal(
		lipgloss.Top,

		lipgloss.NewStyle().Width(fourColsWidth).Align(lipgloss.Center).SetString(
			renderTeamCol(fourColsWidth, params.Match.HomeTeamCode, params.Match.HomeTeamLineup, params.PlayerStatsByTeam),
		).String(),

		lipgloss.NewStyle().Width(fourColsWidth*2).SetString(
			lipgloss.JoinVertical(
				lipgloss.Top,
				lipgloss.JoinHorizontal(
					lipgloss.Top,
					lipgloss.NewStyle().Width(fourColsWidth-10).Align(lipgloss.Right).SetString(
						scoreDigitStyle.Render(params.BigText.Char(strconv.FormatUint(params.Match.HomeTeamScore, 10))),
					).String(),
					lipgloss.NewStyle().Width(10*2).String(),
					lipgloss.NewStyle().Width(fourColsWidth-10).Align(lipgloss.Left).SetString(
						scoreDigitStyle.Render(params.BigText.Char(strconv.FormatUint(params.Match.AwayTeamScore, 10))),
					).String(),
				),
				"",
				lipgloss.JoinHorizontal(
					lipgloss.Top,
					lipgloss.NewStyle().Width(fourColsWidth-10).Align(lipgloss.Right).SetString(renderEvents(params.Match.HomeTeamEvents, true)).String(),
					lipgloss.NewStyle().Width(10*2).String(),
					lipgloss.NewStyle().Width(fourColsWidth-10).Align(lipgloss.Left).SetString(renderEvents(params.Match.AwayTeamEvents, false)).String(),
				),
			),
		).String(),

		lipgloss.NewStyle().Width(fourColsWidth).Align(lipgloss.Center).SetString(
			renderTeamCol(fourColsWidth, params.Match.AwayTeamCode, params.Match.AwayTeamLineup, params.PlayerStatsByTeam),
		).String(),
	)

	return s
}

func renderTeamCol(
	width int,
	countryCode string,
	lineup []data.Player,
	playerStatsByTeam map[string]playerstats.PlayerStats,
) string {
	teamInfo, ok := data.TeamInfoByCode[countryCode]
	teamName := countryCode
	if ok {
		teamName = teamInfo.Name
	}

	playerStats, ok := playerStatsByTeam[countryCode]
	if !ok {
		playerStats = playerstats.PlayerStats{
			GoalsByPlayer:       make(map[string]int64),
			YellowCardsByPlayer: make(map[string]int64),
			RedCardsByPlayer:    make(map[string]int64),
		}
	}
	renderedLineup := renderLineup(lineup, playerStats)

	var s string
	s += lipgloss.NewStyle().Bold(true).Width(width).Align(lipgloss.Center).Render(teamName) + "\n\n"
	s += lipgloss.NewStyle().Width(width).Align(lipgloss.Center).SetString(flags.Render(countryCode)).String() + "\n"
	s += lipgloss.NewStyle().Width(width).Align(lipgloss.Center).SetString(renderedLineup).String() + "\n"
	return s
}

func renderLineup(
	lineup []data.Player,
	playerStats playerstats.PlayerStats,
) string {
	if len(lineup) == 0 {
		return ""
	}

	headingsStyle := lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, false, true, false)
	defaultColStyle := lipgloss.NewStyle().Width(3).Align(lipgloss.Center)
	playerColStyle := lipgloss.NewStyle().Width(27).Align(lipgloss.Left)

	var headings string
	headings += defaultColStyle.Render("#")
	headings += playerColStyle.Render("Player")
	headings += defaultColStyle.Render("⚽")
	headings += defaultColStyle.Render(yellowCard)
	headings += defaultColStyle.Render(redCard)

	var rows string
	for _, player := range lineup {
		var row string
		row += defaultColStyle.Render(strconv.FormatInt(int64(player.ShirtNumber), 10))
		row += playerColStyle.Render(player.Name)
		row += defaultColStyle.Render(strconv.FormatInt(playerStats.GoalsByPlayer[player.Name], 10))
		row += defaultColStyle.Render(strconv.FormatInt(playerStats.YellowCardsByPlayer[player.Name], 10))
		row += defaultColStyle.Render(strconv.FormatInt(playerStats.RedCardsByPlayer[player.Name], 10))

		rows += row + "\n"
	}

	return headingsStyle.Render(headings) + "\n" + rows
}

func renderStatus(match data.Match) string {
	if match.Status == data.StatusLive {
		return fmt.Sprintf("LIVE %s", match.Minute)
	}

	return string(match.Status)
}

func renderEvents(events []data.Event, reverse bool) string {
	var s string
	for _, event := range events {
		s += renderEvent(event, reverse)
		s += "\n"
	}
	return s
}

func renderEvent(event data.Event, reverse bool) string {
	if reverse {
		if event.Canceled {
			return canceledEventStyle.Render(event.Player+" "+event.Minute) + " " + renderEventType(event.Type)
		}
		return event.Player + " " + event.Minute + " " + renderEventType(event.Type)
	}

	if event.Canceled {
		return renderEventType(event.Type) + " " + canceledEventStyle.Render(event.Minute+" "+event.Player)
	}
	return renderEventType(event.Type) + " " + event.Minute + " " + event.Player
}

func renderEventType(eventType string) string {
	eventTypeIcons := map[string]string{
		string(data.EventTypeYellowCard):      yellowCard,
		string(data.EventTypeSeconYellowCard): yellowCard,
		string(data.EventTypeRedCard):         redCard,
		string(data.EventTypeSubIn):           "►",
		string(data.EventTypeSubOut):          "◄",
	}

	icon, ok := eventTypeIcons[eventType]
	if ok {
		return icon
	}

	return eventTypeStyle.SetString(eventType).String()
}
