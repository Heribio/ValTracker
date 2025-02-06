package tui

import (
	"fmt"

	"github.com/Heribio/ValTracker/internal/valorantapi"
	_ "github.com/Heribio/ValTracker/internal/valorantapi"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type matchListState struct {
    list list.Model
}

type Match valorantapi.Match

func (m matchListState) Init() tea.Cmd {
    return nil
}

func (m model) matchListView() string {
    return docStyle.Render(m.state.matchListPage.list.View())
}

func (m model) matchListUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c":
			return m, tea.Quit
        case "<":
            m = m.SwitchPage(loginPage)
            return m, nil
       }
    }

	m.state.matchListPage.list, _= m.state.matchListPage.list.Update(msg)
	return m, nil
}

func (m Match) Title() string {
    return fmt.Sprintf("%s - %s", m.MapName, m.Mode)
}

func (m Match) Description() string {
    return fmt.Sprintf("Character: %s | Kills: %d | Deaths: %d", m.CharacterName, m.Kills, m.Deaths)
}
func (i Match) FilterValue() string { return i.Id}

func getMatch(match valorantapi.Match) Match {
    i := Match{
        Id: match.Id,
        MapName: match.MapName,
        Mode: match.Mode,
        Kills: match.Kills,
        Deaths: match.Deaths,
        CharacterName: match.CharacterName,
    }
    return i
}

func MatchList() matchListState {
    puuid := valorantapi.GetAccountPUUID("Heri", "BLUB")
    resp := valorantapi.GetAccountMatches(puuid)
    formattedMatches := valorantapi.FormatMatches(resp)

    items := make([]list.Item, len(formattedMatches))
    for i := 0; i < len(formattedMatches); i++ { // Iterate safely
        items[i] = getMatch(formattedMatches[i])
    }

    m := matchListState{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
    return m
}
