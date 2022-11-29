package bracket

import (
	"fmt"

	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/data"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().Bold(true).Align(lipgloss.Center)
	matchWidth = 13
	matchStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#3E6D9C")).
			Foreground(lipgloss.Color("#FAFAFA")).
			Align(lipgloss.Center).
			Width(matchWidth)
	cupStyle = lipgloss.NewStyle().Width(matchWidth).Align(lipgloss.Center)
)

func Bracket(matches []data.Match) string {
	if len(matches) != 64 {
		return fmt.Sprintf("Cannot display brackets, want 64 matches, got %d", len(matches))
	}

	var s string
	s = titleStyle.Render("Bracket") + "\n"
	s += lipgloss.JoinHorizontal(
		lipgloss.Top,
		leftRoundOfSixteen(matches),
		leftQuarterFinals(matches),
		leftSemiFinal(matches),
		final(matches),
		rightSemiFinal(matches),
		rightQuarterFinals(matches),
		rightRoundOfSixteen(matches),
	)
	return s
}

func leftRoundOfSixteen(matches []data.Match) string {
	var s1 string
	s1 += match(matches[48]) + "\n"
	s1 += "\n"
	s1 += match(matches[49]) + "\n"
	s1 += "\n"
	s1 += match(matches[52]) + "\n"
	s1 += "\n"
	s1 += match(matches[53]) + "\n"

	var s2 string
	s2 += "──╮" + "\n"
	s2 += "  ├─" + "\n"
	s2 += "──╯" + "\n"
	s2 += "\n"
	s2 += "──╮" + "\n"
	s2 += "  ├─" + "\n"
	s2 += "──╯" + "\n"

	return lipgloss.JoinHorizontal(lipgloss.Top, s1, s2)
}

func rightRoundOfSixteen(matches []data.Match) string {
	var s1 string
	s1 += " ╭──" + "\n"
	s1 += "─┤  " + "\n"
	s1 += " ╰──" + "\n"
	s1 += "\n"
	s1 += " ╭──" + "\n"
	s1 += "─┤  " + "\n"
	s1 += " ╰──" + "\n"

	var s2 string
	s2 += match(matches[50]) + "\n"
	s2 += "\n"
	s2 += match(matches[51]) + "\n"
	s2 += "\n"
	s2 += match(matches[54]) + "\n"
	s2 += "\n"
	s2 += match(matches[55]) + "\n"

	return lipgloss.JoinHorizontal(lipgloss.Top, s1, s2)
}

func leftQuarterFinals(matches []data.Match) string {
	var s1 string
	s1 += "\n"
	s1 += match(matches[57]) + "\n"
	s1 += "\n"
	s1 += "\n"
	s1 += "\n"
	s1 += match(matches[56]) + "\n"
	s1 += "\n"

	var s2 string
	s2 += "\n"
	s2 += "──╮" + "\n"
	s2 += "  │" + "\n"
	s2 += "  ├─" + "\n"
	s2 += "  │" + "\n"
	s2 += "──╯" + "\n"
	s2 += "\n"

	return lipgloss.JoinHorizontal(lipgloss.Top, s1, s2)
}

func rightQuarterFinals(matches []data.Match) string {
	var s1 string
	s1 += "\n"
	s1 += " ╭──" + "\n"
	s1 += " │  " + "\n"
	s1 += "─┤  " + "\n"
	s1 += " │  " + "\n"
	s1 += " ╰──" + "\n"
	s1 += "\n"

	var s2 string
	s2 += "\n"
	s2 += match(matches[59]) + "\n"
	s2 += "\n"
	s2 += "\n"
	s2 += "\n"
	s2 += match(matches[58]) + "\n"
	s2 += "\n"

	return lipgloss.JoinHorizontal(lipgloss.Top, s1, s2)
}

func leftSemiFinal(matches []data.Match) string {
	var s string
	s += "\n"
	s += "\n"
	s += "\n"
	s += match(matches[60]) + "──" + "\n"
	s += "\n"
	s += "\n"
	s += "\n"
	return s
}

func rightSemiFinal(matches []data.Match) string {
	var s string
	s += "\n"
	s += "\n"
	s += "\n"
	s += "──" + match(matches[61]) + "\n"
	s += "\n"
	s += "\n"
	s += "\n"
	return s
}

func final(matches []data.Match) string {
	var s string
	s += "\n"
	s += "\n"
	s += "\n"
	s += match(matches[63]) + "\n"
	s += "\n"
	s += cupStyle.Render("🏆") + "\n"
	s += "\n"
	return s
}

func match(m data.Match) string {
	if m.Status == data.StatusFinished || m.Status == data.StatusLive {
		return matchStyle.Render(fmt.Sprintf("%s %d-%d %s", m.HomeTeamCode, m.HomeTeamScore, m.AwayTeamScore, m.AwayTeamCode))
	}

	return matchStyle.Render(m.HomeTeamCode + "-" + m.AwayTeamCode)
}
