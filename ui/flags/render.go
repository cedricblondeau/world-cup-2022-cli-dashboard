package flags

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func Render(countryCode string) string {
	countryFlag, ok := countryFlags[countryCode]
	if !ok {
		return ""
	}

	height := len(countryFlag)
	b := strings.Builder{}
	for y := 0; y < height; y += 2 {
		if y >= height-1 {
			// we expect a height of 14px, so it should never happen
			continue
		}

		for x := 0; x < 25; x++ {
			color1 := countryFlag[y][x]
			color2 := countryFlag[y+1][x]

			b.WriteString(lipgloss.
				NewStyle().
				SetString("▀").
				Foreground(lipgloss.Color(color1)).
				Background(lipgloss.Color(color2)).
				String(),
			)
		}
		b.WriteString("\n")
	}

	return b.String()
}
