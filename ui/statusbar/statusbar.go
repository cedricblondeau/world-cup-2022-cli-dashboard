package statusbar

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
)

type StatusBarParams struct {
	API        string
	Err        error
	LastUpdate time.Time
	Loading    bool
	Spinner    spinner.Model
	Width      int
}

func StatusBar(params StatusBarParams) string {
	api := lipgloss.
		NewStyle().
		SetString("API: " + params.API).
		Background(lipgloss.Color("#7D56F4")).
		PaddingLeft(1).
		PaddingRight(1).
		String()

	lastSync := lipgloss.
		NewStyle().
		SetString("Last synced at " + params.LastUpdate.Format("3:04 PM")).
		Background(lipgloss.Color("#CC5C87")).
		PaddingLeft(1).
		PaddingRight(1).
		String()

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		api,
		lipgloss.NewStyle().
			SetString(message(params)).
			Width(params.Width-lipgloss.Width(api)-lipgloss.Width(lastSync)).
			Background(lipgloss.Color("#454544")).
			PaddingLeft(1).
			PaddingRight(1).
			String(),
		lastSync,
	)
}

func message(params StatusBarParams) string {
	if params.Loading {
		return params.Spinner.View() + " Refreshing..."
	}

	if params.Err != nil {
		return "❌ " + params.Err.Error()
	}

	return "Enjoy the match."
}
