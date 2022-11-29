package group

import (
	"fmt"
	"strconv"

	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/data"
	"github.com/charmbracelet/lipgloss"
)

var (
	teamColWidth  = 20
	pointColWidth = 5
	groupWidth    = 20 + (8 * pointColWidth)

	groupNameStyle = lipgloss.NewStyle().Width(groupWidth).Bold(true).Align(lipgloss.Center)
	containerStyle = lipgloss.NewStyle().Width(groupWidth).Align(lipgloss.Left)

	headingsStyle     = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, false, true, false)
	pointColStyle     = lipgloss.NewStyle().Width(pointColWidth).Align(lipgloss.Center)
	pointBoldColStyle = lipgloss.NewStyle().Bold(true).Width(pointColWidth).Align(lipgloss.Center)
	teamColStyle      = lipgloss.NewStyle().Width(teamColWidth)
)

func Group(groupTable data.GroupTable) string {
	var s string
	s += renderGroupName(groupTable.Letter) + "\n"
	s += "\n"
	s += headingsStyle.Render(renderHeadings()) + "\n"
	for _, team := range groupTable.Table {
		s += renderRow(team) + "\n"
	}
	return containerStyle.Render(s)
}

func renderGroupName(letter string) string {
	return groupNameStyle.Render(fmt.Sprintf("Group %s", letter))
}

func renderHeadings() string {
	var s string
	s += teamColStyle.Render("Team")
	s += pointColStyle.Render("MP")
	s += pointColStyle.Render("W")
	s += pointColStyle.Render("D")
	s += pointColStyle.Render("L")
	s += pointColStyle.Render("GF")
	s += pointColStyle.Render("GA")
	s += pointColStyle.Render("GD")
	s += pointColStyle.Render("Pts")
	return s
}

func renderRow(team data.GroupTableTeam) string {
	var s string
	s += teamColStyle.Render(renderTeam(team))
	s += pointColStyle.Render(strconv.FormatInt(int64(team.MatchesPlayed), 10))
	s += pointColStyle.Render(strconv.FormatInt(int64(team.Wins), 10))
	s += pointColStyle.Render(strconv.FormatInt(int64(team.Draws), 10))
	s += pointColStyle.Render(strconv.FormatInt(int64(team.Losses), 10))
	s += pointColStyle.Render(strconv.FormatInt(int64(team.GoalsFor), 10))
	s += pointColStyle.Render(strconv.FormatInt(int64(team.GoalsAgainst), 10))
	s += pointColStyle.Render(strconv.FormatInt(int64(team.GoalsDifferential), 10))
	s += pointBoldColStyle.Render(strconv.FormatInt(int64(team.Points), 10))
	return s
}

func renderTeam(team data.GroupTableTeam) string {
	teamInfo, ok := data.TeamInfoByCode[team.Code]
	if !ok {
		return "?? " + team.Code
	}

	firstSquare := lipgloss.NewStyle().SetString("■").Foreground(lipgloss.Color(teamInfo.FirstColor)).String()
	secondSquare := lipgloss.NewStyle().SetString("■").Foreground(lipgloss.Color(teamInfo.SecondColor)).String()
	return firstSquare + secondSquare + " " + teamInfo.Name
}
