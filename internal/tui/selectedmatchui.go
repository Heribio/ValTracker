package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type selectedMatchState struct {
}


func (m model) selectMatchUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c":
			return m, tea.Quit
        case "esc":
            m = m.SwitchPage(matchListPage)
            return m, nil
        }
    }
    var cmd tea.Cmd
	return m, cmd
}

func (m model) selectedMatchView() string {
    return fmt.Sprintf("%s - %s", m.selectedMatch.CharacterName, m.selectedMatch.Mode)
}
