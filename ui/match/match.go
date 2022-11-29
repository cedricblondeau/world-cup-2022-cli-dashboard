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

	shirtNumberStyle = lipgloss.NewStyle().Width(2)
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
			renderTeamCol(fourColsWidth, params.Match.HomeTeamCode, params.Match.HomeTeamLineup, true),
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
			renderTeamCol(fourColsWidth, params.Match.AwayTeamCode, params.Match.AwayTeamLineup, false),
		).String(),
	)

	return s
}

func renderTeamCol(width int, countryCode string, lineup []data.Player, reverse bool) string {
	teamInfo, ok := data.TeamInfoByCode[countryCode]
	teamName := countryCode
	if ok {
		teamName = teamInfo.Name
	}
	var renderedLineup string
	for _, player := range lineup {
		if reverse {
			renderedLineup += player.Name + " " + shirtNumberStyle.Render(strconv.FormatInt(int64(player.ShirtNumber), 10)) + "\n"
		} else {
			renderedLineup += shirtNumberStyle.Render(strconv.FormatInt(int64(player.ShirtNumber), 10)) + " " + player.Name + "\n"
		}
	}

	if reverse {
		var s string
		s += lipgloss.NewStyle().Bold(true).Width(width).Align(lipgloss.Right).Render(teamName) + "\n\n"
		s += lipgloss.NewStyle().Width(width).Align(lipgloss.Right).SetString(flags.Render(countryCode)).String() + "\n"
		s += lipgloss.NewStyle().Width(width).Align(lipgloss.Right).SetString(renderedLineup).String() + "\n"
		return s
	}

	var s string
	s += lipgloss.NewStyle().Bold(true).Width(width).Align(lipgloss.Left).Render(teamName) + "\n\n"
	s += lipgloss.NewStyle().Width(width).Align(lipgloss.Left).SetString(flags.Render(countryCode)).String() + "\n"
	s += lipgloss.NewStyle().Width(width).Align(lipgloss.Left).SetString(renderedLineup).String() + "\n"
	return s
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
