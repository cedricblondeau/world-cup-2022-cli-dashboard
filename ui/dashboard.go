package ui

import (
	"fmt"
	"time"

	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/data"
	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/ui/groups"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type intervalRefreshMsg time.Time

const refreshInterval = time.Duration(10) * time.Second

type dashboard struct {
	dataFetcher         dataFetcher
	dataFetchErr        error
	dataFetchLastUpdate time.Time
	dataFetchLoading    bool
	dataFetchSpinner    spinner.Model

	groupTables []data.GroupTable

	width, height int
}

func NewDashboard(fetcher dataFetcher) tea.Model {
	s := spinner.New()
	s.Spinner = spinner.Globe

	return &dashboard{
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
		return m, refreshCmd()

	case intervalRefreshMsg:
		m.dataFetchLoading = true
		return m, dataFetchCmd(m.dataFetcher)

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

	if m.dataFetchLoading {
		return fullScreenMsgStyle.Render(m.dataFetchSpinner.View() + " " + fmt.Sprintf("Loading data from %s...", m.dataFetcher.Name()))
	}

	if m.dataFetchErr != nil {
		return fullScreenMsgStyle.Render(fmt.Sprintf("❌ HTTP request failed with err: %v...", m.dataFetchErr.Error()))
	}

	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), true, false, false, false).
		PaddingTop(1).
		PaddingBottom(1).
		Width(m.width).
		Align(lipgloss.Center).
		SetString(groups.Groups(m.groupTables)).
		String()
}

func refreshCmd() tea.Cmd {
	return tea.Tick(
		refreshInterval,
		func(t time.Time) tea.Msg {
			return intervalRefreshMsg(t)
		},
	)
}
