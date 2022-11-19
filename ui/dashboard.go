package ui

import (
	"fmt"
	"time"

	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/data"
	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/ui/bigtext"
	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/ui/groups"
	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/ui/match"
	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/ui/nav"
	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/ui/statusbar"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type intervalRefreshMsg time.Time

const refreshInterval = time.Duration(10) * time.Second

type dashboard struct {
	bigtext *bigtext.BigText

	dataFetcher         dataFetcher
	dataFetchErr        error
	dataFetchLastUpdate time.Time
	dataFetchLoading    bool
	dataFetchSpinner    spinner.Model

	groupTables []data.GroupTable
	matches     []data.Match

	matchIndex        int
	matchIndexChanged bool

	width, height int
}

func NewDashboard(fetcher dataFetcher) tea.Model {
	s := spinner.New()
	s.Spinner = spinner.Globe

	return &dashboard{
		bigtext: bigtext.NewBigText(),

		dataFetcher:      fetcher,
		dataFetchLoading: true,
		dataFetchSpinner: s,
	}
}

func (m *dashboard) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen, m.dataFetchSpinner.Tick, dataFetchCmd(m.dataFetcher))
}

func (m *dashboard) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case dataFetchErrMsg:
		m.dataFetchErr = msg.err
		m.dataFetchLoading = false
		return m, nil

	case dataFetchMsg:
		m.dataFetchLoading = false
		m.dataFetchLastUpdate = time.Now()
		m.groupTables = msg.groupTables
		m.matches = msg.matches
		if !m.matchIndexChanged || m.matchIndex > len(msg.matches)-1 {
			m.matchIndex = pickMatchIndex(msg.matches)
		}
		return m, refreshCmd()

	case intervalRefreshMsg:
		m.dataFetchLoading = true
		return m, dataFetchCmd(m.dataFetcher)

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "right", " ", "d":
			if len(m.matches) == 0 {
				return m, nil
			}

			m.matchIndexChanged = true
			m.matchIndex = min(m.matchIndex+1, len(m.matches)-1)
			return m, nil
		case "left", "a":
			if len(m.matches) == 0 {
				return m, nil
			}

			m.matchIndexChanged = true
			m.matchIndex = max(m.matchIndex-1, 0)
			return m, nil
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.dataFetchSpinner, cmd = m.dataFetchSpinner.Update(msg)
		return m, cmd

	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		return m, nil
	}
	return m, nil
}

func (m *dashboard) View() string {
	if m.width == 0 || m.height == 0 {
		return "Initializing..."
	}

	fullScreenMsgStyle := lipgloss.NewStyle().Width(m.width).Height(m.height).Align(lipgloss.Center, lipgloss.Center)

	minWidth := 102
	minHeight := 35
	if m.width < minWidth || m.height < minHeight {
		return fullScreenMsgStyle.Render(fmt.Sprintf("❌ Need at least %d columns and %d rows to render.", minWidth, minHeight))
	}

	if len(m.matches) == 0 {
		if m.dataFetchLoading {
			return fullScreenMsgStyle.Render(m.dataFetchSpinner.View() + " " + fmt.Sprintf("Loading data from %s...", m.dataFetcher.Name()))
		}

		if m.dataFetchErr != nil {
			return fullScreenMsgStyle.Render(fmt.Sprintf("❌ HTTP request failed with err: %v...", m.dataFetchErr.Error()))
		}

		return fullScreenMsgStyle.Render("❓ HTTP request succeeded but no matches available.")
	}

	navContainer := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, true, false).
		PaddingTop(1).
		PaddingBottom(1).
		SetString(nav.Nav(nav.NavParams{
			Index:    m.matchIndex,
			Matches:  m.matches,
			ShowKeys: !m.matchIndexChanged,
			Width:    m.width,
		})).
		String()

	groupsContainer := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), true, false, false, false).
		PaddingTop(1).
		PaddingBottom(1).
		Width(m.width).
		Align(lipgloss.Center).
		SetString(groups.Groups(m.groupTables)).
		String()

	statusBarContainer := lipgloss.NewStyle().
		SetString(statusbar.StatusBar(statusbar.StatusBarParams{
			API:        m.dataFetcher.Name(),
			Err:        m.dataFetchErr,
			LastUpdate: m.dataFetchLastUpdate,
			Loading:    m.dataFetchLoading,
			Spinner:    m.dataFetchSpinner,
			Width:      m.width,
		})).
		String()

	matchContainerHeight := m.height - lipgloss.Height(navContainer) - lipgloss.Height(groupsContainer) - lipgloss.Height(statusBarContainer)
	matchContainer := lipgloss.NewStyle().
		SetString(match.Match(match.MatchParams{
			BigText: m.bigtext,
			Match:   m.matches[min(m.matchIndex, len(m.matches)-1)],
			Width:   m.width - 1 - 1,
		})).
		Height(matchContainerHeight).
		MaxHeight(matchContainerHeight).
		PaddingLeft(1).
		PaddingRight(1).
		String()

	return navContainer + "\n" + matchContainer + "\n" + groupsContainer + "\n" + statusBarContainer
}

func refreshCmd() tea.Cmd {
	return tea.Tick(
		refreshInterval,
		func(t time.Time) tea.Msg {
			return intervalRefreshMsg(t)
		},
	)
}

func pickMatchIndex(matches []data.Match) int {
	for i, match := range matches {
		if match.Status == data.StatusLive || match.Status == data.StatusScheduled {
			return i
		}
	}

	return len(matches) - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
