package match

import (
	"fmt"
	"math"
	"strconv"

	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/data"
	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/ui/bigtext"
	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/ui/flags"
	"github.com/charmbracelet/lipgloss"
)

type MatchParams struct {
	BigText *bigtext.BigText
	Match   data.Match
	Width   int
}

var (
	scoreDigitStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FAFAFA"}).
			Height(11).
			Align(lipgloss.Left)

	eventTypeStyle = lipgloss.NewStyle().
			Background(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FAFAFA"}).
			Foreground(lipgloss.AdaptiveColor{Light: "#FAFAFA", Dark: "#000000"})
)

func Match(params MatchParams) string {
	twoColsWidth := int(math.Floor(float64(params.Width) / 2))
	threeColsWidth := int(math.Floor(float64(params.Width) / 3))
	fourColsWidth := int(math.Floor(float64(params.Width) / 4))

	s := "\n"

	s += lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().Width(threeColsWidth).Align(lipgloss.Left).SetString(params.Match.Venue).String(),
		lipgloss.NewStyle().Width(threeColsWidth).Align(lipgloss.Center).SetString(renderStatus(params.Match)).String(),
		lipgloss.NewStyle().Width(threeColsWidth).Align(lipgloss.Right).SetString(params.Match.Date.Format("Jan 2 3:04 PM")).String(),
	)

	s += "\n\n\n"

	s += lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().Width(fourColsWidth).Align(lipgloss.Center).SetString(
			renderTeam(params.Match.HomeTeamCode),
		).String(),
		lipgloss.NewStyle().Width(fourColsWidth).Align(lipgloss.Center).SetString(
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				scoreDigitStyle.Render(params.BigText.Char(strconv.FormatUint(params.Match.HomeTeamScore, 10))),
			),
		).String(),
		lipgloss.NewStyle().Width(fourColsWidth).Align(lipgloss.Center).SetString(
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				scoreDigitStyle.Render(params.BigText.Char(strconv.FormatUint(params.Match.AwayTeamScore, 10))),
			),
		).String(),
		lipgloss.NewStyle().Width(fourColsWidth).Align(lipgloss.Center).SetString(
			renderTeam(params.Match.AwayTeamCode),
		).String(),
	)

	s += "\n\n"

	s += lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().Width(twoColsWidth-8).Align(lipgloss.Right).SetString(renderEvents(params.Match.HomeTeamEvents, true)).String(),
		lipgloss.NewStyle().Width(8*2).String(),
		lipgloss.NewStyle().Width(twoColsWidth-8).Align(lipgloss.Left).SetString(renderEvents(params.Match.AwayTeamEvents, false)).String(),
	)

	return s
}

func renderTeam(countryCode string) string {
	teamInfo, ok := data.TeamInfoByCode[countryCode]
	teamName := countryCode
	if ok {
		teamName = teamInfo.Name
	}

	return lipgloss.JoinVertical(
		lipgloss.Center,
		lipgloss.NewStyle().SetString(teamName).Bold(true).String(),
		"\n",
		flags.Render(countryCode),
	)
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
		return event.Player + " " + event.Minute + " " + renderEventType(event.Type)
	}
	return renderEventType(event.Type) + " " + event.Minute + " " + event.Player
}

func renderEventType(eventType string) string {
	eventTypeIcons := map[string]string{
		string(data.EventTypeYellowCard):      lipgloss.NewStyle().SetString("■").Foreground(lipgloss.Color("#FFFF00")).String(),
		string(data.EventTypeSeconYellowCard): lipgloss.NewStyle().SetString("■").Foreground(lipgloss.Color("#FFFF00")).String(),
		string(data.EventTypeRedCard):         lipgloss.NewStyle().SetString("■").Foreground(lipgloss.Color("#FF0000")).String(),
		string(data.EventTypeSubIn):           "►",
		string(data.EventTypeSubOut):          "◄",
	}

	icon, ok := eventTypeIcons[eventType]
	if ok {
		return icon
	}

	return eventTypeStyle.SetString(eventType).String()
}
