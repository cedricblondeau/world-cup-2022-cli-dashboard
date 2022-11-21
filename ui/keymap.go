package ui

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Left  key.Binding
	Quit  key.Binding
	Right key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Left, k.Right, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}
