package groups

import (
	"fmt"
	"strconv"

	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/data"
	"github.com/charmbracelet/lipgloss"
)

var (
	groupStyle = lipgloss.NewStyle().
			Align(lipgloss.Center).
			Width(12)

	groupNameStyle = lipgloss.NewStyle().Bold(true).Align(lipgloss.Center)
)

func Groups(groups []data.GroupTable) string {
	renderedGroups := make([]string, len(groups))
	for i, group := range groups {
		renderedGroups[i] = groupStyle.Render(renderGroupName(group.Letter) + "\n\n" + renderGroupTable(group.Table))
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		renderedGroups...,
	)
}

func renderGroupName(letter string) string {
	return groupNameStyle.Render(fmt.Sprintf("Group %s", letter))
}

func renderGroupTable(table []data.GroupTableTeam) string {
	var s string
	for _, team := range table {
		teamInfo, ok := data.TeamInfoByCode[team.Code]
		firstColor := "#000000"
		secondColor := "#ffffff"
		if ok {
			firstColor = teamInfo.FirstColor
			secondColor = teamInfo.SecondColor
		}

		firstColorSquare := lipgloss.NewStyle().SetString("■").Foreground(lipgloss.Color(firstColor)).String()
		secondColorSquare := lipgloss.NewStyle().SetString("■").Foreground(lipgloss.Color(secondColor)).String()
		s += firstColorSquare + secondColorSquare + " " + team.Code + " " + strconv.FormatInt(int64(team.Points), 10) + "\n"
	}
	return s
}
