package nav

import (
	"fmt"
	"math"
	"time"

	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/data"
	"github.com/charmbracelet/lipgloss"
)

const (
	matchItemWidth = 22
)

var (
	matchItemStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#3E6D9C")).
			Foreground(lipgloss.Color("#FAFAFA")).
			Align(lipgloss.Center).
			Width(matchItemWidth).
			Padding(1, 0)

	selectedMatchItemStyle = matchItemStyle.Copy().
				Bold(true).
				Foreground(lipgloss.Color("#FAFAFA")).
				Background(lipgloss.Color("#282A3A"))

	stageStyle = lipgloss.NewStyle().
			Bold(true).
			Background(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FAFAFA"}).
			Foreground(lipgloss.AdaptiveColor{Light: "#EEEEEE", Dark: "#333333"})
)

type NavParams struct {
	Index   int
	Matches []data.Match
	Width   int
}

func Nav(params NavParams) string {
	if params.Index < 0 || params.Index > len(params.Matches)-1 {
		return fmt.Sprintf("Index=%d must be >=0 and <= len(matches)=%d", params.Index, len(params.Matches))
	}

	pageSize := int(math.Floor(float64(params.Width) / float64(matchItemWidth)))
	totalPages := int(math.Ceil(float64(len(params.Matches)) / float64(pageSize)))
	currentPage := int(math.Floor(float64(params.Index) / float64(pageSize)))
	pageIndex := params.Index % pageSize

	matchesStartIndex := currentPage * pageSize
	matchesToRender := params.Matches[matchesStartIndex:min(len(params.Matches), matchesStartIndex+pageSize+1)]

	renderedMatches := make([]string, len(matchesToRender))
	for i, match := range matchesToRender {
		if i == pageIndex {
			renderedMatches[i] = selectedMatchItemStyle.Render(renderMatch(match))
		} else {
			renderedMatches[i] = matchItemStyle.Render(renderMatch(match))
		}
	}

	var s string

	s += lipgloss.JoinHorizontal(
		lipgloss.Top,
		renderedMatches...,
	)
	s += "\n\n"

	s += lipgloss.NewStyle().Width(params.Width).SetString(
		renderPagination(totalPages, currentPage),
	).Align(lipgloss.Center).String()

	return s
}

func renderMatch(m data.Match) string {
	return renderDatetime(m) + "\n" + renderTeams(m) + " " + stageStyle.Render(renderStage(m.Stage, m.HomeTeamCode))
}

func renderDatetime(m data.Match) string {
	if m.Status == data.StatusLive {
		return fmt.Sprintf("LIVE %s", m.Minute)
	}

	localMatchDate := m.Date.Local()

	timeFromNow := time.Until(localMatchDate)
	if timeFromNow > 0 && timeFromNow < time.Duration(6)*(time.Hour*24) {
		return localMatchDate.Format("Monday 3:04 PM")
	}

	return localMatchDate.Format("Jan 2 3:04 PM")
}

func renderTeams(m data.Match) string {
	if m.Status == data.StatusFinished || m.Status == data.StatusLive {
		return fmt.Sprintf("%s %d-%d %s", m.HomeTeamCode, m.HomeTeamScore, m.AwayTeamScore, m.AwayTeamCode)
	}

	return m.HomeTeamCode + "-" + m.AwayTeamCode
}

func renderStage(stage string, homeTeamCode string) string {
	if stage == string(data.StageGroup) {
		if homeTeamInfo, ok := data.TeamInfoByCode[homeTeamCode]; ok {
			return homeTeamInfo.Group
		} else {
			return "?"
		}
	}

	return stage
}

func renderPagination(totalPages int, currentPage int) string {
	var s string
	for i := 0; i < totalPages; i++ {
		if i == currentPage {
			s += lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Padding(0, 1).Render("⬤")
			continue
		}
		s += lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Padding(0, 1).Render("⬤")
	}
	return s
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
