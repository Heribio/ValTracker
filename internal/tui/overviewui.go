package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type overViewState struct {
}

func (m model) overViewUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "<":
			m = m.SwitchPage(matchListPage)
			return m, nil
		}
	}
	var cmd tea.Cmd
	return m, cmd
}

func (m model) overViewView() string {
	return "Overview Page"
}
